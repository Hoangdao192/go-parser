package statement

// A TypeSwitchStatement node represents a type switch statement.
type TypeSwitchStatement struct {
	Switch         int             `json:"switch"`         // position of "switch" keyword
	Initialization Statement       `json:"initialization"` // initialization statement; or nil
	Assign         Statement       `json:"assign"`         // x := y.(type) or y.(type)
	Body           *BlockStatement `json:"body"`           // CaseClauses only
}

func (s *TypeSwitchStatement) Position() int {
	return s.Switch
}

func (s *TypeSwitchStatement) End() int {
	return s.Body.End()
}

func (*TypeSwitchStatement) StatementNode() {}
