// parsing mathematical expressions
//
//go:generate nex lexer.nex
//go:generate go tool yacc -o parser.y.go parser.y
package parser

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"sort"
	"strings"
	"text/template"
	"time"
)

const (
	defaultPrivateString = "pvt"
	defaultDx            = 0.00001
	defaultTemplates     = "src/xoba/ad/parser/templates"
)

func Formula(n int) string {
	var terms []string
	for i := 0; i < n; i++ {
		terms = append(terms, fmt.Sprintf("x[%d]", i))
	}
	return "f:=z*" + strings.Join(terms, "*")
}

func computeDerivatives(w io.Writer, steps []Step, private string) {
	fmt.Fprintf(w, "grad_%s := make(map[string]float64)\n", private)
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

func Run(args []string) {
	var funcName, private, pkg, templates, formula, output string
	var dx float64
	var numerical, main, timeComment, funcs bool
	flags := flag.NewFlagSet("parse", flag.ExitOnError)
	flags.StringVar(&funcName, "name", "ComputeAD", "ad function name")
	flags.StringVar(&formula, "formula", Formula(20), "the formula to parse (or file)")
	flags.StringVar(&private, "private", defaultPrivateString, "the private variable string")
	flags.StringVar(&pkg, "package", "main", "the go package for generated code")
	flags.StringVar(&templates, "templates", defaultTemplates, "directory of go template functions")
	flags.StringVar(&output, "output", "compute.go", "name of go source file to output")
	flags.BoolVar(&main, "main", true, "whether to emit a main method")
	flags.BoolVar(&funcs, "funcs", true, "whether to emit template functions")
	flags.BoolVar(&timeComment, "time", true, "embed time in source code comment")
	flags.BoolVar(&numerical, "numerical", true, "whether to output numerical gradient code")
	flags.Float64Var(&dx, "dx", defaultDx, "infinitesimal for numerical differentiation")
	flags.Parse(args)

	if buf, err := ioutil.ReadFile(formula); err == nil {
		formula = string(buf)
	}

	code, err := Parse(numerical, funcs, main, timeComment, funcName, pkg, private, templates, formula, 0.00001)
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

func Parse(numerical, funcs, main, timeComment bool, name, pkg, private, templates, formula string, dx float64) ([]byte, error) {
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
	vp := &VarParser{
		vars:        vars,
		indexed:     make(map[string]string),
		maxIndicies: make(map[string]int),
	}
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
		var names []string
		for k := range vars {
			names = append(names, k)
		}
		sort.Strings(names)
		for _, k := range names {
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
	fmt.Fprintf(pgm, "if !computeGrad {\n")
	fmt.Fprintf(pgm, "return %s, nil;\n", y)
	fmt.Fprintln(pgm, "}")
	computeDerivatives(pgm, steps, private)

	indexedArgs := make(map[string]bool)
	scalarArgs := make(map[string]bool)

	var argList []string
	for v := range vars {
		argList = append(argList, v)
		if x, ok := vp.indexed[v]; ok {
			indexedArgs[x] = true
		} else {
			scalarArgs[v] = true
		}
	}
	sort.Strings(argList)

	scalarArgList := sort1(scalarArgs)
	indexedArgList := sort1(indexedArgs)

	args := func() string {
		var out []string
		if len(scalarArgList) > 0 {
			var list []string
			for _, a := range scalarArgList {
				list = append(list, a)
			}
			out = append(out, strings.Join(list, ","))
		}
		if len(indexedArgList) > 0 {
			var list []string
			for _, a := range indexedArgList {
				list = append(list, a)
			}
			out = append(out, strings.Join(list, ","))
		}
		return strings.Join(out, ", ")
	}()

	declArgs := func() string {
		var out []string
		if len(scalarArgList) > 0 {
			var list []string
			for _, a := range scalarArgList {
				list = append(list, a)
			}
			out = append(out, strings.Join(list, ",")+" float64")
		}
		if len(indexedArgList) > 0 {
			var list []string
			for _, a := range indexedArgList {
				list = append(list, a)
			}
			out = append(out, strings.Join(list, ",")+" []float64")
		}
		return strings.Join(out, ", ")
	}()

	f := new(bytes.Buffer)

	decls := new(bytes.Buffer)

	for _, v := range indexedArgList {
		fmt.Fprintf(decls, "%s := make([]float64,%d);\n", v, vp.maxIndicies[v])
	}
	for _, v := range argList {
		if _, ok := vp.indexed[v]; ok {
			fmt.Fprintf(decls, "%v = %.20f;\n", v, rand.NormFloat64())
			fmt.Fprintf(decls, "fmt.Printf(\"setting %s = %%+.20f\\n\",%s)\n", v, v)
		} else {
			fmt.Fprintf(decls, "%s := %.20f;\n", v, rand.NormFloat64())
			fmt.Fprintf(decls, "fmt.Printf(\"setting %s = %%+.20f\\n\",%s)\n", v, v)
		}
	}
	fmt.Fprintln(decls, `fmt.Println();`)

	var imports Imports
	if main {
		imports.Add("fmt")
		imports.Add("math")
	}

	returnSig := "(float64,map[string]float64)"

	t := template.Must(template.New("output.go").Parse(`{{ if .timeComment }} // autogenerated {{.time}} {{else}} // autogenerated, do not edit! {{end}}
// see https://github.com/xoba/ad

package {{.package}}
{{.imports}}

// automatically compute the value and gradient of {{.qformula}}
func {{.funcName}}(computeGrad bool,{{.varsDecl}}) {{.returnSig}} {
{{.program}} return {{.y}},grad_{{.private}};
}

{{ if .main }}
func main() {
fmt.Printf("running autodiff code of {{.time}} on %q\n\n", {{ printf "%q" .formula }});
{{.decls}} 

	c1, grad1 := {{.funcName}}(true,{{.vars}})
	fmt.Printf("autodiff value   : %.20f\n", c1)
	fmt.Printf("autodiff gradient: %v\n\n", grad1)

	c2, grad2 := ComputeNumerical(true,{{.vars}})
	fmt.Printf("numeric value   : %.20f\n", c2)
	fmt.Printf("numeric gradient: %v\n\n", grad2)

	var total float64
	add := func(n string, x float64) {
		fmt.Printf("%s difference: %+.20f\n", n, x)
		total += math.Abs(x)
	}
	add("value", c1-c2)
	for k, v := range grad2 {
		add(fmt.Sprintf("grad[%s]", k), grad1[k]-v)
	}
fmt.Printf("\nsum of absolute differences: %+.20f\n",total);
}
{{ end }}

{{if .numerical}}
// numerically compute the value and gradient of {{.qformula}}
func ComputeNumerical(computeGrad bool, {{.varsDecl}}) (float64,map[string]float64) {
grad_{{.private}} := make(map[string]float64)
{{.checker}} return tmp1_{{.private}},grad_{{.private}};
}
{{end}}

{{ if .emitFuncs }}
  {{.funcs}}
{{end}}

`))

	templateImports, code, err := GenTemplates(templates, private)
	if err != nil {
		return nil, err
	}
	imports.AddAll(templateImports)

	t.Execute(f, map[string]interface{}{
		"checker":     checker.String(),
		"decls":       decls.String(),
		"emitFuncs":   funcs,
		"formula":     formula,
		"funcName":    name,
		"funcs":       code,
		"imports":     imports.String(),
		"lhs":         lex.lhs.S,
		"main":        main,
		"package":     pkg,
		"private":     private,
		"program":     pgm.String(),
		"qformula":    fmt.Sprintf("%q", formula),
		"timeComment": timeComment,
		"time":        time.Now().UTC().Format("2006-01-02T15:04:05.000Z"),
		"vars":        args,
		"varsDecl":    declArgs,
		"y":           y,
		"numerical":   numerical,
		"returnSig":   returnSig,
	})

	//return f.Bytes(), nil
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
	i           int
	vars        map[string]string
	indexed     map[string]string
	maxIndicies map[string]int
}

// gets all the vars and name them in-place
func (v *VarParser) getVars(root *Node, private string) (out []Step) {
	switch root.Type {
	case identifierNT, indexedIdentifierNT:
		if _, ok := v.vars[root.S]; !ok {
			root.Name = fmt.Sprintf("v_%d_%s", v.i, private)
			v.vars[root.VarName()] = root.Name
			v.i++
			step := Step{
				decl: true,
				lhs:  root.Name,
				rhs:  root.S,
			}
			if root.Type == indexedIdentifierNT {
				step.rhs = fmt.Sprintf("%s[%d]", root.S, root.I)
				v.indexed[root.VarName()] = root.S
				if v.i > v.maxIndicies[root.S] {
					v.maxIndicies[root.S] = v.i
				}
			}
			out = append(out, step)
		} else {
			root.Name = v.vars[root.S]
		}
	case functionNT:
		for _, n := range root.Children {
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

func IndexedIdentifier(ident, index *Node) *Node {
	return &Node{
		Type: indexedIdentifierNT,
		S:    ident.S,
		I:    int(index.F),
	}
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

func sort1(m map[string]bool) (out []string) {
	for k, v := range m {
		if v {
			out = append(out, k)
		}
	}
	return
}
