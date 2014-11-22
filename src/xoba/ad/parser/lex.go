package parser

import (
	"encoding/json"
	"fmt"
	"strconv"
)

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
}

func (n Node) String() string {
	buf, _ := json.Marshal(n)
	return string(buf)
}

func (n Node) VarName() string {
	switch n.Type {
	case identifierNT:
		return n.S
	case indexedIdentifierNT:
		return fmt.Sprintf("%s[%d]", n.S, n.I)
	default:
		panic(fmt.Sprintf("illegal type: %d", n.Type))
	}
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
