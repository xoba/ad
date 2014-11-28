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
f= sin(3.14)
g=sin(x,y)
z= f+g
z2 = z*z
z3 = beta[0]*z
cos(x) = sin(x)
`

func Run(args []string) {
	lex := NewContext(NewLexer(strings.NewReader(formula)))
	yyParse(lex)
	if len(lex.errors) > 0 {
		log.Fatal(lex.errors)
	}
	defs := make(map[string]*Node)
	for i, s := range lex.statements {
		substitute(defs, s)
		fmt.Printf("statement %d. %s\n", i, s.Formula())
		defs[s.Children[0].S] = s.Children[1]
	}
}

func substitute(defs map[string]*Node, n *Node) {
	switch n.Type {
	case numberNT:
	case identifierNT:
		if v, ok := defs[n.S]; ok {
			n.CopyFrom(v)
		}
	case indexedIdentifierNT:
	case statementNT, functionNT:
		for _, c := range n.Children {
			substitute(defs, c)
		}
	default:
		panic("illegal type: " + n.Type)
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
	Type     NodeType `json:"T,omitempty"`
	S        string   `json:"S,omitempty"`
	F        float64  `json:"F,omitempty"`
	I        int      `json:"I,omitempty"`
	Children []*Node  `json:"C,omitempty"`
}

func (n *Node) CopyFrom(o *Node) {
	n.Type = o.Type
	n.S = o.S
	n.F = o.F
	n.I = o.I
	n.Children = o.Children
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
			fmt.Fprintf(buf, "(%s %s %s)", n.Children[0].Formula(), x, n.Children[1].Formula())
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
			for _, c := range n.Children {
				args = append(args, c.Formula())
			}
			fmt.Fprintf(buf, "%s(%s)", n.S, strings.Join(args, ", "))
		}
	case statementNT:
		fmt.Fprintf(buf, "%s := %s", n.Children[0].Formula(), n.Children[1].Formula())
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
		Type:     statementNT,
		Children: []*Node{lhs, rhs},
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
		Type:     functionNT,
		S:        ident,
		Children: args,
	}
}

func NewArgList(arg *Node) *Node {
	return &Node{
		Type:     argListNT,
		Children: []*Node{arg},
	}
}

func (n *Node) AddChild(c *Node) *Node {
	n.Children = append(n.Children, c)
	return n
}

func FunctionArgs(ident string, args *Node) *Node {
	if args.Type != argListNT {
		panic("illegal type: " + args.Type)
	}
	n := &Node{
		Type:     functionNT,
		S:        ident,
		Children: args.Children,
	}
	return n
}

func Negate(a *Node) *Node {
	return Function("multiply", Number(-1), a)
}
