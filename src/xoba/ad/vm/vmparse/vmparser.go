package vmparse

import "strconv"

//go:generate go tool yacc -o vmparser.y.go vmparser.y
//go:generate nex vmlexer.nex
func Run(args []string) {

}

type Statement struct {
	lhs *Node
	rhs *Node
}

type context struct {
	lhs        *Node
	rhs        *Node
	statements []Statement
	yyLexer
	errors []error
}

type NodeType string

const (
	numberNT            NodeType = "NUM"
	identifierNT        NodeType = "IDENT"
	indexedIdentifierNT NodeType = "INDEXED"
	functionNT          NodeType = "FUNC"
)

type Node struct {
	Type     NodeType `json:"T,omitempty"`
	S        string   `json:",omitempty"`
	F        float64  `json:",omitempty"`
	I        int      `json:",omitempty"`
	Children []*Node  `json:"C,omitempty"`
	Name     string   `json:"N,omitempty"` // name of variable assigned by parser

	// for vm-assembly version:
	register, output int
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

func Negate(a *Node) *Node {
	return Function("multiply", Number(-1), a)
}
