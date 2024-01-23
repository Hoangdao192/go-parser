package ast

// A Comment node represents a single //-style or /*-style comment.
//
// The Text field contains the comment text without carriage returns (\r) that
// may have been present in the source. Because a comment's end position is
// computed using len(Text), the position reported by End() does not match the
// true source end position for comments containing carriage returns.
type Comment struct {
	Node
	Slash int    // position of "/" starting the comment
	Text  string // comment text (excluding '\n' for //-style comments)
}

func (c Comment) Start() int {
	return c.Slash
}
func (c Comment) End() int {
	return c.Slash + len(c.Text)
}
