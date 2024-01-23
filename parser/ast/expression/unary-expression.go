package expression

import "joern-go/parser/ast"

// A UnaryExpression node represents a unary expression.
// Unary "*" expressions are represented via StarExpression nodes.
type UnaryExpression struct {
	ast.Node
	OpPos      int        `json:"opPos"`      // position of Op
	Op         int        `json:"op"`         // operator
	Expression Expression `json:"expression"` // operand
}

func (x UnaryExpression) Start() int {
	return x.OpPos
}

func (x UnaryExpression) End() int {
	return x.Expression.End()
}

func (UnaryExpression) ExpressionNode() {}
