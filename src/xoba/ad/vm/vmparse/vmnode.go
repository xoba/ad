package vmparse

import (
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

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
	C    []*Node  `json:"C,omitempty"`
}

func (n Node) Float64() float64 {
	f, err := strconv.ParseFloat(n.S, 64)
	check(err)
	return f
}

func (n Node) IndexedVar() (string, int) {
	if n.Type != indexedIdentifierNT {
		panic("illegal type: " + n.Type)
	}
	return parseIndex(n.S)
}

func parseIndex(s string) (string, int) {
	p := regexp.MustCompile("^([a-zA-Z]+[a-zA-Z0-9_]*)\\[(\\d+)\\]$")
	x := p.FindStringSubmatch(s)
	if len(x) != 3 {
		panic("can't match " + s)
	}
	i, err := strconv.ParseUint(x[2], 10, 64)
	check(err)
	return x[1], int(i)
}

func (n Node) Name() string {
	return n.S
}

func (n *Node) DeepCopy() *Node {
	out := &Node{
		Type: n.Type,
		S:    n.S,
	}
	for _, c := range n.C {
		out.C = append(out.C, c.DeepCopy())
	}
	return out
}

func (n *Node) CopyFrom(o *Node) {
	n.Type = o.Type
	n.S = o.S
	n.C = o.C
}

func (n Node) String() string {
	buf, _ := json.Marshal(n)
	return string(buf)
}

func (n *Node) AddChild(c *Node) *Node {
	n.C = append(n.C, c)
	return n
}

func (n Node) Formula() string {
	buf := new(bytes.Buffer)
	switch n.Type {
	case numberNT:
		return n.S
	case identifierNT:
		return n.S
	case indexedIdentifierNT:
		return n.S
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

func check(e error) {
	if e != nil {
		panic(e)
	}
}
