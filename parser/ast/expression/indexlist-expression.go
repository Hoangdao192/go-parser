package expression

// An IndexListExpression node represents an expression followed by multiple
// indices.
type IndexListExpression struct {
	Expression   Expression   `json:"expression"`   // expression
	LeftBracket  int          `json:"leftBracket"`  // position of "["
	Indices      []Expression `json:"indices"`      // index expressions
	RightBracket int          `json:"rightBracket"` // position of "]"
}

func (x *IndexListExpression) Position() int {
	return x.Expression.Position()
}

func (x *IndexListExpression) End() int {
	return x.RightBracket + 1
}

func (*IndexListExpression) ExpressionNode() {}
