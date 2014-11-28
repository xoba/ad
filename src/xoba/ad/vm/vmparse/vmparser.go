package vmparse

//go:generate go tool yacc -o vmparser.y.go vmparser.y
//go:generate nex vmlexer.nex

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
)

const formula = `
f := sin(3.14)
abc(x,y) = (x+1)*y
z=1+q
zz=2
g = abc(3.14,z) + f
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
	subsChildren := func() {
		for _, c := range n.C {
			substitute(idents, funcs, c)
		}
	}

	switch n.Type {
	case numberNT:
	case identifierNT:
		if v, ok := idents[n.S]; ok {
			n.CopyFrom(v)
		}
	case indexedIdentifierNT:
	case functionNT:
		subsChildren()
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
		panic(fmt.Sprintf("illegal type: %s", n.Type))
	}
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

type context struct {
	statements []*Node
	yyLexer
	errors []error
}

type NodeType string

const (
	argListNT           NodeType = "ARGS"
	numberNT                     = "NUMBER"
	identifierNT                 = "IDENTIFIER"
	indexedIdentifierNT          = "INDEXED IDENTIFIER"
	functionNT                   = "FUNCTION"
	statementNT                  = "STATEMENT"
)

type Node struct {
	Type NodeType `json:"T,omitempty"`
	S    string   `json:"S,omitempty"`
	F    float64  `json:"F,omitempty"`
	I    int      `json:"I,omitempty"`
	C    []*Node  `json:"C,omitempty"`
}

func (n *Node) DeepCopy() *Node {
	out := &Node{
		Type: n.Type,
		S:    n.S,
		F:    n.F,
		I:    n.I,
	}
	for _, c := range n.C {
		out.C = append(out.C, c.DeepCopy())
	}
	return out
}

func (n *Node) CopyFrom(o *Node) {
	n.Type = o.Type
	n.S = o.S
	n.F = o.F
	n.I = o.I
	n.C = o.C
}

func (n Node) String() string {
	buf, _ := json.Marshal(n)
	return string(buf)
}

func (n Node) Formula() string {
	buf := new(bytes.Buffer)
	switch n.Type {
	case numberNT:
		fmt.Fprintf(buf, "%f", n.F)
	case identifierNT:
		fmt.Fprintf(buf, "%s", n.S)
	case indexedIdentifierNT:
		fmt.Fprintf(buf, "%s[%d]", n.S, n.I)
	case functionNT:
		op := func(x string) {
			fmt.Fprintf(buf, "(%s %s %s)", n.C[0].Formula(), x, n.C[1].Formula())
		}
		switch n.S {
		case "multiply":
			op("*")
		case "divide":
			op("/")
		case "subtract":
			op("-")
		case "add":
			op("+")
		default:
			var args []string
			for _, c := range n.C {
				args = append(args, c.Formula())
			}
			fmt.Fprintf(buf, "%s(%s)", n.S, strings.Join(args, ", "))
		}
	case statementNT:
		fmt.Fprintf(buf, "%s := %s", n.C[0].Formula(), n.C[1].Formula())
	default:
		panic("illegal type: " + n.Type)
	}
	return buf.String()

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

func (n *Node) AddChild(c *Node) *Node {
	n.C = append(n.C, c)
	return n
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
