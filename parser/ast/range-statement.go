package ast

// A RangeStatement represents a for statement with a range clause.
type RangeStatement struct {
	Node
	For        int            `json:"for"`        // position of "for" keyword
	Key        Expression     `json:"key"`        // Key may be nil
	Value      Expression     `json:"value"`      // Value may be nil
	TokenPos   int            `json:"tokenPos"`   // position of Tok; invalid if Key == nil
	Token      int            `json:"token"`      // ILLEGAL if Key == nil, ASSIGN, DEFINE
	Range      int            `json:"range"`      // position of "range" keyword
	Expression Expression     `json:"expression"` // value to range over
	Body       BlockStatement `json:"body"`
}

func (s RangeStatement) Start() int {
	return s.For
}

func (RangeStatement) StatementNode() {}
