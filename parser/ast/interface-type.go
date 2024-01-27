package ast

// An InterfaceType node represents an interface type.
type InterfaceType struct {
	Node
	Interface  int       `json:"interface"`  // position of "interface" keyword
	Methods    FieldList `json:"methods"`    // list of embedded interfaces, methods, or types
	Incomplete bool      `json:"incomplete"` // true if (source) methods or types are missing in the Methods list
}

func (x InterfaceType) Start() int {
	return x.Interface
}

func (x InterfaceType) End() int {
	return x.Methods.End()
}

func (InterfaceType) ExpressionNode() {}
