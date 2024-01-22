package statement

import (
	"joern-go/parser/ast/expression"
)

// An AssignStatement node represents an assignment or
// a short variable declaration.
type AssignStatement struct {
	Lhs      []expression.Expression `json:"lhs"`
	TokenPos int                     `json:"tokenPos"` // position of Tok
	Token    int                     `json:"token"`    // assignment token, DEFINE
	Rhs      []expression.Expression `json:"rhs"`
}

func (s *AssignStatement) Position() int {
	return s.Lhs[0].Position()
}

func (s *AssignStatement) End() int {
	return s.Rhs[len(s.Rhs)-1].End()
}

func (*AssignStatement) StatementNode() {}
