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

f := sin(3.14)
abc(x,y) = (x+1)*y+f
xyz(z) = z^2
z=1+q
zz=2
g = abc(3.14,z) + f
h = abc(abc(1,2),abc(3,xyz(4)))

nn(beta[0],beta[1],x) = beta[0] *beta[1] + x

result := nn(1,2,x[6]) + nn(3,4,q)

`

func Run(args []string) {
	lex := NewContext(NewLexer(strings.NewReader(formula)))
	yyParse(lex)
	if len(lex.errors) > 0 {
		log.Fatal(lex.errors)
	}
	idents := make(map[string]*Node)
	funcs := make(map[string]*Node)
	for i, s := range lex.statements {
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
		fmt.Printf("definition %d. %s\n", i, s.Formula())
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
	log.Printf("oops: %v\n", e)
	c.errors = append(c.errors, e)
}

func LexNumber(s string) *Node {
	n, _ := strconv.ParseFloat(s, 64)
	return Number(n)
}

func LexIdentifier(s string) *Node {
	return &Node{
		Type: identifierNT,
		S:    s,
	}
}

func Number(n float64) *Node {
	return &Node{
		Type: numberNT,
		F:    n,
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
		S:    ident.S,
		I:    int(index.F),
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
	return Function("multiply", Number(-1), a)
}
