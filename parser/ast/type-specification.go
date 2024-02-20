package ast

// A TypeSpecification node represents a type declaration (TypeSpecification production).
type TypeSpecification struct {
	Node
	Doc        *CommentGroup `json:"doc"`        // associated documentation; or nil
	Name       *Identifier   `json:"name"`       // type name
	TypeParams *FieldList    `json:"typeParams"` // type parameters; or nil
	Assign     int           `json:"assign"`     // position of '=', if any
	Type       Expression    `json:"type"`       // *Ident, *ParenExpression, *SelectorExpression, *StarExpression, or any of the *XxxTypes
	Comment    *CommentGroup `json:"comment"`    // line comments; or nil
}

func (s *TypeSpecification) Start() int {
	return s.Name.Start()
}

func (s *TypeSpecification) End() int {
	return s.Type.End()
}

func (*TypeSpecification) specificationNode() {}
