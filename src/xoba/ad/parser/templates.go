package parser

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func RunTemplates(args []string) {
	GenTemplates("src/xoba/ad/parser/templates", "src/xoba/ad/parser/templates/gen/gen.go")
}

func GenTemplates(dir, output string) {
	body := new(bytes.Buffer)
	var imports []string
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, dir, nil, parser.AllErrors)
	check(err)
	for _, p := range pkgs {
		for _, f := range p.Files {
			for _, s := range f.Imports {
				imports = append(imports, s.Path.Value)
			}
			for _, s := range f.Decls {
				switch t := s.(type) {

				case *ast.FuncDecl:
					fmt.Fprintf(body, "func %s(", t.Name)
					fmt.Fprint(body, fields(t.Type.Params.List))
					fmt.Fprint(body, ")")
					fmt.Fprint(body, "(")
					fmt.Fprint(body, fields(t.Type.Results.List))
					fmt.Fprint(body, ") {")
					start := t.Body.Lbrace
					file := fset.File(start)
					end := t.Body.Rbrace
					buf, err := ioutil.ReadFile(file.Name())
					check(err)
					fmt.Fprint(body, string(buf[start:end-1]))
					fmt.Fprintln(body, "}\n")
				}
			}
		}
	}
	f, err := os.Create(output)
	check(err)
	fmt.Fprintf(f, "package %s\n", filepath.Base(filepath.Dir(output)))
	fmt.Fprintf(f, "import (\n")
	for _, i := range imports {
		fmt.Fprintf(f, "%s\n", i)
	}
	fmt.Fprintf(f, ")\n")
	f.Write(body.Bytes())
	defer f.Close()
	Gofmt(output)
}

func fields(fields []*ast.Field) string {
	var params []string
	for _, r := range fields {
		var names []string
		for _, i := range r.Names {
			names = append(names, i.Name)
		}
		n := fmt.Sprintf("%s %s", strings.Join(names, ","), r.Type)
		params = append(params, n)
	}
	return strings.Join(params, ",")
}
