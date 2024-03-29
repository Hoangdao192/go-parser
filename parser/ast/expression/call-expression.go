package expression

// A CallExpression node represents an expression followed by an argument list.
type CallExpression struct {
	Function Expression   `json:"fun"`      // function expression
	Lparen   int          `json:"lparen"`   // position of "("
	Args     []Expression `json:"args"`     // function arguments; or nil
	Ellipsis int          `json:"ellipsis"` // position of "..." (token.NoPos if there is no "...")
	Rparen   int          `json:"rparen"`   // position of ")"
}

func (x *CallExpression) Pos() int {
	return x.Function.Position()
}

func (x *CallExpression) End() int {
	return x.Rparen + 1
}
