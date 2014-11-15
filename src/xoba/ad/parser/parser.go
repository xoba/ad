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
	N    []Node   `json:",omitempty"`
}

func (n Node) String() string {
	buf, err := json.Marshal(n)
	check(err)
	return string(buf)
}

const formula = `a = f(5*x + 55)`

func Run(args []string) {
	lex := NewContext(NewLexer(strings.NewReader(formula)))
	yyParse(lex)
	fmt.Printf("lhs = %s\n", lex.lhs)
	fmt.Printf("rhs = %s\n", lex.rhs)
}

func Function(ident string, args ...Node) Node {
	return Node{
		Type: functionNT,
		S:    ident,
		N:    args,
	}
}

func Number(n float64) Node {
	return Node{
		Type: numberNT,
		F:    n,
	}
}

func LexIdentifier(s string) Node {
	return Node{
		Type: identifierNT,
		S:    s,
	}
}

func LexNumber(s string) Node {
	n, err := strconv.ParseFloat(s, 64)
	check(err)
	return Number(n)
}

func Binary(op rune, a, b Node) Node {
	return Node{
		Type: binaryNT,
		S:    fmt.Sprintf("%c", op),
		N:    []Node{a, b},
	}
}

func Negate(a Node) Node {
	return Binary('*', Number(-1), a)
}

type context struct {
	lhs Node
	rhs Node
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
