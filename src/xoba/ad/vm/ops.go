package vm

type VmOp uint64

const (
	_ VmOp = iota

	Abs
	Acos
	Add
	Asin
	Atan
	Cos
	Cosh
	Divide
	Exp
	Exp10
	Exp2
	Halt
	HaltIfDmodelNil
	Literal
	Log
	Log10
	Log2
	Multiply
	Pow
	SetScalarOutput
	SetVectorOutput
	Sin
	Sinh
	Sqrt
	Subtract
	Tan
	Tanh
)

func (o VmOp) String() string {
	switch o {
	case Abs:
		return "Abs"
	case Acos:
		return "Acos"
	case Add:
		return "Add"
	case Asin:
		return "Asin"
	case Atan:
		return "Atan"
	case Cos:
		return "Cos"
	case Cosh:
		return "Cosh"
	case Divide:
		return "Divide"
	case Exp:
		return "Exp"
	case Exp10:
		return "Exp10"
	case Exp2:
		return "Exp2"
	case Halt:
		return "Halt"
	case HaltIfDmodelNil:
		return "HaltIfDmodelNil"
	case Literal:
		return "Literal"
	case Log:
		return "Log"
	case Log10:
		return "Log10"
	case Log2:
		return "Log2"
	case Multiply:
		return "Multiply"
	case Pow:
		return "Pow"
	case SetScalarOutput:
		return "SetScalarOutput"
	case SetVectorOutput:
		return "SetVectorOutput"
	case Sin:
		return "Sin"
	case Sinh:
		return "Sinh"
	case Sqrt:
		return "Sqrt"
	case Subtract:
		return "Subtract"
	case Tan:
		return "Tan"
	case Tanh:
		return "Tanh"
	}
	panic("illegal state")
}

// the following can be copied over to "ops" var (self-reproduction)
/*
var ops []string = []string{
"Abs",
"Acos",
"Add",
"Asin",
"Atan",
"Cos",
"Cosh",
"Divide",
"Exp",
"Exp10",
"Exp2",
"Halt",
"HaltIfDmodelNil",
"Literal",
"Log",
"Log10",
"Log2",
"Multiply",
"Pow",
"SetScalarOutput",
"SetVectorOutput",
"Sin",
"Sinh",
"Sqrt",
"Subtract",
"Tan",
"Tanh",
}
*/

// the following can serve as template for single-arg funcs
/*

	twos := map[VmOp]string{
Abs:"math.Abs",
Acos:"math.Acos",
Add:"math.Add",
Asin:"math.Asin",
Atan:"math.Atan",
Cos:"math.Cos",
Cosh:"math.Cosh",
Divide:"math.Divide",
Exp:"math.Exp",
Exp10:"math.Exp10",
Exp2:"math.Exp2",
Halt:"math.Halt",
HaltIfDmodelNil:"math.HaltIfDmodelNil",
Literal:"math.Literal",
Log:"math.Log",
Log10:"math.Log10",
Log2:"math.Log2",
Multiply:"math.Multiply",
Pow:"math.Pow",
SetScalarOutput:"math.SetScalarOutput",
SetVectorOutput:"math.SetVectorOutput",
Sin:"math.Sin",
Sinh:"math.Sinh",
Sqrt:"math.Sqrt",
Subtract:"math.Subtract",
Tan:"math.Tan",
Tanh:"math.Tanh",

}
*/
