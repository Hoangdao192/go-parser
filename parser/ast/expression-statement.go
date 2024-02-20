package ast

// An ExpressionStatement node represents a (stand-alone) expression
// in a statement list.
type ExpressionStatement struct {
	Node
	Expression Expression `json:"Expression"`
}

func (s *ExpressionStatement) Start() int {
	return s.Expression.Start()
}

func (s *ExpressionStatement) End() int {
	return s.Expression.End()
}

func (*ExpressionStatement) StatementNode() {}
