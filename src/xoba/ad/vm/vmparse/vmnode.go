package vmparse

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	F    float64  `json:"F,omitempty"`
	I    int      `json:"I,omitempty"`
	C    []*Node  `json:"C,omitempty"`
}

func (n Node) Name() string {
	switch n.Type {
	case identifierNT, functionNT:
		return n.S
	case indexedIdentifierNT:
		return fmt.Sprintf("%s[%d]", n.S, n.I)
	default:
		panic("illegal type: " + n.Type)
	}
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

func (n *Node) AddChild(c *Node) *Node {
	n.C = append(n.C, c)
	return n
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
