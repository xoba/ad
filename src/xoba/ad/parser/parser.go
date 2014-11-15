// parsing mathematical expressions
//
//go:generate nex lexer.nex
//go:generate go tool yacc -o parser.y.go parser.y
package parser

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"
	"text/template"
)

const formula = `a := x+sqrt(pow(z,3)) + 99*(-5*x + 55)/6 + y`

func computeDerivatives(steps []Step) {
	for i := 0; i < len(steps); i++ {
		j := len(steps) - i - 1
		fmt.Println(steps[j])
	}
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
	pgm := new(bytes.Buffer)
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

	computeDerivatives(steps)
	for _, s := range steps {
		fmt.Fprintln(pgm, s)
	}
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
rand.Seed(time.Now().UTC().UnixNano())
{{.decls}} {{.formula}} 
fmt.Printf("formula: %f\n",{{.lhs}});
fmt.Printf("parsed : %f\n",Compute({{.vars}}))
fmt.Printf("diff   : %f\n",{{.lhs}}-Compute({{.vars}}))

{{.computeGradients}}

}

func Compute({{.vars}} float64) float64 {
{{.program}} return {{.y}};
}

func add(a, b float64) float64 {
	return a + b
}

func multiply(a, b float64) float64 {
	return a * b
}

func subtract(a, b float64) float64 {
	return a - b
}

func divide(a, b float64) float64 {
	return a / b
}

func sqrt(a float64) float64 {
	return math.Sqrt(a)
}

func exp(a float64) float64 {
	return math.Exp(a)
}

func log(a float64) float64 {
	return math.Log(a)
}

func pow(a, b float64) float64 {
	return math.Pow(a, b)
}
`))

	t.Execute(f, map[string]interface{}{
		"decls":            decls.String(),
		"formula":          formula,
		"lhs":              lex.lhs.S,
		"vars":             strings.Join(list, ", "),
		"program":          pgm.String(),
		"y":                y,
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

type Step struct {
	decl bool
	lhs  string
	rhs  string
	f    string
	args []string
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
