package ast

// A LabeledStatement node represents a labeled statement.
type LabeledStatement struct {
	Node
	Statement Statement   `json:"statement"`
	Label     *Identifier `json:"label"`
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
