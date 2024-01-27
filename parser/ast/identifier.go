package ast

// An Identifier node represents an identifier.
type Identifier struct {
	Node
	// identifier position
	NamePos int `json:"NamePos"`
	// identifier name
	Name string `json:"Name"`
	// denoted object; or nil
	Object *Object `json:"Object"`
}

func (x Identifier) Start() int {
	return x.NamePos
}

func (x Identifier) End() int {
	return x.NamePos + len(x.Name)
}

func (Identifier) ExpressionNode() {}
