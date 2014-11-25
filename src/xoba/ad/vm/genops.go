package vm

import (
	"os"
	"sort"
	"strings"
	"text/template"
	"xoba/ad/defs"
)

const ops_source = "ops.go"

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
	"Exp10",
	"Exp2",
	"Halt",
	"HaltIfDmodelNil",
	"Literal",
	"Log",
	"Log10",
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
	f, err := os.Create(ops_source)
	check(err)
	t := template.Must(template.New(ops_source).Parse(`// autogenerated, do not edit!
package vm

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

// the following can serve as template for single-arg funcs
/*

	twos := map[VmOp]string{
{{range .ops}}{{.}}:"math.{{.}}",
{{end}}
}
*/

`))
	t.Execute(f, map[string]interface{}{
		"ops": ops,
	})
	f.Close()
	defs.Gofmt(ops_source)
}

func init() {
	// add to and re-sort the ops list
	var additional = "log10,exp10"
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
