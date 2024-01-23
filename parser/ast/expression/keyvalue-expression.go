package expression

import "joern-go/parser/ast"

// A KeyValueExpression node represents (key : value) pairs
// in composite literals.
type KeyValueExpression struct {
	ast.Node
	Key   Expression `json:"key"`
	Colon int        `json:"colon"` // position of ":"
	Value Expression `json:"value"`
}

func (x KeyValueExpression) Start() int {
	return x.Key.Start()
}

func (x KeyValueExpression) End() int {
	return x.Value.End()
}

func (KeyValueExpression) ExpressionNode() {}
