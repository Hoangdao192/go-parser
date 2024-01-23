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

func (s *LabeledStatement) Start() int {
	return s.Label.Start()
}

func (s *LabeledStatement) End() int {
	return s.Statement.End()
}

func (*LabeledStatement) StatementNode() {}
