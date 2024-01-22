package statement

// A BadStatement node is a placeholder for statements containing
// syntax errors for which no correct statement nodes can be
// created.
type BadStatement struct {
	From int `json:"from"`
	To   int `json:"to"` // position range of bad statement
}

func (s *BadStatement) Position() int {
	return s.From
}

func (s *BadStatement) End() int {
	return s.To
}

func (*BadStatement) StatementNode() {}
