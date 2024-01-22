package statement

import (
	"joern-go/parser/ast/expression"
)

// An ExpressionStatement node represents a (stand-alone) expression
// in a statement list.
type ExpressionStatement struct {
	Expression expression.Expression `json:"Expression"`
}

func (s *ExpressionStatement) Position() int {
	return s.Expression.Position()
}

func (s *ExpressionStatement) End() int {
	return s.Expression.End()
}

func (*ExpressionStatement) StatementNode() {}
