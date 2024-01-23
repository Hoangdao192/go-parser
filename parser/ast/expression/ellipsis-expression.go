package expression

import "joern-go/parser/ast"

// An Ellipsis node stands for the "..." type in a
// parameter list or the "..." length in an array type.
type Ellipsis struct {
	ast.Node
	Ellipsis int        `json:"ellipsis"` // position of "..."
	Element  Expression `json:"element"`  // ellipsis element type (parameter lists only); or nil
}

func (x Ellipsis) Start() int {
	return x.Ellipsis
}

func (x Ellipsis) End() int {
	if x.Element != nil {
		return x.Element.End()
	}
	return x.Ellipsis + 3 // len("...")
}

func (Ellipsis) ExpressionNode() {}
