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
	"math/rand"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"
	"text/template"
)

// don't name any varibles like this prefix in your expressions:
const private = "x_a8b32499"

var formula string = Formula(3)

func Formula(n int) string {
	var terms []string
	for i := 0; i < n; i++ {
		term := fmt.Sprintf("%f * x%d", rand.Float64(), i)
		terms = append(terms, term)
	}
	return fmt.Sprintf("f := s2 + log(1 + exp(-y * (%s)))", strings.Join(terms, "+"))
}

func computeDerivatives(w io.Writer, steps []Step) {
	n := len(steps)
	for i := 0; i < n; i++ {
		j := n - i - 1
		step := steps[j]
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
		if step.decl {
			fmt.Fprintf(w, "grad_%s[\"%s\"] = b%s\n", private, step.rhs, step.lhs)
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
		list = append(list, fmt.Sprintf("d_%s_%s(%d,%s)", private, num.f, a, strings.Join(num.args, ",")))
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

	checker := new(bytes.Buffer)
	if true {
		dx := 0.00001
		fmt.Fprintf(checker, "delta := %f\n", dx)
		fmt.Fprintf(checker, `calc := func() float64 {
%s
return %s
}
`, formula, lex.lhs.S)
		fmt.Fprintln(checker, "tmp1 := calc()")
		for k := range vars {
			fmt.Fprintln(checker, "{")
			fmt.Fprintf(checker, "%s += delta\n", k)
			fmt.Fprintln(checker, "tmp2 := calc()")
			fmt.Fprintf(checker, "%s -= delta\n", k)
			fmt.Fprintf(checker, "grad_%s[%q] = (tmp2 - tmp1)/delta\n", private, k)
			fmt.Fprintln(checker, "}")
		}
	}

	pgm := new(bytes.Buffer)
	for _, s := range steps {
		fmt.Fprintln(pgm, s)
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
{{.decls}} 

	c1, grad1 := ComputeAD({{.vars}})
	fmt.Printf("ad value: %f\n", c1)
	fmt.Printf("ad grad : %v\n", grad1)

	c2, grad2 := ComputeNum({{.vars}})
	fmt.Printf("num value: %f\n", c2)
	fmt.Printf("num grad : %v\n", grad2)

	var total float64
	add := func(n string, x float64) {
		fmt.Printf("diff %s: %f\n", n, x)
		total += math.Abs(x)
	}
	add("value", c1-c2)
	for k, v := range grad2 {
		add(fmt.Sprintf("grad[%3s]", k), grad1[k]-v)
	}
fmt.Printf("*** total diffs: %f\n",total);
}

{{.funcs}}

func ComputeAD({{.vars}} float64) (float64,map[string]float64) {
grad_{{.private}} := make(map[string]float64)
{{.program}} return {{.y}},grad_{{.private}};
}

func ComputeNum({{.vars}} float64) (float64,map[string]float64) {
grad_{{.private}} := make(map[string]float64)
{{.checker}} return tmp1,grad_{{.private}};
}

`))

	t.Execute(f, map[string]interface{}{
		"decls":   decls.String(),
		"formula": formula,
		"lhs":     lex.lhs.S,
		"vars":    strings.Join(list, ", "),
		"program": pgm.String(),
		"checker": checker.String(),
		"y":       y,
		"funcs":   Functions(private),
		"private": private,
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
			rhs.Name = fmt.Sprintf("v_%d_%s", v.i, private)
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
	return fmt.Sprintf("%s := %s;", s.lhs, s.rhs)
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
			lhs:  fmt.Sprintf("s_%d_%s", s.i, private),
			rhs:  fmt.Sprintf("%s_%s(%s)", private, rhs.S, strings.Join(args, ",")),
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
