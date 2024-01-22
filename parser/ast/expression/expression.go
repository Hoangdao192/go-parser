package expression

import "joern-go/parser/ast"

// Expression All expression nodes implement the Expression interface.
type Expression interface {
	ast.Node
	expressionNode()
}
