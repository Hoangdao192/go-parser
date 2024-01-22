package expression

// A StarExpression node represents an expression of the form "*" Expression.
// Semantically it could be a unary "*" expression, or a pointer type.
type StarExpression struct {
	Star       int        `json:"star"`       // position of "*"
	Expression Expression `json:"expression"` // operand
}

func (x *StarExpression) Position() int {
	return x.Star
}

func (x *StarExpression) End() int {
	return x.Expression.End()
}

func (*StarExpression) ExpressionNode() {}