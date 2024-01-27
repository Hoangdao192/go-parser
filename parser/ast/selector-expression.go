package ast

// A SelectorExpression node represents an expression followed by a selector.
type SelectorExpression struct {
	Node
	Expression Expression `json:"expression"` // expression
	Sel        Identifier `json:"sel"`        // field selector
}

func (x SelectorExpression) Start() int {
	return x.Expression.Start()
}

func (x SelectorExpression) End() int {
	return x.Sel.End()
}

func (SelectorExpression) ExpressionNode() {}
