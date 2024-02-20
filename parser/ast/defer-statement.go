package ast

// A DeferStatement node represents a defer statement.
type DeferStatement struct {
	Node
	Defer int             `json:"defer"` // position of "defer" keyword
	Call  *CallExpression `json:"call"`
}

func (s *DeferStatement) Start() int {
	return s.Defer
}

func (s *DeferStatement) End() int {
	return s.Call.End()
}

func (*DeferStatement) StatementNode() {}
