package declaration

import "joern-go/parser/ast"

type Declaration interface {
	ast.INode
	declarationNode()
}
