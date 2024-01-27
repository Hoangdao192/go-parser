package ast

// A FunctionLiteral node represents a function literal.
type FunctionLiteral struct {
	Node
	Type FunctionType   `json:"type"` // function type
	Body BlockStatement `json:"body"` // function body
}

func (x FunctionLiteral) Start() int {
	return x.Type.Start()
}

func (x FunctionLiteral) End() int {
	return x.Body.End()
}

func (FunctionLiteral) ExpressionNode() {}
