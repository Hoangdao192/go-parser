package ast

type Declaration interface {
	INode
	declarationNode()
}
