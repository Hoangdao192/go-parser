package statement

import "joern-go/parser/ast"

// A Statement is represented by a tree consisting of one
// or more of the following concrete statement nodes.
type Statement interface {
	ast.INode
	StatementNode()
}
