package declaration

import "joern-go/parser/ast"

type Declaration interface {
	ast.Node
	declarationNode()
}
