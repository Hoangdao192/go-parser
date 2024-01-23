package types

import "joern-go/parser/ast"

// A StructType node represents a struct type.
type StructType struct {
	Struct     int            `json:"struct"`     // position of "struct" keyword
	Fields     *ast.FieldList `json:"fields"`     // list of field declarations
	Incomplete bool           `json:"incomplete"` // true if (source) fields are missing in the Fields list
}

func (x *StructType) Start() int {
	return x.Struct
}

func (x *StructType) End() int {
	return x.Fields.End()
}

func (*StructType) ExpressionNode() {}
