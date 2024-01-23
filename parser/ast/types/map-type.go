package types

import "joern-go/parser/ast"

// A MapType node represents a map type.
type MapType struct {
	Map   int            `json:"map"` // position of "map" keyword
	Key   ast.Expression `json:"key"`
	Value ast.Expression `json:"value"`
}

func (x *MapType) Start() int {
	return x.Map
}

func (x *MapType) End() int {
	return x.Value.End()
}

func (*MapType) ExpressionNode() {}
