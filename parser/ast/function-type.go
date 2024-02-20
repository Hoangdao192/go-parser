package ast

type FunctionType struct {
	Node
	Function   int        // position of "func" keyword (token.NoPos if there is no "func")
	TypeParams *FieldList // type parameters; or nil
	Params     *FieldList // (incoming) parameters; non-nil
	Results    *FieldList // (outgoing) results; or nil
}

func (x *FunctionType) Start() int {
	if x.Function == 0 { // see issue 3870
		return x.Function
	}
	return x.Params.Start() // interface method declarations have no "func" keyword
}

func (x *FunctionType) End() int {
	if x.Results.Start() != 0 {
		return x.Results.End()
	}
	return x.Params.End()
}

func (*FunctionType) ExpressionNode() {}
