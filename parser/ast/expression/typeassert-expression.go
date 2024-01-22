package expression

// A TypeAssertExpression node represents an expression followed by a
// type assertion.
type TypeAssertExpression struct {
	Expression Expression `json:"expression"` // expression
	Lparen     int        `json:"lparen"`     // position of "("
	Type       Expression `json:"type"`       // asserted type; nil means type switch X.(type)
	Rparen     int        `json:"rparen"`     // position of ")"
}

func (x *TypeAssertExpression) Position() int {
	return x.Expression.Position()
}

func (x *TypeAssertExpression) End() int {
	return x.Rparen + 1
}

func (*TypeAssertExpression) ExpressionNode() {}
