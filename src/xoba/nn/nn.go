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
	flags := flag.NewFlagSet("parse", flag.ExitOnError)
	flags.BoolVar(&gen, "gen", false, "whether to generate formula")
	flags.Parse(args)
	fmt.Println("generating formula")
	f, err := os.Create("nn.txt")
	check(err)
	defer f.Close()
	fmt.Fprintf(f, "f := a+b\n")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
