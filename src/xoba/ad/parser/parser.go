// parsing mathematical expressions
package parser

import "os"

//go:generate nex lexer.nex
//go:generate go tool yacc -o parser.y.go parser.y
const (
	numberNT NodeType = "NUM"
)

type NodeType string

type Node struct {
	Type NodeType `json:"T,omitempty"`
	S    string   `json:",omitempty"`
	F    float64  `json:",omitempty"`
	N    []Node   `json:",omitempty"`
}

func Run(args []string) {
	lex := NewLexer(os.Stdin)
	yyParse(lex)
}

func Num(n float64) Node {
	return Node{
		Type: numberNT,
		F:    n,
	}
}

func Str(t NodeType, s string) Node {
	return Node{
		Type: t,
		S:    s,
	}
}
