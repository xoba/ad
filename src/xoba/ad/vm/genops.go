package vm

import (
	"os"
	"sort"
	"text/template"
	"xoba/ad/defs"
)

const source = "ops.go"

var ops = []string{
	"Add",
	"Divide",
	"Halt",
	"HaltIfDmodelNil",
	"Literal",
	"Multiply",
	"SetScalarOutput",
	"SetVectorOutput",
	"Subtract",
}

func init() {
	sort.Strings(ops)
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
`))
	t.Execute(f, map[string]interface{}{
		"ops": ops,
	})
	f.Close()
	defs.Gofmt(source)
}
