package ast

import (
	"joern-go/parser/ast/expression"
)

// A Field represents a Field declaration list in a struct type,
// a method list in an interface type, or a parameter/result declaration
// in a signature.
// Field.Names is nil for unnamed parameters (parameter lists which only contain types)
// and embedded struct fields. In the latter case, the field name is the type name.
type Field struct {
	Doc     *CommentGroup            `json:"doc"`     // associated documentation; or nil
	Names   []*expression.Identifier `json:"names"`   // field/method/(type) parameter names; or nil
	Type    Expression               `json:"type"`    // field/method/parameter type; or nil
	Tag     *expression.BasicLiteral `json:"tag"`     // field tag; or nil
	Comment *CommentGroup            `json:"comment"` // line comments; or nil
}

func (f *Field) Start() int {
	if len(f.Names) > 0 {
		return f.Names[0].Start()
	}
	if f.Type != nil {
		return f.Type.Position()
	}
	return 0
}

func (f *Field) End() int {
	if f.Tag != nil {
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

func (f *FieldList) Start() int {
	if f.Opening != 0 {
		return f.Opening
	}
	// the list should not be empty in this case;
	// be conservative and guard against bad ASTs
	if len(f.List) > 0 {
		return f.List[0].Start()
	}
	return 0
}

func (f *FieldList) End() int {
	if f.Closing != 0 {
		return f.Closing + 1
	}
	// the list should not be empty in this case;
	// be conservative and guard against bad ASTs
	if n := len(f.List); n > 0 {
		return f.List[n-1].End()
	}
	return 0
}

// NumFields returns the number of parameters or struct fields represented by a FieldList.
func (f *FieldList) NumFields() int {
	n := 0
	if f != nil {
		for _, g := range f.List {
			m := len(g.Names)
			if m == 0 {
				m = 1
			}
			n += m
		}
	}
	return n
}
