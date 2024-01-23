package parser

import (
	"go/ast"
	data "joern-go/parser/ast"
	"joern-go/parser/ast/expression"
)

func BuildComment(comment *ast.Comment) data.Comment {
	parsed := data.Comment{
		Node: data.Node{
			Children:      []data.INode{},
			StartPosition: 0,
			EndPosition:   0,
		},
		Slash: int(comment.Slash),
		Text:  comment.Text,
	}
	parsed.Node.StartPosition = parsed.Start()
	parsed.Node.EndPosition = parsed.End()
	return parsed
}

func BuildCommentGroup(commentGroup *ast.CommentGroup) data.CommentGroup {
	parsed := data.CommentGroup{
		Node: data.Node{
			Children:      []data.INode{},
			StartPosition: int(commentGroup.Pos()),
			EndPosition:   int(commentGroup.End()),
		},
		Comments: []*data.Comment{},
	}
	return parsed
}

func BuildField(field *ast.Field) data.Field {
	return data.Field{
		Node: data.Node{
			Children:      []data.INode{},
			StartPosition: int(field.Pos()),
			EndPosition:   int(field.End()),
		},
	}
}

func BuildFieldList(fieldList *ast.FieldList) data.FieldList {
	return data.FieldList{
		Node: data.Node{
			Children:      []data.INode{},
			StartPosition: int(fieldList.Pos()),
			EndPosition:   int(fieldList.End()),
		},
		Opening: int(fieldList.Opening),
		List:    []data.Field{},
		Closing: int(fieldList.Closing),
	}
}

func BuildBadExpression(badExpression *ast.BadExpr) expression.BadExpression {
	return expression.BadExpression{
		Node: data.Node{
			Children:      []data.INode{},
			StartPosition: int(badExpression.Pos()),
			EndPosition:   int(badExpression.End()),
		},
		From: int(badExpression.From),
		To:   int(badExpression.To),
	}
}

func BuildIdentifier(identifier *ast.Ident) expression.Identifier {
	return expression.Identifier{
		Node:    BuildNode(identifier),
		NamePos: int(identifier.NamePos),
		Name:    identifier.Name,
		//Object: identifier.
	}
}

func BuildBasicLiteral(basicLiteral *ast.BasicLit) expression.BasicLiteral {
	return expression.BasicLiteral{
		Node:     BuildNode(basicLiteral),
		ValuePos: int(basicLiteral.ValuePos),
		Kind:     int(basicLiteral.Kind),
		Value:    basicLiteral.Value,
	}
}

func BuildEllipsis(ellipsis *ast.Ellipsis) expression.Ellipsis {
	return expression.Ellipsis{
		Node:     BuildNode(ellipsis),
		Ellipsis: int(ellipsis.Ellipsis),
	}
}

func BuildFunctionLiteral(functionLiteral *ast.FuncLit) expression.FunctionLiteral {
	return expression.FunctionLiteral{
		Node: BuildNode(functionLiteral),
	}
}

func BuildCompositeLiteral(compositeLiteral *ast.CompositeLit) expression.CompositeLiteral {
	return expression.CompositeLiteral{
		Node:       BuildNode(compositeLiteral),
		Lbrace:     int(compositeLiteral.Lbrace),
		Rbrace:     int(compositeLiteral.Rbrace),
		Elements:   []expression.Expression{},
		Incomplete: compositeLiteral.Incomplete,
	}
}

func BuildParenthesizedExpression(expr *ast.ParenExpr) expression.ParenthesizedExpression {
	return expression.ParenthesizedExpression{
		Node:   BuildNode(expr),
		Lparen: int(expr.Lparen),
		Rparen: int(expr.Rparen),
	}
}

func BuildNode(n ast.Node) data.Node {
	return data.Node{
		Children:      []data.INode{},
		StartPosition: int(n.Pos()),
		EndPosition:   int(n.End()),
	}
}
