package expression

// A TypeAssertExpression node represents an expression followed by a
// type assertion.
type TypeAssertExpression struct {
	Expression Expression `json:"expression"` // expression
	Lparen     int        `json:"lparen"`     // position of "("
	Type       Expression `json:"type"`       // asserted type; nil means type switch X.(type)
	Rparen     int        `json:"rparen"`     // position of ")"
}

func (x *TypeAssertExpression) Start() int {
	return x.Expression.Start()
}

func (x *TypeAssertExpression) End() int {
	return x.Rparen + 1
}

func (*TypeAssertExpression) ExpressionNode() {}
