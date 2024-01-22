package expression

// A BinaryExpression node represents a binary expression.
type BinaryExpression struct {
	LeftExpression  Expression `json:"leftExpression"`  // left operand
	OpPos           int        `json:"opPos"`           // position of Op
	Op              int        `json:"op"`              // operator
	RightExpression Expression `json:"rightExpression"` // right operand
}

func (x *BinaryExpression) Position() int {
	return x.LeftExpression.Position()
}

func (x *BinaryExpression) End() int {
	return x.RightExpression.End()
}

func (*BinaryExpression) ExpressionNode() {}
