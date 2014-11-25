package vm

type VmOp uint64

const (
	_ VmOp = iota

	Add
	Divide
	Halt
	HaltIfDmodelNil
	Literal
	Multiply
	SetScalarOutput
	SetVectorOutput
	Subtract
)

func (o VmOp) String() string {
	switch o {
	case Add:
		return "Add"
	case Divide:
		return "Divide"
	case Halt:
		return "Halt"
	case HaltIfDmodelNil:
		return "HaltIfDmodelNil"
	case Literal:
		return "Literal"
	case Multiply:
		return "Multiply"
	case SetScalarOutput:
		return "SetScalarOutput"
	case SetVectorOutput:
		return "SetVectorOutput"
	case Subtract:
		return "Subtract"
	}
	panic("illegal state")
}
