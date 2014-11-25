package vm

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
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
		case Abs:
			src, dest := two()
			registers[dest] = math.Abs(registers[src])
		case Acos:
			src, dest := two()
			registers[dest] = math.Acos(registers[src])
		case Asin:
			src, dest := two()
			registers[dest] = math.Asin(registers[src])
		case Atan:
			src, dest := two()
			registers[dest] = math.Atan(registers[src])
		case Cos:
			src, dest := two()
			registers[dest] = math.Cos(registers[src])
		case Cosh:
			src, dest := two()
			registers[dest] = math.Cosh(registers[src])
		case Exp:
			src, dest := two()
			registers[dest] = math.Exp(registers[src])
		case Exp10:
			src, dest := two()
			registers[dest] = exp10(registers[src])
		case Exp2:
			src, dest := two()
			registers[dest] = math.Exp2(registers[src])
		case Log:
			src, dest := two()
			registers[dest] = math.Log(registers[src])
		case Log10:
			src, dest := two()
			registers[dest] = math.Log10(registers[src])
		case Log2:
			src, dest := two()
			registers[dest] = math.Log2(registers[src])
		case Sin:
			src, dest := two()
			registers[dest] = math.Sin(registers[src])
		case Sinh:
			src, dest := two()
			registers[dest] = math.Sinh(registers[src])
		case Sqrt:
			src, dest := two()
			registers[dest] = math.Sqrt(registers[src])
		case Tan:
			src, dest := two()
			registers[dest] = math.Tan(registers[src])
		case Tanh:
			src, dest := two()
			registers[dest] = math.Tanh(registers[src])
		case Add:
			a, b, dest := three()
			registers[dest] = registers[a] + registers[b]
		case Divide:
			a, b, dest := three()
			registers[dest] = registers[a] / registers[b]
		case Multiply:
			a, b, dest := three()
			registers[dest] = registers[a] * registers[b]
		case Subtract:
			a, b, dest := three()
			registers[dest] = registers[a] - registers[b]
		default:
			return 0, fmt.Errorf("unhandled op %s", VmOp(c))
		}
	}
	return
}
