// vm for running evaluation and autodiff code
package vm

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

//go:generate run genops
func Run(args []string) {
	var asm string
	flags := flag.NewFlagSet("vm", flag.ExitOnError)
	flags.StringVar(&asm, "asm", "lib/test.asm", "assmebly file to run")
	flags.Parse(args)
	f, err := os.Open(asm)
	check(err)
	defer f.Close()
	w := new(bytes.Buffer)
	p := Program{
		Name:      asm,
		Registers: 10,
	}
	s := bufio.NewScanner(f)
	var fields []string
	tmp := make([]byte, 20)
	putOp := func(o VmOp) {
		n := binary.PutUvarint(tmp, uint64(o))
		w.Write(tmp[:n])
	}
	putInt := func(i int) {
		v, err := strconv.ParseUint(fields[i], 10, 64)
		check(err)
		n := binary.PutUvarint(tmp, v)
		w.Write(tmp[:n])
	}
	putFloat := func(i int) {
		v, err := strconv.ParseFloat(fields[i], 64)
		check(err)
		binary.Write(w, order, v)
	}
	for s.Scan() {
		line := s.Text()
		line = strings.TrimSpace(strings.ToLower(line))
		if len(line) == 0 || line[0] == '#' {
			continue
		}
		fields = strings.Fields(line)
		switch fields[0] {
		case "halt":
			putOp(Halt)
		case "haltifdmodelnil":
			putOp(HaltIfDmodelNil)
		case "literal":
			putOp(Literal)
			putInt(1)
			putFloat(2)
		case "setscalaroutput":
			putOp(SetScalarOutput)
			putInt(1)
		case "setvectoroutput":
			putOp(SetVectorOutput)
			putInt(1)
			putInt(2)
		case "multiply":
			putOp(Multiply)
			putInt(1)
			putInt(2)
			putInt(3)
		case "divide":
			putOp(Divide)
			putInt(1)
			putInt(2)
			putInt(3)
		case "subtract":
			putOp(Subtract)
			putInt(1)
			putInt(2)
			putInt(3)
		case "add":
			putOp(Add)
			putInt(1)
			putInt(2)
			putInt(3)
		default:
			log.Fatalf("unknown opcode: %s", fields[0])
		}
	}
	check(s.Err())
	p.Code = w.Bytes()
	var x, model, dmodel []float64
	y, err := Execute(p, x, model, dmodel)
	check(err)
	fmt.Println(y)
}

func bytesToFloat(buf []byte) (f float64) {
	binary.Read(bytes.NewReader(buf), binary.BigEndian, &f)
	return
}

func floatToBytes(f float64) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, f)
	return buf.Bytes()
}

func bytesToInt(buf []byte) (f uint64) {
	binary.Read(bytes.NewReader(buf), binary.BigEndian, &f)
	return
}

func intToBytes(f uint64) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, f)
	return buf.Bytes()
}

type Program struct {
	Name      string
	Registers int
	Code      []byte
}

var (
	order          = binary.BigEndian
	DimensionError = fmt.Errorf("dimension mismatch")
)

// if dmodel == nil, don't calculate gradient
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
		case Add:
			a, b, dest := three()
			registers[dest] = registers[a] + registers[b]
		case Subtract:
			a, b, dest := three()
			registers[dest] = registers[a] - registers[b]
		case Divide:
			a, b, dest := three()
			registers[dest] = registers[a] / registers[b]
		case Multiply:
			a, b, dest := three()
			registers[dest] = registers[a] * registers[b]
		default:
			return 0, fmt.Errorf("unhandled op %s", VmOp(c))
		}
	}
	return
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
