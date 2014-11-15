package parser

import (
	"bytes"
	"text/template"
)

func Functions(p string) string {
	t := template.Must(template.New("compute.go").Parse(xfuncs))
	out := new(bytes.Buffer)
	t.Execute(out, map[string]interface{}{
		"private": private,
	})
	return out.String()
}

const xfuncs = `

func {{.private}}_add(a, b float64) float64 {
	return a + b
}
func d_{{.private}}_add(i int, a, b float64) float64 {
	return 1
}

func {{.private}}_multiply(a, b float64) float64 {
	return a * b
}
func d_{{.private}}_multiply(i int, a, b float64) float64 {
	switch i {
	case 0:
		return b
	case 1:
		return a
	default:
		panic("illegal index")
	}
}

func {{.private}}_subtract(a, b float64) float64 {
	return a - b
}
func d_{{.private}}_subtract(i int, a, b float64) float64 {
	switch i {
	case 0:
		return 1
	case 1:
		return -1
	default:
		panic("illegal index")
	}
}

func {{.private}}_divide(a, b float64) float64 {
	return a / b
}
func d_{{.private}}_divide(i int, a, b float64) float64 {
	switch i {
	case 0:
		return 1 / b
	case 1:
		return -a / (b * b)
	default:
		panic("illegal index")
	}
}

func {{.private}}_sqrt(a float64) float64 {
	return math.Sqrt(a)
}
func d_{{.private}}_sqrt(_int, a float64) float64 {
	return 0.5 * math.Pow(a, -1.5)
}

func {{.private}}_exp(a float64) float64 {
	return math.Exp(a)
}
func d_{{.private}}_exp(_ int, a float64) float64 {
	return math.Exp(a)
}

func {{.private}}_log(a float64) float64 {
	return math.Log(a)
}
func d_{{.private}}_log(_ int, a float64) float64 {
	return 1 / a
}

func {{.private}}_pow(a, b float64) float64 {
	return math.Pow(a, b)
}
func d_{{.private}}_pow(i int, a, b float64) float64 {
	switch i {
	case 0:
		return b * math.Pow(a, b-1)
	case 1:
		return math.Pow(a, b) * math.Log(a)
	default:
		panic("illegal index")
	}
}

func log(a float64) float64 {
	return math.Log(a)
}
func exp(a float64) float64 {
	return math.Exp(a)
}


`
