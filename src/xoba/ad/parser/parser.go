// parsing mathematical expressions
package parser

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"
	"text/template"
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

func Run(args []string) {
	lex := NewContext(NewLexer(strings.NewReader(formula)))
	yyParse(lex)
	pgm := new(bytes.Buffer)
	vars := make(map[string]string)
	vp := &VarParser{vars: vars}
	vp.getVars(pgm, lex.rhs)
	sp := &StepParser{i: 1, vars: vars}
	y := sp.linearize(pgm, lex.rhs)
	var list []string
	for v := range vars {
		list = append(list, v)
	}
	sort.Strings(list)
	f, err := os.Create("compute.go")
	check(err)

	decls := new(bytes.Buffer)
	for _, v := range list {
		fmt.Fprintf(decls, "%s := rand.Float64();\n", v)
	}

	t := template.Must(template.New("compute.go").Parse(`package main
import (
"fmt"
"math"
"time"
"math/rand"
)
func main() {
rand.Seed(time.Now().UTC().UnixNano())
{{.decls}} {{.formula}} 
fmt.Printf("formula: %f\n",{{.lhs}});
fmt.Printf("parsed : %f\n",Compute({{.vars}}))
fmt.Printf("diff   : %f\n",{{.lhs}}-Compute({{.vars}}))
}

func Compute({{.vars}} float64) float64 {
{{.program}} return {{.y}};
}

func sqrt(a float64) float64 {
return math.Sqrt(a);
}

func pow(a,b float64) float64 {
return math.Pow(a,b)
}
`))

	t.Execute(f, map[string]interface{}{
		"decls":   decls.String(),
		"formula": formula,
		"lhs":     lex.lhs.S,
		"vars":    strings.Join(list, ", "),
		"program": pgm.String(),
		"y":       y,
	})

	f.Close()
	cmd := exec.Command("gofmt", "-w", "compute.go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("oops: %v\n", err)
	}
}

const formula = `a := x+sqrt(pow(z,3)) + 99*(5*x + 55)/6 + y`

type VarParser struct {
	i    int
	vars map[string]string
}

// gets all the vars and names them in tree
func (v *VarParser) getVars(w io.Writer, rhs *Node) {
	switch rhs.Type {
	case identifierNT:
		if _, ok := v.vars[rhs.S]; !ok {
			rhs.name = fmt.Sprintf("v_%d", v.i)
			v.vars[rhs.S] = rhs.name
			v.i++
			fmt.Fprintf(w, "%s := %s\n", rhs.name, rhs.S)
		} else {
			rhs.name = v.vars[rhs.S]
		}
	case functionNT, binaryNT:
		for _, n := range rhs.N {
			v.getVars(w, n)
		}
	}
}

type StepParser struct {
	i    int
	vars map[string]string
}

func (s *StepParser) linearize(w io.Writer, rhs *Node) string {
	switch rhs.Type {
	case numberNT:
		rhs.name = fmt.Sprintf("%f", rhs.F)
	case functionNT:
		var args []string
		for _, n := range rhs.N {
			s.linearize(w, n)
			args = append(args, n.name)
		}
		fmt.Fprintf(w, "v%d := %s (%s)\n", s.i, rhs.S, strings.Join(args, ","))
		rhs.name = fmt.Sprintf("v%d", s.i)
		s.i++
	case binaryNT:
		for _, n := range rhs.N {
			s.linearize(w, n)
		}
		fmt.Fprintf(w, "v%d := %s %s %s\n", s.i, rhs.N[0].name, rhs.S, rhs.N[1].name)
		rhs.name = fmt.Sprintf("v%d", s.i)
		s.i++
	}
	return rhs.name
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
