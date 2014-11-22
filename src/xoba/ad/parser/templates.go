package parser

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
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
	for pname, p := range pkgs {
		fmt.Printf("pkg = %q\n", pname)
		for n, f := range p.Files {
			if strings.HasSuffix(n, "_test.go") {
				log.Printf("skipping %s\n", n)
				continue
			}
			log.Printf("parsing %s\n", n)
			for _, s := range f.Imports {
				imports[s.Path.Value] = struct{}{}
			}
			for _, s := range f.Decls {
				switch t := s.(type) {
				case *ast.FuncDecl:
					output := func(name string, testing bool) error {
						fmt.Printf("name = %q\n", name)
						if genTests && testing && !strings.HasPrefix(name, "d_") {
							switch count(t.Type.Params.List) {
							case 1:
								fmt.Printf("test1d(%q,%s,d_%s,t)\n", name, name, name)
							case 2:
								fmt.Printf("test2d(%q,%s,d_%s,t)\n", name, name, name)
							}
						}
						fmt.Fprintf(body, "func %s(", name)
						fmt.Fprint(body, fields(t.Type.Params.List))
						fmt.Fprint(body, ")")
						if t.Type.Results != nil {
							fmt.Fprint(body, "(")
							fmt.Fprint(body, fields(t.Type.Results.List))
							fmt.Fprint(body, ")")
						}
						fmt.Fprint(body, "{")
						start := t.Body.Lbrace
						file := fset.File(start)
						end := t.Body.Rbrace
						buf, err := ioutil.ReadFile(file.Name())
						if err != nil {
							return err
						}
						fmt.Printf("%s (%d bytes) %d to %d\n", file.Name(), len(buf), start, end)
						fmt.Fprint(body, string(buf[start:end-1]))
						fmt.Fprintln(body, "}\n")
						return nil
					}
					if err := output(t.Name.Name, true); err != nil {
						return nil, "", err
					}
					if err := output(fmt.Sprintf("%s_%s", t.Name.Name, private), false); err != nil {
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
