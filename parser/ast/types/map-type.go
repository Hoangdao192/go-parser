package types

import (
	"joern-go/parser/ast"
	"joern-go/parser/ast/expression"
)

// A MapType node represents a map type.
type MapType struct {
	ast.Node
	Map   int                   `json:"map"` // position of "map" keyword
	Key   expression.Expression `json:"key"`
	Value expression.Expression `json:"value"`
}

func (x MapType) Start() int {
	return x.Map
}

func (x MapType) End() int {
	return x.Value.End()
}

func (MapType) ExpressionNode() {}
