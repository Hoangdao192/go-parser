package specification

import (
	"joern-go/parser/ast"
	"joern-go/parser/ast/expression"
)

// An ImportSpecification node represents a single package import.
type ImportSpecification struct {
	Doc         *ast.CommentGroup        `json:"doc"`         // associated documentation; or nil
	Name        *expression.Identifier   `json:"name"`        // local package name (including "."); or nil
	Path        *expression.BasicLiteral `json:"path"`        // import path
	Comment     *ast.CommentGroup        `json:"comment"`     // line comments; or nil
	EndPosition int                      `json:"endPosition"` // end of spec (overrides Path.Pos if nonzero)
}

func (s *ImportSpecification) Start() int {
	if s.Name != nil {
		return s.Name.Start()
	}
	return s.Path.Start()
}

func (s *ImportSpecification) End() int {
	if s.EndPosition != 0 {
		return s.EndPosition
	}
	return s.Path.End()
}

func (*ImportSpecification) specificationNode() {}
