package statement

import (
	"joern-go/parser/ast/expression"
)

// A LabeledStatement node represents a labeled statement.
type LabeledStatement struct {
	Statement Statement              `json:"statement"`
	Label     *expression.Identifier `json:"label"`
	// position of ":"
	Colon int `json:"colon"`
}

func (s *LabeledStatement) Position() int {
	return s.Label.Position()
}

func (s *LabeledStatement) End() int {
	return s.Statement.End()
}

func (*LabeledStatement) StatementNode() {}
