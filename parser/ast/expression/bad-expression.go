package expression

import "joern-go/parser/ast"

// A BadExpression node is a placeholder for an expression containing
// syntax errors for which a correct expression node cannot be
// created.
type BadExpression struct {
	ast.Node
	From int `json:"from"` // position range of bad expression
	To   int `json:"to"`
}

func (x BadExpression) Start() int {
	return x.From
}

func (x BadExpression) End() int {
	return x.To
}

func (x BadExpression) ExpressionNode() {}
