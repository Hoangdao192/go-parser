package ast

// A ParenthesizedExpression node represents a parenthesized expression.
type ParenthesizedExpression struct {
	Node
	Lparen     int        `json:"lparen"`     // position of "("
	Expression Expression `json:"expression"` // parenthesized expression
	Rparen     int        `json:"rparen"`     // position of ")"
}

func (x ParenthesizedExpression) Start() int {
	return x.Lparen
}

func (x ParenthesizedExpression) End() int {
	return x.Rparen + 1
}

func (ParenthesizedExpression) ExpressionNode() {}
