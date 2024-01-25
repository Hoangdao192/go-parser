package statement

import (
	"joern-go/parser/ast"
	"joern-go/parser/ast/expression"
)

// An IfStatement node represents an if statement.
type IfStatement struct {
	ast.Node
	If             int                   `json:"if"`             // position of "if" keyword
	Initialization Statement             `json:"initialization"` // initialization statement; or nil
	Condition      expression.Expression `json:"condition"`      // condition
	Body           BlockStatement        `json:"body"`
	Else           Statement             `json:"else"` // else branch; or nil
}

func (s IfStatement) Start() int {
	return s.If
}

func (s IfStatement) End() int {
	if s.Else != nil {
		return s.Else.End()
	}
	return s.Body.End()
}

func (IfStatement) StatementNode() {}
