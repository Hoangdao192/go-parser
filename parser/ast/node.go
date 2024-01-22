package ast

type Node interface {
	Position() int // position of first character belonging to the node
	End() int      // position of first character immediately after the node
}
