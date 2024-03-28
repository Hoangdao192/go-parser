package ast

// An ArrayType node represents an array or slice type.
type ArrayType struct {
	Node
	LeftBracket int        `json:"leftBracket"` // position of "["
	Length      Expression `json:"length"`      // Ellipsis node for [...]T array types, nil for slice types
	Element     Expression `json:"element"`     // element type
}

func (x *ArrayType) Start() int {
	return x.LeftBracket
}

func (x *ArrayType) End() int {
	return x.Element.End()
}

func (*ArrayType) ExpressionNode() {}
