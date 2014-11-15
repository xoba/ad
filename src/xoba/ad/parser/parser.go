// parsing mathematical expressions
package parser

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

//go:generate nex lexer.nex
//go:generate go tool yacc -o parser.y.go parser.y
const (
	numberNT     NodeType = "NUM"
	identifierNT NodeType = "IDENT"
	binaryNT     NodeType = "BIN"
	functionNT   NodeType = "FUNC"
)

type NodeType string

type Node struct {
	Type NodeType `json:"T,omitempty"`
	S    string   `json:",omitempty"`
	F    float64  `json:",omitempty"`
	N    []*Node  `json:",omitempty"`
	name string   `json:",omitempty"`
}

func (n Node) String() string {
	buf, err := json.Marshal(n)
	check(err)
	return string(buf)
}

const formula = `a = 99*(5*x + 55)/6`

func Run(args []string) {
	lex := NewContext(NewLexer(strings.NewReader(formula)))
	yyParse(lex)
	linearize(lex.rhs)
}

var i int

func linearize(rhs *Node) {
	switch rhs.Type {
	case numberNT:
		fmt.Printf("v%d = %f\n", i, rhs.F)
		rhs.name = fmt.Sprintf("v%d", i)
		i++
	case identifierNT:
		fmt.Printf("v%d = %s\n", i, rhs.S)
		rhs.name = fmt.Sprintf("v%d", i)
		i++
	case binaryNT:
		linearize(rhs.N[0])
		linearize(rhs.N[1])
		fmt.Printf("v%d = %s %s %s\n", i, rhs.N[0].name, rhs.S, rhs.N[1].name)
		rhs.name = fmt.Sprintf("v%d", i)
		i++
	}
}

func Function(ident string, args ...*Node) *Node {
	return &Node{
		Type: functionNT,
		S:    ident,
		N:    args,
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

func Binary(op rune, a, b *Node) *Node {
	return &Node{
		Type: binaryNT,
		S:    fmt.Sprintf("%c", op),
		N:    []*Node{a, b},
	}
}

func Negate(a *Node) *Node {
	return Binary('*', Number(-1), a)
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
