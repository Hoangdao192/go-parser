package ast

// An Identifier node represents an identifier.
type Identifier struct {
	// identifier position
	NamePos int `json:"NamePos"`
	// identifier name
	Name string `json:"Name"`
	// denoted object; or nil
	Object *Object `json:"Object"`
}

func (x *Identifier) Position() int {
	return x.NamePos
}

func (x *Identifier) End() int {
	return x.NamePos + len(x.Name)
}
