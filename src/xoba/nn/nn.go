// a simple neural network example using automatic differentiation
package nn

import (
	"flag"
	"fmt"
	"os"
)

//go:generate run nn -gen
//go:generate run compile -formula=nn.txt -output nn_ad.go -templates "../ad/parser/templates" -package nn -main=false -time=false
func Run(args []string) {
	var gen bool
	var nclass, hidden, inputs int
	flags := flag.NewFlagSet("parse", flag.ExitOnError)
	flags.BoolVar(&gen, "gen", false, "whether to generate formula")
	flags.IntVar(&inputs, "inputs", 2, "number of inputs")
	flags.IntVar(&nclass, "classes", 3, "number of classes")
	flags.IntVar(&hidden, "hidden", 3, "number of hidden units")
	flags.Parse(args)
	if gen {
		fmt.Println("generating formula")
		f, err := os.Create("nn.txt")
		check(err)
		defer f.Close()
		fmt.Fprintf(f, "f := a+b\n")
		return
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
