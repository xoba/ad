// parsing mathematical expressions
//
//go:generate nex lexer.nex
//go:generate go tool yacc -o parser.y.go parser.y
package parser

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"
	"text/template"
)

const (
	formula  = `a := log(z+y*x*(exp(x) + x*y+x/y))`
	formula2 = `a := x+sqrt(pow(z,3)) + 99*(-5*x + 55)/6 + y`
)

func computeDerivatives(w io.Writer, steps []Step) {
	n := len(steps)
	for i := 0; i < n; i++ {
		j := n - i - 1
		step := steps[j]
		//fmt.Printf("%d. %s\n", j, step)
		x0 := 0.0
		if i == 0 {
			x0 = 1.0
		}
		fmt.Fprintf(w, "b%s := %f\n", step.lhs, x0)
		for k := j + 1; k < n; k++ {
			s3 := steps[k]
			if d := derivative(s3, step); len(d) > 0 {
				fmt.Fprintf(w, "b%s += b%s * %s\n", step.lhs, s3.lhs, d)
			}
		}
		//	fmt.Fprintf(w, "fmt.Printf(\"b%s = %%f\\n\",b%s) // %v: %s\n", step.lhs, step.lhs, step.decl, step.rhs)
		if step.decl {
			fmt.Fprintf(w, "grad[\"%s\"] = b%s\n", step.rhs, step.lhs)
		}
	}
}

type Step struct {
	decl bool
	lhs  string
	rhs  string
	f    string
	args []string
}

func derivative(num, denom Step) string {
	var args []int
	for i, a := range num.args {
		if a == denom.lhs {
			args = append(args, i)
		}
	}
	if len(args) == 0 {
		return ""
	}
	var list []string
	for _, a := range args {
		list = append(list, fmt.Sprintf("d%s(%d,%s)", num.f, a, strings.Join(num.args, ",")))
	}
	return strings.Join(list, "+")
}

type NodeType string

const (
	numberNT     NodeType = "NUM"
	identifierNT NodeType = "IDENT"
	functionNT   NodeType = "FUNC"
)

type Node struct {
	Type     NodeType `json:"T,omitempty"`
	S        string   `json:",omitempty"`
	F        float64  `json:",omitempty"`
	Children []*Node  `json:"C,omitempty"`
	Name     string   `json:"N,omitempty"`
}

func (n Node) String() string {
	buf, err := json.Marshal(n)
	check(err)
	return string(buf)
}

func Run(args []string) {
	lex := NewContext(NewLexer(strings.NewReader(formula)))
	yyParse(lex)
	vars := make(map[string]string)
	vp := &VarParser{vars: vars}
	sp := &StepParser{vars: vars}
	var steps []Step
	steps = append(steps, vp.getVars(lex.rhs)...)
	steps = append(steps, sp.program(lex.rhs)...)
	y := steps[len(steps)-1].lhs

	cg := new(bytes.Buffer)
	if true {
		dx := 0.00001
		fmt.Fprintf(cg, "delta := %f\n", dx)
		fmt.Fprintf(cg, "tmp := %s\n", lex.lhs.S)
		for k := range vars {
			fmt.Fprintln(cg, "{")
			fmt.Fprintf(cg, "%s += delta\n", k)
			fmt.Fprintln(cg, formula)
			fmt.Fprintf(cg, "%s -= delta\n", k)
			fmt.Fprintf(cg, "fmt.Printf(\"df/d%s = %%f\\n\",(a-tmp)/delta)\n", k)
			fmt.Fprintln(cg, "}")
		}
	}

	pgm := new(bytes.Buffer)
	for _, s := range steps {
		fmt.Fprintln(pgm, s)
		fmt.Fprintf(pgm, "fmt.Printf(\"%s = %%f\\n\",%s);\n", s.lhs, s.lhs)
	}
	computeDerivatives(pgm, steps)
	var list []string
	for v := range vars {
		list = append(list, v)
	}
	sort.Strings(list)
	f, err := os.Create("compute.go")
	check(err)

	decls := new(bytes.Buffer)
	for _, v := range list {
		fmt.Fprintf(decls, "%s := rand.Float64();\n", v)
		fmt.Fprintf(decls, "fmt.Printf(\"%s = %%f\\n\",%s)\n", v, v)
	}

	t := template.Must(template.New("compute.go").Parse(`package main
import (
"fmt"
"math"
"time"
"math/rand"
)
func main() {
fmt.Println("running compute.go");
rand.Seed(time.Now().UTC().UnixNano())
{{.decls}} {{.formula}} 
fmt.Printf("formula: %f\n",{{.lhs}});
c, grad:= Compute({{.vars}})
fmt.Printf("parsed : %f\n",c)
fmt.Printf("diff   : %f\n",{{.lhs}}-c)
fmt.Printf("grad = %v\n",grad);
{{.computeGradients}}

}

{{.funcs}}

func Compute({{.vars}} float64) (float64,map[string]float64) {
grad := make(map[string]float64)
{{.program}} return {{.y}},grad;
}

`))

	const funcs = `func add(a, b float64) float64 {
	return a + b
}
func dadd(i int, a, b float64) float64 {
	return 1
}

func multiply(a, b float64) float64 {
	return a * b
}
func dmultiply(i int, a, b float64) float64 {
	switch i {
	case 0:
		return b
	case 1:
		return a
	default:
		panic("illegal index")
	}
}

func subtract(a, b float64) float64 {
	return a - b
}
func dsubtract(i int, a, b float64) float64 {
	switch i {
	case 0:
		return 1
	case 1:
		return -1
	default:
		panic("illegal index")
	}
}

func divide(a, b float64) float64 {
	return a / b
}
func ddivide(i int, a, b float64) float64 {
	switch i {
	case 0:
		return 1 / b
	case 1:
		return -a / (b * b)
	default:
		panic("illegal index")
	}
	panic("unimplemented")
}

func sqrt(a float64) float64 {
	return math.Sqrt(a)
}
func dsqrt(_int, a float64) float64 {
	return 0.5 * math.Pow(a, -1.5)
}

func exp(a float64) float64 {
	return math.Exp(a)
}
func dexp(_ int, a float64) float64 {
	return exp(a)
}

func log(a float64) float64 {
	return math.Log(a)
}
func dlog(_ int, a float64) float64 {
	return 1 / a
}

func pow(a, b float64) float64 {
	return math.Pow(a, b)
}
func dpow(i int, a, b float64) float64 {
	panic("unimplemented")
}
`

	t.Execute(f, map[string]interface{}{
		"decls":            decls.String(),
		"formula":          formula,
		"lhs":              lex.lhs.S,
		"vars":             strings.Join(list, ", "),
		"program":          pgm.String(),
		"y":                y,
		"funcs":            funcs,
		"computeGradients": cg.String(),
	})

	f.Close()
	cmd := exec.Command("gofmt", "-w", "compute.go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("oops: %v\n", err)
	}
}

