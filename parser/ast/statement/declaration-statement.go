package statement

import (
	"joern-go/parser/ast/declaration"
)

// A DeclarationStatement node represents a declaration in a statement list.
type DeclarationStatement struct {
	// *GenDecl with CONST, TYPE, or VAR token
	Declaration declaration.Declaration `json:"declaration"`
}

func (s *DeclarationStatement) Position() int {
	return s.Declaration.Position()
}

func (s *DeclarationStatement) End() int {
	return s.Declaration.End()
}

func (*DeclarationStatement) StatementNode() {}
