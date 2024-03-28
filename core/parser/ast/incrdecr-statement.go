package ast

// An IncrementDecrementStatement node represents an increment or decrement statement.
type IncrementDecrementStatement struct {
	Node
	Expression    Expression `json:"Expression"`
	TokenPosition int        `json:"TokenPosition"` // position of Token
	Token         int        `json:"Token"`         // INC or DEC
}

func (s *IncrementDecrementStatement) Start() int {
	return s.Expression.Start()
}

func (s *IncrementDecrementStatement) End() int {
	return s.TokenPosition + 2 /* len("++") */
}

func (*IncrementDecrementStatement) StatementNode() {}