type VarParser struct {
	i    int
	vars map[string]string
}

// gets all the vars and names them in tree
func (v *VarParser) getVars(rhs *Node) (out []Step) {
	switch rhs.Type {
	case identifierNT:
		if _, ok := v.vars[rhs.S]; !ok {
			rhs.Name = fmt.Sprintf("v%d", v.i)
			v.vars[rhs.S] = rhs.Name
			v.i++
			step := Step{
				decl: true,
				lhs:  rhs.Name,
				rhs:  rhs.S,
			}
			out = append(out, step)
		} else {
			rhs.Name = v.vars[rhs.S]
		}
	case functionNT:
		for _, n := range rhs.Children {
			out = append(out, v.getVars(n)...)
		}
	}
	return
}

type StepParser struct {
	i    int
	vars map[string]string
}

func (s Step) String() string {
	if true {
		return fmt.Sprintf("%s := %s;", s.lhs, s.rhs)
	} else {
		return fmt.Sprintf("%s := %s; // f = %q; args = %q", s.lhs, s.rhs, s.f, s.args)
	}
}

func (s *StepParser) program(rhs *Node) (out []Step) {
	switch rhs.Type {
	case numberNT:
		rhs.Name = fmt.Sprintf("%f", rhs.F)
	case functionNT:
		var args []string
		for _, n := range rhs.Children {
			steps := s.program(n)
			out = append(out, steps...)
			args = append(args, n.Name)
		}
		step := Step{
			lhs:  fmt.Sprintf("s%d", s.i),
			rhs:  fmt.Sprintf("%s(%s)", rhs.S, strings.Join(args, ",")),
			f:    rhs.S,
			args: args,
		}
		out = append(out, step)
		rhs.Name = step.lhs
		s.i++
	}
	return out
}

func Function(ident string, args ...*Node) *Node {
	return &Node{
		Type:     functionNT,
		S:        ident,
		Children: args,
	}
}

func Number(n float64) *Node {
	return &Node{
		Type: numberNT,
		F:    n,
	}
}

func LexIdentifier(s string) *Node {
	return &Node{
		Type: identifierNT,
		S:    s,
	}
}

func LexNumber(s string) *Node {
	n, err := strconv.ParseFloat(s, 64)
	check(err)
	return Number(n)
}

func Negate(a *Node) *Node {
	return Function("multiply", Number(-1), a)
}

type context struct {
	lhs *Node
	rhs *Node
	yyLexer
}

func NewContext(y yyLexer) *context {
	return &context{yyLexer: y}
}

func (context) Error(e string) {
	log.Printf("oops: %v\n", e)
	os.Exit(1)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
