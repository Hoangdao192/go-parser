package types

import "joern-go/parser/ast"

// An InterfaceType node represents an interface type.
type InterfaceType struct {
	Interface  int            `json:"interface"`  // position of "interface" keyword
	Methods    *ast.FieldList `json:"methods"`    // list of embedded interfaces, methods, or types
	Incomplete bool           `json:"incomplete"` // true if (source) methods or types are missing in the Methods list
}

func (x *InterfaceType) Position() int {
	return x.Interface
}

func (x *InterfaceType) End() int {
	return x.Methods.End()
}

func (*InterfaceType) ExpressionNode() {}
