package vm

import (
	"os"
	"text/template"
	"xoba/ad/defs"
)

const ops_source = "ops.go"

var ops map[string]string = map[string]string{
	"Abs":       "absolute value",
	"Acos":      "",
	"Add":       "",
	"Asin":      "",
	"Atan":      "",
	"Cos":       "",
	"Cosh":      "",
	"Divide":    "",
	"Exp":       "",
	"Exp10":     "10^x",
	"Exp2":      "2^x",
	"Halt":      "",
	"Inputs":    "validate input dimension is large enough",
	"Literal":   "",
	"Log":       "",
	"Log10":     "",
	"Log2":      "",
	"Multiply":  "",
	"Outputs":   "validate output dimension is large enough",
	"Pow":       "",
	"Registers": "1 argument, sets the number of registers",
	"SetOutput": "",
	"Sin":       "",
	"Sinh":      "",
	"Sqrt":      "",
	"Subtract":  "",
	"Tan":       "",
	"Tanh":      "",
}

type Signature string

const (
	One     Signature = "SR,DR"    // source register -> destination register
	Two               = "SR,SR,DR" // source registers -> destination register
	None              = "N/A"      // no arguments
	Integer           = "I"        // an integer
)

var sigs map[string]Signature = map[string]Signature{
	"Abs":             One,
	"Acos":            One,
	"Add":             Two,
	"Asin":            One,
	"Atan":            One,
	"Cos":             One,
	"Cosh":            One,
	"Divide":          Two,
	"Exp":             One,
	"Exp10":           One,
	"Exp2":            One,
	"Halt":            None,
	"HaltIfDmodelNil": None,
	"Inputs":          Integer, // specify dimension of input
	"Literal":         "DR,F",  // destination register, floating-point value
	"Log":             One,
	"Log10":           One,
	"Log2":            One,
	"Multiply":        Two,
	"Outputs":         Integer, // specify dimension of output
	"Pow":             Two,
	"Registers":       Integer, // specify number of registers
	"SetScalarOutput": "SR",    // register to take output from
	"SetVectorOutput": "SR,DI", // source register, destination index
	"Sin":             One,
	"Sinh":            One,
	"Sqrt":            One,
	"Subtract":        Two,
	"Tan":             One,
	"Tanh":            One,
}

func GenOps(args []string) {
	f, err := os.Create(ops_source)
	check(err)
	t := template.Must(template.New(ops_source).Parse(`// autogenerated, do not edit!
package vm

type VmOp uint64

const (
	_ VmOp = iota

{{range $op,$desc := .ops}}{{$op}} // {{$desc}}
{{end}}
)

func (o VmOp) String() string {
	switch o {
{{range $op,$desc := .ops}}case {{$op}}:
return "{{$op}}"
{{end}}}
panic("illegal state")
}

`))

	t.Execute(f, map[string]interface{}{
		"ops": ops,
	})
	f.Close()
	defs.Gofmt(ops_source)
}
