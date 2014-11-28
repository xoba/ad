package vmparse

//go:generate go tool yacc -o vmparser.y.go vmparser.y
//go:generate nex vmlexer.nex

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

const formula = `

f := sqrt(a*b*sin(a))
g = x[0] * y[0]
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
		switch lhs.Type {
		case identifierNT:
			idents[s.C[0].S] = s.C[1]
		case functionNT:
			funcs[s.C[0].S] = s
		default:
			panic("illegal lhs: " + lhs.Type)
		}
		fmt.Println(s.Formula())
	}
}

func substitute(idents, funcs map[string]*Node, n *Node) {
	switch n.Type {
	case numberNT:
	case identifierNT, indexedIdentifierNT:
		if v, ok := idents[n.Name()]; ok {
			n.CopyFrom(v)
		}
	case functionNT:
		for _, c := range n.C {
			substitute(idents, funcs, c)
		}
		if v, ok := funcs[n.Name()]; ok {
			lhsDef := v.C[0]
			newIdents := make(map[string]*Node)
			for i := 0; i < len(lhsDef.C); i++ {
				lhs := lhsDef.C[i]
				rhs := n.C[i]
				newIdents[lhs.Name()] = rhs
			}
			c := v.C[1].DeepCopy()
			substitute(newIdents, nil, c)
			n.CopyFrom(c)
		}
	default:
		panic(fmt.Sprintf("illegal type: %s", n.Type))
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
		Type: identifierNT,
		S:    s,
	}
}

func LexNumber(n string) *Node {
	if _, err := strconv.ParseFloat(n, 64); err != nil {
		check(err)
	}
	return &Node{
		Type: numberNT,
		S:    n,
	}
}

func NewStatement(lhs, rhs *Node) *Node {
	return &Node{
		Type: statementNT,
		C:    []*Node{lhs, rhs},
	}
}

func IndexedIdentifier(ident, index *Node) *Node {
	return &Node{
		Type: indexedIdentifierNT,
		S:    fmt.Sprintf("%s[%s]", ident.S, index.S),
	}
}

func Function(ident string, args ...*Node) *Node {
	return &Node{
		Type: functionNT,
		S:    ident,
		C:    args,
	}
}

func NewArgList(arg *Node) *Node {
	return &Node{
		Type: argListNT,
		C:    []*Node{arg},
	}
}

func FunctionArgs(ident string, args *Node) *Node {
	if args.Type != argListNT {
		panic("illegal type: " + args.Type)
	}
	n := &Node{
		Type: functionNT,
		S:    ident,
		C:    args.C,
	}
	return n
}

func Negate(a *Node) *Node {
	return Function("multiply", LexNumber("-1"), a)
}
