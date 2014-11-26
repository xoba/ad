// vm for running evaluation and autodiff code
package vm

import (
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"os"
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
	p := Compile(f)
	p.Registers = 10
	var x, model, out []float64
	out = make([]float64, 10)
	check(Execute(p, x, model, out))
	fmt.Println(out)
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
