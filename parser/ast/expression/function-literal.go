package expression

import (
	"joern-go/parser/ast"
	"joern-go/parser/ast/statement"
	"joern-go/parser/ast/types"
)

// A FunctionLiteral node represents a function literal.
type FunctionLiteral struct {
	ast.Node
	Type types.FunctionType       `json:"type"` // function type
	Body statement.BlockStatement `json:"body"` // function body
}

func (x FunctionLiteral) Start() int {
	return x.Type.Start()
}

func (x FunctionLiteral) End() int {
	return x.Body.End()
}

func (FunctionLiteral) ExpressionNode() {}
