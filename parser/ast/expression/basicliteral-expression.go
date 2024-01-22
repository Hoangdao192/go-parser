package expression

// A BasicLiteral node represents a literal of basic type.
type BasicLiteral struct {
	ValuePos int    `json:"valuePos"` // literal position
	Kind     int    `json:"kind"`     // token.INT, token.FLOAT, token.IMAG, token.CHAR, or token.STRING
	Value    string `json:"value"`    // literal string; e.g. 42, 0x7f, 3.14, 1e-9, 2.4i, 'a', '\x7f', "foo" or `\m\n\o`
}

func (x *BasicLiteral) Position() int {
	return x.ValuePos
}

func (x *BasicLiteral) End() int {
	return x.ValuePos + len(x.Value)
}

func (*BasicLiteral) ExpressionNode() {}
