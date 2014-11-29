package vmparse

//go:generate go tool yacc -o vmparser.y.go vmparser.y
//go:generate nex vmlexer.nex

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"xoba/ad/vm"
)

const formula = `

f:= 5 + sqrt(a*b*sin(a))

`

func Run(args []string) {
	lex := NewContext(NewLexer(strings.NewReader(formula)))
	yyParse(lex)
	if len(lex.errors) > 0 {
		log.Fatalf("oops, errors: %q\n", lex.errors)
	}
	idents := make(map[string]*Node)
	funcs := make(map[string]*Node)
	for _, s := range lex.statements {
		lhs := s.C[0]
		rhs := s.C[1]
		substitute(idents, funcs, rhs)
		switch lhs.T {
		case identifierNT:
			idents[s.C[0].S] = s.C[1]
		case functionNT:
			funcs[s.C[0].S] = s
		default:
			panic("illegal lhs: " + lhs.T)
		}
		fmt.Println(s.Formula())
		var steps []Step
		var r int
		newreg := func() (out int) {
			out = r
			r++
			return
		}
		vars := AssignVars(rhs, newreg)
		fmt.Printf("vars = %v\n", vars)
		for k, v := range vars {
			r := newreg()
			steps = append(steps, Step{
				Register: r,
				Args:     []int{v},
				Code:     fmt.Sprintf("getinput %d %d // copy variable %s from input %d to register %d", v, r, k, v, r),
			})
		}
		steps = append(steps, CreateSteps(rhs, vars, newreg)...)
		for i, s := range steps {
			if len(s.Code) == 0 {
				continue
			}
			fmt.Printf("step %d: %#v\n", i, s)
		}
		buf := new(bytes.Buffer)
		fmt.Fprintf(buf, "inputs %d\n", len(vars))
		fmt.Fprintf(buf, "outputs 1\n")
		fmt.Fprintf(buf, "registers %d\n", r)
		for _, s := range steps {
			if len(s.Code) == 0 {
				continue
			}
			fmt.Fprintln(buf, s.Code)
		}
		fmt.Fprintf(buf, "setoutput %d 0\n", r-1)
		fmt.Println(buf.String())
		p := vm.Compile(buf)
		out := make([]float64, p.Outputs)
		in := make([]float64, p.Inputs)
		model := make([]float64, p.Models)
		in[0] = 3
		in[1] = 4
		var e vm.Executor = vm.Execute
		check(e(p, model, in, out))
		a := in[0]
		b := in[1]
		fmt.Printf("output = %.3f vs %.3f\n", out, 5+math.Sqrt(a*b*math.Sin(a)))

	}
}

type Step struct {
	Register int    // register being assigned
	Function string // the function being computed
	Args     []int  // register args of the function
	Code     string
}

func CreateSteps(n *Node, vars map[string]int, newreg func() int) (out []Step) {
	switch n.T {
	case numberNT:
		r := newreg()
		n.Step = Step{
			Register: r,
			Code:     fmt.Sprintf("literal %d %f", r, n.Float64()),
		}
		out = append(out, n.Step)
	case identifierNT, indexedIdentifierNT:
		n.Step = Step{
			Register: vars[n.S],
		}
		out = append(out, n.Step)
	case functionNT:
		for _, c := range n.C {
			out = append(out, CreateSteps(c, vars, newreg)...)
		}
		r := newreg()
		var args []int
		n.Step = Step{
			Register: r,
			Function: n.S,
			Args:     args,
		}
		switch len(n.C) {
		case 1:
			n.Step.Code = fmt.Sprintf("%s %d %d", n.S, n.C[0].Step.Register, r)
		case 2:
			n.Step.Code = fmt.Sprintf("%s %d %d %d", n.S, n.C[0].Step.Register, n.C[1].Step.Register, r)
		default:
			panic("illegal state for " + n.S)
		}
		out = append(out, n.Step)
	default:
		panic("illegal type: " + n.T)
	}
	return
}

func AssignVars(n *Node, newreg func() int) map[string]int {
	out := make(map[string]int)
	add := func(v string) {
		if _, ok := out[v]; !ok {
			out[v] = len(out)
		}
	}
	switch n.T {
	case identifierNT, indexedIdentifierNT:
		add(n.S)
	}
	for _, c := range n.C {
		for k := range AssignVars(c, newreg) {
			add(k)
		}
	}
	return out
}

func substitute(idents, funcs map[string]*Node, n *Node) {
	switch n.T {
	case numberNT:
	case identifierNT, indexedIdentifierNT:
		if v, ok := idents[n.S]; ok {
			n.CopyFrom(v)
		}
	case functionNT:
		for _, c := range n.C {
			substitute(idents, funcs, c)
		}
		if v, ok := funcs[n.S]; ok {
			lhsDef := v.C[0]
			newIdents := make(map[string]*Node)
			for i := 0; i < len(lhsDef.C); i++ {
				lhs := lhsDef.C[i]
				rhs := n.C[i]
				newIdents[lhs.S] = rhs
			}
			c := v.C[1].DeepCopy()
			substitute(newIdents, nil, c)
			n.CopyFrom(c)
		}
	default:
		panic(fmt.Sprintf("illegal type: %s", n.T))
	}
}

type context struct {
	statements []*Node
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
	lexer := c.yyLexer.(*Lexer)
	c.errors = append(c.errors, fmt.Errorf("in line %d: %v", lexer.Line(), e))
}

func LexIdentifier(s string) *Node {
	return &Node{
		T: identifierNT,
		S: s,
	}
}

func LexNumber(n string) *Node {
	if _, err := strconv.ParseFloat(n, 64); err != nil {
		check(err)
	}
	return &Node{
		T: numberNT,
		S: n,
	}
}

func NewStatement(lhs, rhs *Node) *Node {
	return &Node{
		T: statementNT,
		C: []*Node{lhs, rhs},
	}
}

func IndexedIdentifier(ident, index *Node) *Node {
	return &Node{
		T: indexedIdentifierNT,
		S: fmt.Sprintf("%s[%s]", ident.S, index.S),
	}
}

func Function(ident string, args ...*Node) *Node {
	return &Node{
		T: functionNT,
		S: ident,
		C: args,
	}
}

func NewArgList(arg *Node) *Node {
	return &Node{
		T: argListNT,
		C: []*Node{arg},
	}
}

func FunctionArgs(ident string, args *Node) *Node {
	if args.T != argListNT {
		panic("illegal type: " + args.T)
	}
	n := &Node{
		T: functionNT,
		S: ident,
		C: args.C,
	}
	return n
}

func Negate(a *Node) *Node {
	return Function("multiply", LexNumber("-1"), a)
}
