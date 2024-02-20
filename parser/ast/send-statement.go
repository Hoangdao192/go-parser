package ast

// A SendStatement node represents a send statement.
type SendStatement struct {
	Node
	Chanel Expression `json:"Chan"`
	// position of "<-"
	Arrow int        `json:"arrow"`
	Value Expression `json:"Value"`
}

func (s *SendStatement) Start() int {
	return s.Chanel.Start()
}

func (s *SendStatement) End() int {
	return s.Value.End()
}

func (*SendStatement) StatementNode() {}
