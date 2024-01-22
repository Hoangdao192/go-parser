package specification

import "joern-go/parser/ast"

// The Specification type stands for any of *ImportSpecification, *ValueSpecification, and *TypeSpecification.
type Specification interface {
	ast.Node
	specificationNode()
}
