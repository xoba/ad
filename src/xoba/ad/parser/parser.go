// parsing mathematical expressions
//
//go:generate nex lexer.nex
//go:generate go tool yacc -o parser.y.go parser.y
package parser

import (
	"bytes"
	"encoding/json"
	"flag"
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

func Formula(n int) string {
	if true {
		var terms []string
		for i := 0; i < n; i++ {
			terms = append(terms, fmt.Sprintf("x%d", i))
		}
		return "f:=" + strings.Join(terms, "*")
	}
	var terms []string
	for i := 0; i < n; i++ {
		term := fmt.Sprintf("%f * x%d", rand.Float64(), i)
		terms = append(terms, term)
	}
	return fmt.Sprintf("f := sqrt(a*x) + pow(a,b) + a-b+c/d + log(1 + exp(-y * (%s)))", strings.Join(terms, "+"))
}

func computeDerivatives(w io.Writer, steps []Step, private string) {
	n := len(steps)
	for i := 0; i < n; i++ {
		j := n - i - 1
		step := steps[j]
		x0 := 0.0
		if i == 0 {
			x0 = 1.0
		}
		fmt.Fprintf(w, "b_%s := %f\n", step.lhs, x0)
		for k := j + 1; k < n; k++ {
			s3 := steps[k]
			if d := derivative(s3, step, private); len(d) > 0 {
				fmt.Fprintf(w, "b_%s += b_%s * (%s)\n", step.lhs, s3.lhs, d)
			}
		}
		if step.decl {
			fmt.Fprintf(w, "grad_%s[\"%s\"] = b_%s\n", private, step.rhs, step.lhs)
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

func derivative(num, denom Step, private string) string {
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
		list = append(list, fmt.Sprintf("d_%s_%s(%d,%s)", num.f, private, a, strings.Join(num.args, ",")))
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
	var private, templates, formula, output string
	flags := flag.NewFlagSet("parse", flag.ExitOnError)
	flags.StringVar(&formula, "formula", Formula(10), "the formula to parse")
	flags.StringVar(&private, "private", "x", "the private variable string")
	flags.StringVar(&templates, "templates", "src/xoba/ad/parser/templates", "directory of go template functions")
	flags.StringVar(&output, "output", "compute.go", "name of go program to output")
	flags.Parse(args)

	code := Parse(private, templates, formula)

	f, err := os.Create(output)
	check(err)
	f.Write(code)
	f.Close()
	if err := Gofmt(output); err != nil {
		log.Fatalf("oops: %v\n", err)
	}

}

func Parse(private, templates, formula string) []byte {
	lex := NewContext(NewLexer(strings.NewReader(formula)))
	yyParse(lex)
	vars := make(map[string]string)
	vp := &VarParser{vars: vars}
	sp := &StepParser{vars: vars}
	var steps []Step
	steps = append(steps, vp.getVars(lex.rhs, private)...)
	steps = append(steps, sp.program(lex.rhs, private)...)
	y := steps[len(steps)-1].lhs

	checker := new(bytes.Buffer)
	{
		dx := 0.00001
		fmt.Fprintf(checker, "delta_%s := %f\n", private, dx)
		fmt.Fprintf(checker, `calc_%s := func() float64 {
%s
return %s
}
`, private, formula, lex.lhs.S)
		fmt.Fprintf(checker, "tmp1_%s := calc_%s()\n", private, private)
		for k := range vars {
			fmt.Fprintln(checker, "{")
			fmt.Fprintf(checker, "%s += delta_%s\n", k, private)
			fmt.Fprintf(checker, "tmp2_%s := calc_%s()\n", private, private)
			fmt.Fprintf(checker, "%s -= delta_%s\n", k, private)
			fmt.Fprintf(checker, "grad_%s[%q] = (tmp2_%s - tmp1_%s)/delta_%s\n", private, k, private, private, private)
			fmt.Fprintln(checker, "}")
		}
	}

	pgm := new(bytes.Buffer)
	for _, s := range steps {
		fmt.Fprintln(pgm, s)
	}
	computeDerivatives(pgm, steps, private)
	var list []string
	for v := range vars {
		list = append(list, v)
	}
	sort.Strings(list)

	f := new(bytes.Buffer)

	decls := new(bytes.Buffer)
	for _, v := range list {
		fmt.Fprintf(decls, "%s := rand.Float64();\n", v)
		fmt.Fprintf(decls, "fmt.Printf(\"setting %s = %%f\\n\",%s)\n", v, v)
	}
	fmt.Fprintln(decls, `fmt.Println();`)

	var imports Imports
	imports.Add("fmt")
	imports.Add("math")
	imports.Add("time")
	imports.Add("math/rand")

	t := template.Must(template.New("output.go").Parse(`package main
{{.imports}}
func main() {
fmt.Println("running autodiff code on {{.formula}}\n");
rand.Seed(time.Now().UTC().UnixNano())
{{.decls}} 

	c1, grad1 := ComputeAD({{.vars}})
	fmt.Printf("autodiff value   : %f\n", c1)
	fmt.Printf("autodiff gradient: %v\n\n", grad1)

	c2, grad2 := ComputeNumerical({{.vars}})
	fmt.Printf("numeric value   : %f\n", c2)
	fmt.Printf("numeric gradient: %v\n\n", grad2)

	var total float64
	add := func(n string, x float64) {
		fmt.Printf("%s difference: %f\n", n, x)
		total += math.Abs(x)
	}
	add("value", c1-c2)
	for k, v := range grad2 {
		add(fmt.Sprintf("grad[%3s]", k), grad1[k]-v)
	}
fmt.Printf("\nsum of absolute differences: %f\n",total);
}

{{.funcs}}

func ComputeAD({{.vars}} float64) (float64,map[string]float64) {
grad_{{.private}} := make(map[string]float64)
{{.program}} return {{.y}},grad_{{.private}};
}

func ComputeNumerical({{.vars}} float64) (float64,map[string]float64) {
grad_{{.private}} := make(map[string]float64)
{{.checker}} return tmp1_{{.private}},grad_{{.private}};
}

`))

	templateImports, code := GenTemplates(templates, private)
	imports.AddAll(templateImports)

	t.Execute(f, map[string]interface{}{
		"decls":   decls.String(),
		"formula": formula,
		"lhs":     lex.lhs.S,
		"vars":    strings.Join(list, ", "),
		"program": pgm.String(),
		"checker": checker.String(),
		"y":       y,
		"funcs":   code,
		"private": private,
		"imports": imports.String(),
	})

	return f.Bytes()
}

type Imports struct {
	list []string
}

func (i Imports) String() string {
	m := make(map[string]bool)
	for _, x := range i.list {
		x = strings.Replace(x, `"`, ``, -1)
		m[x] = true
	}
	out := new(bytes.Buffer)
	fmt.Fprintf(out, "import (\n")
	for k := range m {
		fmt.Fprintf(out, "%q\n", k)
	}
	fmt.Fprintf(out, ")\n")
	return out.String()
}

func (i *Imports) Add(x string) {
	i.list = append(i.list, x)
}

func (i *Imports) AddAll(list []string) {
	for _, x := range list {
		i.Add(x)
	}
}

func formatImports(list []string) string {
	var out []string
	m := make(map[string]bool)
	for _, x := range list {
		m[x] = true
	}
	for k := range m {
		out = append(out, k)
	}
	return strings.Join(out, "\n")
}

func Gofmt(p string) error {
	cmd := exec.Command("gofmt", "-w", p)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

type VarParser struct {
	i    int
	vars map[string]string
}

// gets all the vars and name them in-place
func (v *VarParser) getVars(rhs *Node, private string) (out []Step) {
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
			out = append(out, v.getVars(n, private)...)
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

func (s *StepParser) program(rhs *Node, private string) (out []Step) {
	switch rhs.Type {
	case numberNT:
		rhs.Name = fmt.Sprintf("%f", rhs.F)
	case functionNT:
		var args []string
		for _, n := range rhs.Children {
			steps := s.program(n, private)
			out = append(out, steps...)
			args = append(args, n.Name)
		}
		step := Step{
			lhs:  fmt.Sprintf("s_%d_%s", s.i, private),
			rhs:  fmt.Sprintf("%s_%s(%s)", rhs.S, private, strings.Join(args, ",")),
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
