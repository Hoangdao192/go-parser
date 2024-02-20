package ast

// A ChanelType node represents a channel type.
type ChanelType struct {
	Node
	Begin     int        `json:"begin"`     // position of "chan" keyword or "<-" (whichever comes first)
	Arrow     int        `json:"arrow"`     // position of "<-" (token.NoPos if there is no "<-")
	Direction int        `json:"direction"` // channel direction
	Value     Expression `json:"value"`     // value type
}

func (x *ChanelType) Start() int {
	return x.Begin
}

func (x *ChanelType) End() int {
	return x.Value.End()
}

func (*ChanelType) ExpressionNode() {}
