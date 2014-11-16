package main

import (
	"xoba/ad/parser"

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
	add("templates", "run templates", parser.RunTemplates)
	add("parse", "emit formula code", parser.Run)
}
