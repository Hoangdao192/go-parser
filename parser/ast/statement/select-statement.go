package statement

// A SelectStatement node represents a select statement.
type SelectStatement struct {
	Select int             `json:"select"` // position of "select" keyword
	Body   *BlockStatement `json:"body"`   // CommClauses only
}

func (s *SelectStatement) Position() int {
	return s.Select
}

func (s *SelectStatement) End() int {
	return s.Body.End()
}

func (*SelectStatement) StatementNode() {}
