package statement

import "joern-go/parser/ast/expression"

// A GoStatement node represents a go statement.
type GoStatement struct {
	Go   int                        `json:"go"` // position of "go" keyword
	Call *expression.CallExpression `json:"call"`
}

func (s *GoStatement) Start() int {
	return s.Go
}

func (s *GoStatement) End() int {
	return s.Call.End()
}

func (*GoStatement) StatementNode() {}
