package vm

import (
	"os"
	"text/template"
	"xoba/ad/defs"
)

const vm_source = "execute.go"

func GenVm(args []string) {

	twos := map[VmOp]string{
		Abs:   "math.Abs",
		Acos:  "math.Acos",
		Asin:  "math.Asin",
		Atan:  "math.Atan",
		Cos:   "math.Cos",
		Cosh:  "math.Cosh",
		Exp:   "math.Exp",
		Exp10: "exp10",
		Exp2:  "math.Exp2",
		Log:   "math.Log",
		Log10: "math.Log10",
		Log2:  "math.Log2",
		Sin:   "math.Sin",
		Sinh:  "math.Sinh",
		Sqrt:  "math.Sqrt",
		Tan:   "math.Tan",
		Tanh:  "math.Tanh",
	}

	threes := map[VmOp]string{
		Add:      "+",
		Multiply: "*",
		Divide:   "/",
		Subtract: "-",
	}

	f, err := os.Create(vm_source)
	check(err)
	t := template.Must(template.New(vm_source).Parse(`package vm

import (
"math"
	"bytes"
	"encoding/binary"
"fmt"
)

func Execute(p Program, x, model, dmodel []float64) (y float64, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered from: %v", r)
		}
	}()
	fmt.Printf("%d byte program %q = 0x%x\n", len(p.Code), p.Name, p.Code)
	r := bytes.NewReader(p.Code)
	one := func() uint64 {
		a, err := binary.ReadUvarint(r)
		check(err)
		return a
	}
	two := func() (uint64, uint64) {
		a, err := binary.ReadUvarint(r)
		check(err)
		b, err := binary.ReadUvarint(r)
		check(err)
		return a, b
	}
	three := func() (uint64, uint64, uint64) {
		a, err := binary.ReadUvarint(r)
		check(err)
		b, err := binary.ReadUvarint(r)
		check(err)
		c, err := binary.ReadUvarint(r)
		check(err)
		return a, b, c
	}
	registers := make([]float64, p.Registers)

Loop:
	for {
		c, err := binary.ReadUvarint(r)
		check(err)
		fmt.Printf("op = %s\n", VmOp(c))
		// general rules:
		// locations stored first, values later
		// source first, destination after
		switch VmOp(c) {
		case Literal: // store a literal to register
			loc, err := binary.ReadUvarint(r)
			check(err)
			var lit float64
			binary.Read(r, order, &lit)
			registers[loc] = lit
		case SetScalarOutput: // set output from register
			y = registers[one()]
		case SetVectorOutput: // set output from register
			src, dest := two()
			dmodel[dest] = registers[src]
		case HaltIfDmodelNil:
			if dmodel == nil {
				break Loop
			}
		case Halt:
			break Loop
{{range $name,$func := .twos}}case {{$name}}:
			src, dest := two()
			registers[dest] = {{$func}}(registers[src])
{{end}} {{range $name,$op := .threes}}case {{$name}}:
			a, b, dest := three()
			registers[dest] = registers[a] {{$op}} registers[b]
{{end}} 	default:
			return 0, fmt.Errorf("unhandled op %s", VmOp(c))
		}
	}
	return
}

`))
	t.Execute(f, map[string]interface{}{
		"twos":   twos,
		"threes": threes,
	})
	f.Close()
	defs.Gofmt(vm_source)
}
