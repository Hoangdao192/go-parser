package statement

import "joern-go/parser/ast/expression"

// A SwitchStatement node represents an expression switch statement.
type SwitchStatement struct {
	Switch         int                   `json:"switch"`         // position of "switch" keyword
	Initialization Statement             `json:"initialization"` // initialization statement; or nil
	Tag            expression.Expression `json:"tag"`            // tag expression; or nil
	Body           *BlockStatement       `json:"body"`           // CaseClauses only
}

func (s *SwitchStatement) Position() int {
	return s.Switch
}

func (s *SwitchStatement) End() int {
	return s.Body.End()
}

func (*SwitchStatement) StatementNode() {}
