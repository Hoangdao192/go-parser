package ast

// A Field represents a Field declaration list in a struct type,
// a method list in an interface type, or a parameter/result declaration
// in a signature.
// Field.Names is nil for unnamed parameters (parameter lists which only contain types)
// and embedded struct fields. In the latter case, the field name is the type name.
type Field struct {
	Node
	Doc     CommentGroup `json:"doc"`     // associated documentation; or nil
	Names   []Identifier `json:"names"`   // field/method/(type) parameter names; or nil
	Type    Expression   `json:"type"`    // field/method/parameter type; or nil
	Tag     BasicLiteral `json:"tag"`     // field tag; or nil
	Comment CommentGroup `json:"comment"` // line comments; or nil
}

func (f Field) Start() int {
	if len(f.Names) > 0 {
		return f.Names[0].Start()
	}
	if f.Type != nil {
		return f.Type.Start()
	}
	return 0
}

func (f Field) End() int {
	if f.Tag.Start() != 0 {
		return f.Tag.End()
	}
	if f.Type != nil {
		return f.Type.End()
	}
	if len(f.Names) > 0 {
		return f.Names[len(f.Names)-1].End()
	}
	return 0
}
