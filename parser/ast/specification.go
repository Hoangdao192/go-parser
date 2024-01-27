package ast

// The Specification type stands for any of *ImportSpecification, *ValueSpecification, and *TypeSpecification.
type Specification interface {
	INode
	specificationNode()
}
