package expression

import (
	"joern-go/parser/ast/statement"
	"joern-go/parser/ast/types"
)

// A FunctionLiteral node represents a function literal.
type FunctionLiteral struct {
	Type *types.FunctionType       `json:"type"` // function type
	Body *statement.BlockStatement `json:"body"` // function body
}

func (x *FunctionLiteral) Position() int {
	return x.Type.Position()
}

func (x *FunctionLiteral) End() int {
	return x.Body.End()
}

func (*FunctionLiteral) ExpressionNode() {}
