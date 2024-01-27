package ast

// A ReturnStatement node represents a return statement.
type ReturnStatement struct {
	Node
	Return  int          `json:"return"`  // position of "return" keyword
	Results []Expression `json:"results"` // result expressions; or nil
}

func (s ReturnStatement) Start() int {
	return s.Return
}

func (s ReturnStatement) End() int {
	if n := len(s.Results); n > 0 {
		return s.Results[n-1].End()
	}
	return s.Return + 6 // len("return")
}

func (ReturnStatement) StatementNode() {}
