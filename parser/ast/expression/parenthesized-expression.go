package expression

// A ParenthesizedExpression node represents a parenthesized expression.
type ParenthesizedExpression struct {
	Lparen     int        `json:"lparen"`     // position of "("
	Expression Expression `json:"expression"` // parenthesized expression
	Rparen     int        `json:"rparen"`     // position of ")"
}

func (x *ParenthesizedExpression) Position() int {
	return x.Lparen
}

func (x *ParenthesizedExpression) End() int {
	return x.Rparen + 1
}

func (*ParenthesizedExpression) ExpressionNode() {}
