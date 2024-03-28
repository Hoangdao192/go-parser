package ast

// A CaseClause represents a case of an expression or type switch statement.
type CaseClause struct {
	Node
	Case  int          `json:"case"`  // position of "case" or "default" keyword
	List  []Expression `json:"list"`  // list of expressions or types; nil means default case
	Colon int          `json:"colon"` // position of ":"
	Body  []Statement  `json:"body"`  // statement list; or nil
}

func (s *CaseClause) Start() int {
	return s.Case
}

func (s *CaseClause) End() int {
	if n := len(s.Body); n > 0 {
		return s.Body[n-1].End()
	}
	return s.Colon + 1
}

func (*CaseClause) StatementNode() {}