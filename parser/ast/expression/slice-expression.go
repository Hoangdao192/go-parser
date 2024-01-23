package expression

// A SliceExpression node represents an expression followed by slice indices.
type SliceExpression struct {
	Expression   Expression `json:"expression"`   // expression
	LeftBracket  int        `json:"leftBracket"`  // position of "["
	Low          Expression `json:"low"`          // begin of slice range; or nil
	High         Expression `json:"high"`         // end of slice range; or nil
	Max          Expression `json:"max"`          // maximum capacity of slice; or nil
	Slice3       bool       `json:"slice3"`       // true if 3-index slice (2 colons present)
	RightBracket int        `json:"rightBracket"` // position of "]"
}

func (x *SliceExpression) Start() int {
	return x.Expression.Start()
}

func (x *SliceExpression) End() int {
	return x.RightBracket + 1
}

func (*SliceExpression) ExpressionNode() {}
