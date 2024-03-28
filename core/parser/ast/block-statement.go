package ast

// A BlockStatement node represents a braced statement list.
type BlockStatement struct {
	Node
	Lbrace int         `json:"lbrace"` // position of "{"
	List   []Statement `json:"list"`
	Rbrace int         `json:"rbrace"` // position of "}", if any (may be absent due to syntax error)
}

func (s *BlockStatement) Start() int {
	return s.Lbrace
}

func (s *BlockStatement) End() int {
	if s.Rbrace != 0 {
		return s.Rbrace + 1
	}
	if n := len(s.List); n > 0 {
		return s.List[n-1].End()
	}
	return s.Lbrace + 1
}

func (*BlockStatement) StatementNode() {}
