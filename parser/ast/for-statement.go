package ast

// A ForStatement represents a for statement.
type ForStatement struct {
	Node
	For            int             `json:"for"`            // position of "for" keyword
	Initialization Statement       `json:"initialization"` // initialization statement; or nil
	Condition      Expression      `json:"condition"`      // condition; or nil
	Post           Statement       `json:"post"`           // post iteration statement; or nil
	Body           *BlockStatement `json:"body"`
}

func (s *ForStatement) Position() int {
	return s.For
}

func (s *ForStatement) End() int {
	return s.Body.End()
}

func (*ForStatement) StatementNode() {}
