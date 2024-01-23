package ast

// A FieldList represents a list of Fields, enclosed by parentheses,
// curly braces, or square brackets.
type FieldList struct {
	Node
	Opening int     `json:"opening"` // position of opening parenthesis/brace/bracket, if any
	List    []Field `json:"list"`    // field list; or nil
	Closing int     `json:"closing"` // position of closing parenthesis/brace/bracket, if any
}
