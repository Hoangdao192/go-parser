package specification

import (
	"joern-go/parser/ast"
	"joern-go/parser/ast/expression"
)

// A TypeSpecification node represents a type declaration (TypeSpecification production).
type TypeSpecification struct {
	ast.Node
	Doc        ast.CommentGroup      `json:"doc"`        // associated documentation; or nil
	Name       expression.Identifier `json:"name"`       // type name
	TypeParams ast.FieldList         `json:"typeParams"` // type parameters; or nil
	Assign     int                   `json:"assign"`     // position of '=', if any
	Type       expression.Expression `json:"type"`       // *Ident, *ParenExpression, *SelectorExpression, *StarExpression, or any of the *XxxTypes
	Comment    ast.CommentGroup      `json:"comment"`    // line comments; or nil
}

func (s TypeSpecification) Start() int {
	return s.Name.Start()
}

func (s TypeSpecification) End() int {
	return s.Type.End()
}

func (TypeSpecification) specificationNode() {}
