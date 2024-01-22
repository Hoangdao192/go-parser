package ast

// A CommentGroup represents a sequence of comments
// with no other tokens and no empty lines between.
type CommentGroup struct {
	List []*Comment // len(List) > 0
}

func (g *CommentGroup) Position() int {
	return g.List[0].Position()
}
func (g *CommentGroup) End() int {
	return g.List[len(g.List)-1].End()
}
