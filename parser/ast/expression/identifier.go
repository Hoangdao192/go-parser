package expression

import "joern-go/parser/ast"

// An Identifier node represents an identifier.
type Identifier struct {
	// identifier position
	NamePos int `json:"NamePos"`
	// identifier name
	Name string `json:"Name"`
	// denoted object; or nil
	Object *ast.Object `json:"Object"`
}

func (x *Identifier) Start() int {
	return x.NamePos
}

func (x *Identifier) End() int {
	return x.NamePos + len(x.Name)
}

func (*Identifier) ExpressionNode() {}
