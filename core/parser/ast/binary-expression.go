package ast

// A BinaryExpression node represents a binary expression.
type BinaryExpression struct {
	Node
	LeftExpression  Expression `json:"leftExpression"`  // left operand
	OpPos           int        `json:"opPos"`           // position of Op
	Op              int        `json:"op"`              // operator
	RightExpression Expression `json:"rightExpression"` // right operand
}

func (x *BinaryExpression) Start() int {
	return x.LeftExpression.Start()
}

func (x *BinaryExpression) End() int {
	return x.RightExpression.End()
}

func (*BinaryExpression) ExpressionNode() {}
