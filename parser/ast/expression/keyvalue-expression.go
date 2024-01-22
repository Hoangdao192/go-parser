package expression

// A KeyValueExpression node represents (key : value) pairs
// in composite literals.
type KeyValueExpression struct {
	Key   Expression `json:"key"`
	Colon int        `json:"colon"` // position of ":"
	Value Expression `json:"value"`
}

func (x *KeyValueExpression) Position() int {
	return x.Key.Position()
}

func (x *KeyValueExpression) End() int {
	return x.Value.End()
}

func (*KeyValueExpression) ExpressionNode() {}
