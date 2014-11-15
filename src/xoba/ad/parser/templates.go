package parser

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"strings"
)

func RunTemplates(args []string) {
	fmt.Println(GenTemplates("src/xoba/ad/parser/templates", "abc"))
}

func GenTemplates(dir, private string) (imports map[string]struct{}, code string) {
	imports = make(map[string]struct{})
	body := new(bytes.Buffer)
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, dir, nil, parser.AllErrors)
	check(err)
	for _, p := range pkgs {
		for _, f := range p.Files {
			for _, s := range f.Imports {
				imports[s.Path.Value] = struct{}{}
			}
			for _, s := range f.Decls {
				switch t := s.(type) {
				case *ast.FuncDecl:
					output := func(name string) {
						fmt.Fprintf(body, "func %s(", name)
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
					output(t.Name.Name)
					output(fmt.Sprintf("%s_%s", t.Name.Name, private))
				}
			}
		}
	}
	return imports, body.String()
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
