package declaration

import (
	"joern-go/parser/ast"
	"joern-go/parser/ast/expression"
	"joern-go/parser/ast/statement"
	"joern-go/parser/ast/types"
)

// A FunctionDeclaration node represents a function declaration.
type FunctionDeclaration struct {
	Doc      *ast.CommentGroup         `json:"doc"`      // associated documentation; or nil
	Receiver *ast.FieldList            `json:"receiver"` // receiver (methods); or nil (functions)
	Name     *expression.Identifier    `json:"name"`     // function/method name
	Type     *types.FunctionType       `json:"type"`     // function signature: type and value parameters, results, and position of "func" keyword
	Body     *statement.BlockStatement `json:"body"`     // function body; or nil for external (non-Go) function
}

func (d *FunctionDeclaration) Position() int {
	return d.Type.Position()
}

func (d *FunctionDeclaration) End() int {
	if d.Body != nil {
		return d.Body.End()
	}
	return d.Type.End()
}

func (*FunctionDeclaration) declarationNode() {}