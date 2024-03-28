package ast

// A BranchStatement node represents a break, continue, goto,
// or fallthrough statement.
type BranchStatement struct {
	Node
	TokenPosition int         `json:"tokenPosition"` // position of Token
	Token         int         `json:"token"`         // keyword token (BREAK, CONTINUE, GOTO, FALLTHROUGH)
	Label         *Identifier `json:"label"`         // label name; or nil
}

func (s *BranchStatement) Start() int {
	return s.TokenPosition
}

func (s *BranchStatement) End() int {
	if s.Label.End() != 0 {
		return s.Label.End()
	}
	return int(s.TokenPosition + len(string(rune(s.Token))))
}

func (*BranchStatement) StatementNode() {}
