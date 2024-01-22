package expression

// A CompositeLiteral node represents a composite literal.
type CompositeLiteral struct {
	Type       Expression   `json:"type"`       // literal type; or nil
	Lbrace     int          `json:"lbrace"`     // position of "{"
	Elements   []Expression `json:"elements"`   // list of composite elements; or nil
	Rbrace     int          `json:"rbrace"`     // position of "}"
	Incomplete bool         `json:"incomplete"` // true if (source) expressions are missing in the Elements list
}

func (x *CompositeLiteral) Position() int {
	if x.Type != nil {
		return x.Type.Position()
	}
	return x.Lbrace
}

func (x *CompositeLiteral) End() int {
	return x.Rbrace + 1
}

func (*CompositeLiteral) ExpressionNode() {}