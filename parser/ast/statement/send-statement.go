package statement

import (
	"joern-go/parser/ast/expression"
)

// A SendStatement node represents a send statement.
type SendStatement struct {
	Chanel expression.Expression `json:"Chan"`
	// position of "<-"
	Arrow int                   `json:"arrow"`
	Value expression.Expression `json:"Value"`
}

func (s *SendStatement) Position() int {
	return s.Chanel.Position()
}

func (s *SendStatement) End() int {
	return s.Value.End()
}

func (*SendStatement) StatementNode() {}
