package statement

import "joern-go/parser/ast"

// A SelectStatement node represents a select statement.
type SelectStatement struct {
	ast.Node
	Select int            `json:"select"` // position of "select" keyword
	Body   BlockStatement `json:"body"`   // CommClauses only
}

func (s SelectStatement) Start() int {
	return s.Select
}

func (s SelectStatement) End() int {
	return s.Body.End()
}

func (SelectStatement) StatementNode() {}
