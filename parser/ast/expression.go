package ast

// Expression All expression nodes implement the Expression interface.
type Expression interface {
	INode
	ExpressionNode()
	//Start() int
	//End() int
}
