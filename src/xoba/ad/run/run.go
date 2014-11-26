package main

import (
	"xoba/ad/parser"
	"xoba/ad/vm"
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
	add("vm", "play with vm code", vm.Run)
	add("genops", "generate vm ops definitions", vm.GenOps)
	add("orgops", "generate vm ops definitions", vm.OrganizeOps)
	add("genvm", "generate vm", vm.GenVm)
}
