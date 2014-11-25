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
//go:generate run genvm
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
		case "log":
			putOp(Log)
			putInt(1)
			putInt(2)
		case "log10":
			putOp(Log10)
			putInt(1)
			putInt(2)
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

func check(e error) {
	if e != nil {
		panic(e)
	}
}
