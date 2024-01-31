package ast

// An Identifier node represents an identifier.
type Identifier struct {
	Node
	// identifier position
	NamePos int `json:"namePos"`
	// identifier name
	Name string `json:"name"`
	// denoted object; or nil
	Object *Object `json:"object"`
}

func (x Identifier) Start() int {
	return x.NamePos
}

func (x Identifier) End() int {
	return x.NamePos + len(x.Name)
}

func (Identifier) ExpressionNode() {}
