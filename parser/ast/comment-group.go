package ast

// A CommentGroup represents a sequence of comments
// with no other tokens and no empty lines between.
type CommentGroup struct {
	Node
	Comments []*Comment // len(Comments) > 0
}

func (g CommentGroup) Start() int {
	return g.Comments[0].Start()
}
func (g CommentGroup) End() int {
	return g.Comments[len(g.Comments)-1].End()
}
