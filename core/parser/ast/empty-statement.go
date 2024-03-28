package ast

// An EmptyStatement node represents an empty statement.
// The "position" of the empty statement is the position
// of the immediately following (explicit or implicit) semicolon.
type EmptyStatement struct {
	Node
	// position of following ";"
	Semicolon int `json:"semicolon"`
	// if set, ";" was omitted in the source
	Implicit bool `json:"implicit"`
}

func (s *EmptyStatement) Start() int {
	return s.Semicolon
}

func (s *EmptyStatement) End() int {
	if s.Implicit {
		return s.Semicolon
	}
	return s.Semicolon + 1 /* len(";") */
}

func (*EmptyStatement) StatementNode() {}
