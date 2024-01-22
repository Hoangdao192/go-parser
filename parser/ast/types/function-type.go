package types

import "joern-go/parser/ast"

type FunctionType struct {
	Function   int            // position of "func" keyword (token.NoPos if there is no "func")
	TypeParams *ast.FieldList // type parameters; or nil
	Params     *ast.FieldList // (incoming) parameters; non-nil
	Results    *ast.FieldList // (outgoing) results; or nil
}

func (x *FunctionType) Position() int {
	if x.Function == 0 || x.Params == nil { // see issue 3870
		return x.Function
	}
	return x.Params.Position() // interface method declarations have no "func" keyword
}

func (x *FunctionType) End() int {
	if x.Results != nil {
		return x.Results.End()
	}
	return x.Params.End()
}

func (*FunctionType) ExpressionNode() {}
