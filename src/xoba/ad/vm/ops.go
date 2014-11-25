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
