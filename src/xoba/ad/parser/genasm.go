package parser

import (
	"fmt"
	"io"
	"strings"
)

// todo:
// multiple formulas in parser
// lexer/parser takes interface, not struct
// assembly compiler outputs a map of variable names to assignments (input, output, registers)
// assembly compiler can't really know about model variables... just inputs?
//
//

func GenerateVmAssembly(formula string, w io.Writer) error {
	lex := NewContext(NewLexer(strings.NewReader(formula)))
	yyParse(lex)

	vp := &VarProcessor{
		inputs: make(map[string]int),
	}

	vp.getVars(lex.rhs)
	fmt.Fprintf(w, "# compute value and gradient of %s\n", formula)
	fmt.Fprintf(w, "# inputs = %v\n", vp.inputs)
	fmt.Fprintf(w, "inputs %d\n", len(vp.inputs))
	fmt.Fprintf(w, "outputs %d\n", 1+len(vp.inputs))

	//	ep := &EvalProcessor{}

	return nil
}

type VarProcessor struct {
	inputs map[string]int
}

func (v *VarProcessor) getVars(root *Node) {
	switch root.Type {
	case identifierNT, indexedIdentifierNT:
		if _, ok := v.inputs[root.S]; !ok {
			v.inputs[root.S] = len(v.inputs)
		}
	case functionNT:
		for _, n := range root.Children {
			v.getVars(n)
		}
	}
	return
}

type AsmLine struct {
	lhs int
	rhs int
}

type EvalProcessor struct {
	registers map[int]string
}

func (s *EvalProcessor) program(rhs *Node, private string) (out []AsmLine) {
	switch rhs.Type {
	case numberNT:
		rhs.Name = fmt.Sprintf("%f", rhs.F)
	case functionNT:
	}
	return out
}
