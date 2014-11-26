// vm for running evaluation and autodiff code
package vm

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

// e.g., if "model" has dim n, then first n dims of "out" could be gradient, dim n+1 can be scalar output
type Executor func(p Program, model, in, out []float64) (err error)

//go:generate run genops
//go:generate run genvm
func Run(args []string) {
	var e Executor = Execute
	var asm string
	flags := flag.NewFlagSet("vm", flag.ExitOnError)
	flags.StringVar(&asm, "asm", "lib/test.asm", "assmebly file to run")
	flags.Parse(args)
	f, err := os.Open(asm)
	check(err)
	defer f.Close()
	p := Compile(f)
	out := make([]float64, p.Outputs)
	in := make([]float64, p.Inputs)
	model := make([]float64, p.Models)
	in[0] = 99
	model[0] = -3.14
	check(e(p, model, in, out))
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
	Name      string `json:",omitempty"`
	Registers uint64 `json:",omitempty"`
	Inputs    uint64 `json:",omitempty"`
	Models    uint64 `json:",omitempty"`
	Outputs   uint64 `json:",omitempty"`
	Code      []byte
}

func (p Program) String() string {
	buf, _ := json.MarshalIndent(p, "", "  ")
	return string(buf)
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
