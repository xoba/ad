package vm

import (
	"os"
	"sort"
	"strings"
	"text/template"
	"xoba/ad/defs"
)

const source = "ops.go"

var ops []string = []string{
	"Abs",
	"Acos",
	"Add",
	"Asin",
	"Atan",
	"Cos",
	"Cosh",
	"Divide",
	"Exp",
	"Exp2",
	"Halt",
	"HaltIfDmodelNil",
	"Literal",
	"Log",
	"Log2",
	"Multiply",
	"Pow",
	"SetScalarOutput",
	"SetVectorOutput",
	"Sin",
	"Sinh",
	"Sqrt",
	"Subtract",
	"Tan",
	"Tanh",
}

func GenOps(args []string) {
	f, err := os.Create(source)
	check(err)
	t := template.Must(template.New("output.go").Parse(`package vm

type VmOp uint64

const (
	_ VmOp = iota

{{range .ops}}{{.}}
{{end}}
)

func (o VmOp) String() string {
	switch o {
{{range .ops}}case {{.}}:
return "{{.}}"
{{end}}}
panic("illegal state")
}

// the following can be copied over to "ops" var (self-reproduction)
/*
var ops []string = []string{
{{range .ops}}"{{.}}",
{{end}}}
*/
`))
	t.Execute(f, map[string]interface{}{
		"ops": ops,
	})
	f.Close()
	defs.Gofmt(source)
}

func init() {
	// add to and re-sort the ops list
	var additional = ""
	ops = append(ops, strings.Split(additional, ",")...)
	opsMap := make(map[string]bool)
	for _, o := range ops {
		opsMap[strings.Title(o)] = true
	}
	ops = make([]string, 0)
	for o := range opsMap {
		if len(o) > 0 {
			ops = append(ops, o)
		}
	}
	sort.Strings(ops)
}
