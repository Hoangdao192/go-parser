package expression

// A SelectorExpression node represents an expression followed by a selector.
type SelectorExpression struct {
	Expression Expression  `json:"expression"` // expression
	Sel        *Identifier `json:"sel"`        // field selector
}

func (x *SelectorExpression) Position() int {
	return x.Expression.Position()
}

func (x *SelectorExpression) End() int {
	return x.Sel.End()
}

func (*SelectorExpression) ExpressionNode() {}
