package types

import (
	"joern-go/parser/ast"
	"joern-go/parser/ast/expression"
)

// An ArrayType node represents an array or slice type.
type ArrayType struct {
	ast.Node
	LeftBracket int                   `json:"leftBracket"` // position of "["
	Length      expression.Expression `json:"length"`      // Ellipsis node for [...]T array types, nil for slice types
	Element     expression.Expression `json:"element"`     // element type
}

func (x ArrayType) Start() int {
	return x.LeftBracket
}

func (x ArrayType) End() int {
	return x.Element.End()
}

func (ArrayType) ExpressionNode() {}
