package types

import "joern-go/parser/ast"

// An ArrayType node represents an array or slice type.
type ArrayType struct {
	LeftBracket int            `json:"leftBracket"` // position of "["
	Length      ast.Expression `json:"length"`      // Ellipsis node for [...]T array types, nil for slice types
	Element     ast.Expression `json:"element"`     // element type
}

func (x *ArrayType) Position() int {
	return x.LeftBracket
}

func (x *ArrayType) End() int {
	return x.Element.End()
}

func (*ArrayType) ExpressionNode() {}
