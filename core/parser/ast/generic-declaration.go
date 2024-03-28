package ast

// A GenericDeclaration node (generic declaration node) represents an import,
// constant, type or variable declaration. A valid Lparen position
// (Lparen.IsValid()) indicates a parenthesized declaration.
//
// Relationship between Tok value and Specs element type:
//
//	token.IMPORT  *ImportSpec
//	token.CONST   *ValueSpec
//	token.TYPE    *TypeSpec
//	token.VAR     *ValueSpec
type GenericDeclaration struct {
	Node
	Doc            *CommentGroup   `json:"doc"`           // associated documentation; or nil
	TokenPosition  int             `json:"tokenPosition"` // position of Token
	Token          int             `json:"token"`         // IMPORT, CONST, TYPE, or VAR
	Lparen         int             `json:"lparen"`        // position of '(', if any
	Specifications []Specification `json:"specifications"`
	Rparen         int             `json:"rparen"` // position of ')', if any
}

func (d *GenericDeclaration) Start() int {
	return d.TokenPosition
}

func (d *GenericDeclaration) End() int {
	if d.Rparen != 0 {
		return d.Rparen + 1
	}
	return d.Specifications[0].End()
}

func (*GenericDeclaration) declarationNode() {}
