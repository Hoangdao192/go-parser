package ast

// A ValueSpecification node represents a constant or variable declaration
// (ConstSpecification or VarSpecification production).
type ValueSpecification struct {
	Node
	Doc     *CommentGroup `json:"doc"`     // associated documentation; or nil
	Names   []*Identifier `json:"names"`   // value names (len(Names) > 0)
	Type    Expression    `json:"type"`    // value type; or nil
	Values  []Expression  `json:"values"`  // initial values; or nil
	Comment *CommentGroup `json:"comment"` // line comments; or nil
}

func (s *ValueSpecification) Start() int {
	return s.Names[0].Start()
}

func (s *ValueSpecification) End() int {
	if n := len(s.Values); n > 0 {
		return s.Values[n-1].End()
	}
	if s.Type != nil {
		return s.Type.End()
	}
	return s.Names[len(s.Names)-1].End()
}

func (*ValueSpecification) specificationNode() {}
