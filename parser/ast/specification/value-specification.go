package specification

import (
	"joern-go/parser/ast"
	"joern-go/parser/ast/expression"
)

// A ValueSpecification node represents a constant or variable declaration
// (ConstSpecification or VarSpecification production).
type ValueSpecification struct {
	ast.Node
	Doc     ast.CommentGroup        `json:"doc"`     // associated documentation; or nil
	Names   []expression.Identifier `json:"names"`   // value names (len(Names) > 0)
	Type    expression.Expression   `json:"type"`    // value type; or nil
	Values  []expression.Expression `json:"values"`  // initial values; or nil
	Comment ast.CommentGroup        `json:"comment"` // line comments; or nil
}

func (s ValueSpecification) Start() int {
	return s.Names[0].Start()
}

func (s ValueSpecification) End() int {
	if n := len(s.Values); n > 0 {
		return s.Values[n-1].End()
	}
	if s.Type != nil {
		return s.Type.End()
	}
	return s.Names[len(s.Names)-1].End()
}

func (ValueSpecification) specificationNode() {}
