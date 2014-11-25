package vm

import (
	"os"
	"strings"
	"text/template"
	"xoba/ad/defs"
)

const source = "ops.go"

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
		"ops": strings.Split("Add,Divide,Halt,HaltIfDmodelNil,Literal,Multiply,SetScalarOutput,SetVectorOutput,Subtract", ","),
	})
	f.Close()
	defs.Gofmt(source)
}
