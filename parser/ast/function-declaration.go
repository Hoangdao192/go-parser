package ast

// A FunctionDeclaration node represents a function declaration.
type FunctionDeclaration struct {
	Node
	Doc      *CommentGroup   `json:"doc"`      // associated documentation; or nil
	Receiver *FieldList      `json:"receiver"` // receiver (methods); or nil (functions)
	Name     *Identifier     `json:"name"`     // function/method name
	Type     *FunctionType   `json:"type"`     // function signature: type and value parameters, results, and position of "func" keyword
	Body     *BlockStatement `json:"body"`     // function body; or nil for external (non-Go) function
}

func (d *FunctionDeclaration) Start() int {
	return d.Type.Start()
}

func (d *FunctionDeclaration) End() int {
	if d.Body.End() != 0 {
		return d.Body.End()
	}
	return d.Type.End()
}

func (*FunctionDeclaration) declarationNode() {}
