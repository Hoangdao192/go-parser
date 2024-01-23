package parser

import (
	"go/ast"
	data "joern-go/parser/ast"
)

func Build(comment *ast.Comment) data.Comment {
	parsed := data.Comment{
		Node: data.Node{
			Children: []*data.INode{},
			Start:    0,
			End:      0,
		},
		Slash: int(comment.Slash),
		Text:  comment.Text,
	}
	parsed.Node.Start = parsed.Start()
	parsed.Node.End = parsed.End()
	return parsed
}

func BuildCommentGroup(commentGroup *ast.CommentGroup) data.CommentGroup {
	parsed := data.CommentGroup{
		Node: data.Node{
			Children: []*data.INode{},
			Start:    0,
			End:      0,
		},
		Comments: []*data.Comment{},
	}
	return parsed
}
