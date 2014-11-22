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

const genTests = false

// generate function templates, output imports and actual code
func GenTemplates(dir, private string) ([]string, string, error) {
	imports := make(map[string]struct{})
	body := new(bytes.Buffer)
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, dir, nil, parser.AllErrors)
	if err != nil {
		return nil, "", err
	}
	for _, p := range pkgs {
		for n, f := range p.Files {
			if strings.HasSuffix(n, "_test.go") {
				continue
			}
			for _, s := range f.Imports {
				imports[s.Path.Value] = struct{}{}
			}
			for _, s := range f.Decls {
				switch t := s.(type) {
				case *ast.FuncDecl:
					output := func(name string) error {
						fmt.Fprintf(body, "func %s(", name)
						fmt.Fprint(body, fields(t.Type.Params.List))
						fmt.Fprint(body, ")")
						if t.Type.Results != nil {
							fmt.Fprint(body, "(")
							fmt.Fprint(body, fields(t.Type.Results.List))
							fmt.Fprint(body, ")")
						}
						start := t.Body.Lbrace
						file := fset.File(start)
						end := t.Body.Rbrace
						buf, err := ioutil.ReadFile(file.Name())
						if err != nil {
							return err
						}
						p0 := fset.Position(start).Offset
						p1 := fset.Position(end).Offset
						fmt.Fprint(body, string(buf[p0:p1]))
						fmt.Fprintln(body, "}\n")
						return nil
					}
					if err := output(t.Name.Name); err != nil {
						return nil, "", err
					}
					if err := output(fmt.Sprintf("%s_%s", t.Name.Name, private)); err != nil {
						return nil, "", err
					}
				}
			}
		}
	}
	var out []string
	for k := range imports {
		out = append(out, k)
	}
	return out, body.String(), nil
}

func count(fields []*ast.Field) (out int) {
	for _, r := range fields {
		for range r.Names {
			out++
		}
	}
	return
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
