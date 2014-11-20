package main

import (
	"xoba/ad/parser"
	"xoba/nn"

	"github.com/xoba/goutil"
	"github.com/xoba/goutil/tool"
)

func main() {
	tool.Run()
}

func init() {
	goutil.PlatformInit()
	add := func(name, desc string, f func([]string)) {
		tool.Register(tool.Named(name+","+desc, tool.RunFunc(f)))
	}
	add("compile", "compile formula to go", parser.Run)
	add("nn", "neural network example", nn.Run)
}
