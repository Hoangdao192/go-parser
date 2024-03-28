package ast

// A MapType node represents a map type.
type MapType struct {
	Node
	Map   int        `json:"map"` // position of "map" keyword
	Key   Expression `json:"key"`
	Value Expression `json:"value"`
}

func (x *MapType) Start() int {
	return x.Map
}

func (x *MapType) End() int {
	return x.Value.End()
}

func (*MapType) ExpressionNode() {}
