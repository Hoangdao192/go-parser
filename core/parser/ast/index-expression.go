package ast

// An IndexExpression node represents an expression followed by an index.
type IndexExpression struct {
	Node
	Expression   Expression `json:"expression"`   // expression
	LeftBracket  int        `json:"leftBracket"`  // position of "["
	Index        Expression `json:"index"`        // index expression
	RightBracket int        `json:"rightBracket"` // position of "]"
}

func (x *IndexExpression) Start() int {
	return x.Expression.Start()
}

func (x *IndexExpression) End() int {
	return x.RightBracket + 1
}

func (*IndexExpression) ExpressionNode() {}
