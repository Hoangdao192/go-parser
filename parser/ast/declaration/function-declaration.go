package declaration

import (
	"joern-go/parser/ast"
	"joern-go/parser/ast/expression"
	"joern-go/parser/ast/statement"
	"joern-go/parser/ast/types"
)

// A FunctionDeclaration node represents a function declaration.
type FunctionDeclaration struct {
	ast.Node
	Doc      ast.CommentGroup         `json:"doc"`      // associated documentation; or nil
	Receiver ast.FieldList            `json:"receiver"` // receiver (methods); or nil (functions)
	Name     expression.Identifier    `json:"name"`     // function/method name
	Type     types.FunctionType       `json:"type"`     // function signature: type and value parameters, results, and position of "func" keyword
	Body     statement.BlockStatement `json:"body"`     // function body; or nil for external (non-Go) function
}

func (d FunctionDeclaration) Start() int {
	return d.Type.Start()
}

func (d FunctionDeclaration) End() int {
	if d.Body.End() != 0 {
		return d.Body.End()
	}
	return d.Type.End()
}

func (FunctionDeclaration) declarationNode() {}
