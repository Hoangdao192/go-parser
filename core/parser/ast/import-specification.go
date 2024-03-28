package ast

// An ImportSpecification node represents a single package import.
type ImportSpecification struct {
	Node
	Doc         *CommentGroup `json:"doc"`         // associated documentation; or nil
	Name        *Identifier   `json:"name"`        // local package name (including "."); or nil
	Path        *BasicLiteral `json:"path"`        // import path
	Comment     *CommentGroup `json:"comment"`     // line comments; or nil
	EndPosition int           `json:"endPosition"` // end of spec (overrides Path.Pos if nonzero)
}

func (s *ImportSpecification) Start() int {
	//if s.Name != nil {
	//	return s.Name.Start()
	//}
	if s.Name.Start() != 0 {
		return s.Name.Start()
	}
	return s.Path.Start()
}

func (s *ImportSpecification) End() int {
	if s.EndPosition != 0 {
		return s.EndPosition
	}
	return s.Path.End()
}

func (*ImportSpecification) specificationNode() {}
