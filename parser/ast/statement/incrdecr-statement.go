package statement

import (
	"joern-go/parser/ast/expression"
)

// An IncrementDecrementStatement node represents an increment or decrement statement.
type IncrementDecrementStatement struct {
	Expression    expression.Expression `json:"Expression"`
	TokenPosition int                   `json:"TokenPosition"` // position of Token
	Token         int                   `json:"Token"`         // INC or DEC
}

func (s *IncrementDecrementStatement) Position() int {
	return s.Expression.Position()
}

func (s *IncrementDecrementStatement) End() int {
	return s.TokenPosition + 2 /* len("++") */
}

func (*IncrementDecrementStatement) StatementNode() {}
