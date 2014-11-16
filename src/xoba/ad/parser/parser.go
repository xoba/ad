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
	"time"
)

func Formula(n int) string {
	var terms []string
	for i := 0; i < n; i++ {
		terms = append(terms, fmt.Sprintf("x%d", i))
	}
	return "f:=" + strings.Join(terms, "*")
}

func computeDerivatives(w io.Writer, steps []Step, private string) {
	n := len(steps)
	for i := 0; i < n; i++ {
		j := n - i - 1
		step := steps[j]
		if i == 0 {
			fmt.Fprintf(w, "const b_%s = 1.0\n", step.lhs)
		} else {
			fmt.Fprintf(w, "b_%s := 0.0\n", step.lhs)
		}
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
	singleArg := len(num.args) == 1
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
		if singleArg {
			list = append(list, fmt.Sprintf("d_%s_%s(%s)", num.f, private, strings.Join(num.args, ",")))
		} else {
			list = append(list, fmt.Sprintf("d_%s_%s(%d,%s)", num.f, private, a, strings.Join(num.args, ",")))
		}
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
	buf, _ := json.Marshal(n)
	return string(buf)
}

func Run(args []string) {
	var private, templates, formula, output string
	var dx float64
	flags := flag.NewFlagSet("parse", flag.ExitOnError)
	flags.StringVar(&formula, "formula", Formula(10), "the formula to parse")
	flags.StringVar(&private, "private", defaultPrivateString, "the private variable string")
	flags.StringVar(&templates, "templates", defaultTemplates, "directory of go template functions")
	flags.StringVar(&output, "output", "compute.go", "name of go program to output")
	flags.Float64Var(&dx, "dx", defaultDx, "infinitesimal for numerical differentiation")
	flags.Parse(args)

	code, err := Parse(private, templates, formula, 0.00001)
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Create(output)
	if err != nil {
		log.Fatal(err)
	}
	f.Write(code)
	f.Close()
}

const (
	defaultPrivateString = "pvt"
	defaultDx            = 0.00001
	defaultTemplates     = "src/xoba/ad/parser/templates"
)

func Parse(private, templates, formula string, dx float64) ([]byte, error) {
	if len(private) == 0 {
		private = defaultPrivateString
	}
	if dx == 0 {
		dx = defaultDx
	}
	if len(templates) == 0 {
		templates = defaultTemplates
	}
	lex := NewContext(NewLexer(strings.NewReader(formula)))
	yyParse(lex)
	if len(lex.errors) > 0 {
		return nil, lex.errors[0]
	}
	if lex.lhs == nil || lex.rhs == nil {
		return nil, fmt.Errorf("parse error: lhs or rhs is nil")
	}
	vars := make(map[string]string)
	vp := &VarParser{vars: vars}
	sp := &StepParser{vars: vars}
	var steps []Step
	steps = append(steps, vp.getVars(lex.rhs, private)...)
	steps = append(steps, sp.program(lex.rhs, private)...)
	y := steps[len(steps)-1].lhs

	checker := new(bytes.Buffer)
	{
		fmt.Fprintf(checker, "const delta_%s = %f\n", private, dx)
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
		fmt.Fprintf(decls, "%s := %.20f;\n", v, rand.NormFloat64())
		fmt.Fprintf(decls, "fmt.Printf(\"setting %s = %%.20f\\n\",%s)\n", v, v)
	}
	fmt.Fprintln(decls, `fmt.Println();`)

	var imports Imports
	imports.Add("fmt")
	imports.Add("math")

	t := template.Must(template.New("output.go").Parse(`// created {{.time}}
// see https://github.com/xoba/ad

package main
{{.imports}}

// automatically compute the value and gradient of {{.qformula}}
func ComputeAD({{.vars}} float64) (float64,map[string]float64) {
grad_{{.private}} := make(map[string]float64)
{{.program}} return {{.y}},grad_{{.private}};
}

func main() {
fmt.Printf("running autodiff code of {{.time}} on %q\n\n", {{ printf "%q" .formula }});
{{.decls}} 

	c1, grad1 := ComputeAD({{.vars}})
	fmt.Printf("autodiff value   : %.20f\n", c1)
	fmt.Printf("autodiff gradient: %v\n\n", grad1)

	c2, grad2 := ComputeNumerical({{.vars}})
	fmt.Printf("numeric value   : %.20f\n", c2)
	fmt.Printf("numeric gradient: %v\n\n", grad2)

	var total float64
	add := func(n string, x float64) {
		fmt.Printf("%s difference: %.20f\n", n, x)
		total += math.Abs(x)
	}
	add("value", c1-c2)
	for k, v := range grad2 {
		add(fmt.Sprintf("grad[%3s]", k), grad1[k]-v)
	}
fmt.Printf("\nsum of absolute differences: %.20f\n",total);
}

// numerically compute the value and gradient of {{.qformula}}
func ComputeNumerical({{.vars}} float64) (float64,map[string]float64) {
grad_{{.private}} := make(map[string]float64)
{{.checker}} return tmp1_{{.private}},grad_{{.private}};
}

{{.funcs}}


`))

	templateImports, code, err := GenTemplates(templates, private)
	if err != nil {
		return nil, err
	}
	imports.AddAll(templateImports)

	t.Execute(f, map[string]interface{}{
		"time":     time.Now().UTC().Format("2006-01-02T15:04:05.000Z"),
		"decls":    decls.String(),
		"formula":  formula,
		"qformula": fmt.Sprintf("%q", formula),
		"lhs":      lex.lhs.S,
		"vars":     strings.Join(list, ", "),
		"program":  pgm.String(),
		"checker":  checker.String(),
		"y":        y,
		"funcs":    code,
		"private":  private,
		"imports":  imports.String(),
	})

	return GofmtBuffer(f.Bytes())
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

func GofmtBuffer(code []byte) ([]byte, error) {
	out := new(bytes.Buffer)
	cmd := exec.Command("gofmt")
	cmd.Stdin = bytes.NewReader(code)
	cmd.Stdout = out
	if err := cmd.Run(); err != nil {
		return nil, err
	}
	return out.Bytes(), nil
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
	n, _ := strconv.ParseFloat(s, 64)
	return Number(n)
}

func Negate(a *Node) *Node {
	return Function("multiply", Number(-1), a)
}

type context struct {
	lhs *Node
	rhs *Node
	yyLexer
	errors []error
}

func NewContext(y yyLexer) *context {
	return &context{yyLexer: y}
}

func (c *context) Error(e string) {
	c.Error2(fmt.Errorf("%s", e))
}

func (c *context) Error2(e error) {
	log.Printf("oops: %v\n", e)
	c.errors = append(c.errors, e)
}
