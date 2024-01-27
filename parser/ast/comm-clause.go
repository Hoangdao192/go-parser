package ast

// A CommClause node represents a case of a select statement.
type CommClause struct {
	Node
	Case  int         `json:"case"`  // position of "case" or "default" keyword
	Comm  Statement   `json:"comm"`  // send or receive statement; nil means default case
	Colon int         `json:"colon"` // position of ":"
	Body  []Statement `json:"body"`  // statement list; or nil
}

func (s CommClause) Start() int {
	return s.Case
}

func (s CommClause) End() int {
	if n := len(s.Body); n > 0 {
		return s.Body[n-1].End()
	}
	return s.Colon + 1
}

func (CommClause) StatementNode() {}
