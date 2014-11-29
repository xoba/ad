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
	T NodeType `json:"T,omitempty"` // type
	S string   `json:"S,omitempty"` // string value
	C []*Node  `json:"C,omitempty"` // children

	Step Step
}

func (n Node) Float64() float64 {
	if n.T != numberNT {
		panic("illegal type: " + n.T)
	}
	f, err := strconv.ParseFloat(n.S, 64)
	check(err)
	return f
}

func (n Node) IndexedVar() (string, int) {
	if n.T != indexedIdentifierNT {
		panic("illegal type: " + n.T)
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

func (n *Node) DeepCopy() *Node {
	out := &Node{
		T: n.T,
		S: n.S,
	}
	for _, c := range n.C {
		out.C = append(out.C, c.DeepCopy())
	}
	return out
}

func (n *Node) CopyFrom(o *Node) {
	n.T = o.T
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
	return n.formula(true)
}

func (n Node) formula(top bool) string {
	buf := new(bytes.Buffer)
	switch n.T {
	case numberNT:
		return n.S
	case identifierNT:
		return n.S
	case indexedIdentifierNT:
		return n.S
	case functionNT:
		op := func(x string) {
			fmt.Fprintf(buf, "%s %s %s", n.C[0].formula(false), x, n.C[1].formula(false))
		}
		opParen := func(x string) {
			if top {
				op(x)
				return
			}
			fmt.Fprintf(buf, "(%s %s %s)", n.C[0].formula(false), x, n.C[1].formula(false))
		}
		switch n.S {
		case "multiply":
			op("*")
		case "divide":
			op("/")
		case "subtract":
			opParen("-")
		case "add":
			opParen("+")
		default:
			var args []string
			for _, c := range n.C {
				args = append(args, c.formula(false))
			}
			fmt.Fprintf(buf, "%s(%s)", n.S, strings.Join(args, ", "))
		}
	case statementNT:
		fmt.Fprintf(buf, "%s := %s", n.C[0].formula(top), n.C[1].formula(top))
	default:
		panic("illegal type: " + n.T)
	}
	return buf.String()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
