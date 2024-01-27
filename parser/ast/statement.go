package ast

// A Statement is represented by a tree consisting of one
// or more of the following concrete statement nodes.
type Statement interface {
	INode
	StatementNode()
}
