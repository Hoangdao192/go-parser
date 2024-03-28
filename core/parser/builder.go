package parser

import (
	"bytes"
	"fmt"
	"github.com/google/uuid"
	"go/ast"
	"go/printer"
	"go/token"
	data "joern-go/core/parser/ast"
	"time"
)

func BuildComment(comment *ast.Comment) data.Comment {
	parsed := data.Comment{
		Node:  BuildNode(comment),
		Slash: int(comment.Slash),
		Text:  comment.Text,
	}
	parsed.Node.StartPosition = parsed.Start()
	parsed.Node.EndPosition = parsed.End()
	return parsed
}

func BuildCommentGroup(commentGroup *ast.CommentGroup) data.CommentGroup {
	parsed := data.CommentGroup{
		Node:     BuildNode(commentGroup),
		Comments: []*data.Comment{},
	}
	return parsed
}

func BuildField(field *ast.Field) data.Field {
	return data.Field{
		Node: BuildNode(field),
	}
}

func BuildFieldList(fieldList *ast.FieldList) data.FieldList {
	return data.FieldList{
		Node:    BuildNode(fieldList),
		Opening: int(fieldList.Opening),
		List:    []*data.Field{},
		Closing: int(fieldList.Closing),
	}
}

func BuildBadExpression(badExpression *ast.BadExpr) data.BadExpression {
	return data.BadExpression{
		Node: BuildNode(badExpression),
		From: int(badExpression.From),
		To:   int(badExpression.To),
	}
}

func BuildIdentifier(identifier *ast.Ident) data.Identifier {
	return data.Identifier{
		Node:    BuildNode(identifier),
		NamePos: int(identifier.NamePos),
		Name:    identifier.Name,
		//Object: identifier.
	}
}

func BuildBasicLiteral(basicLiteral *ast.BasicLit) data.BasicLiteral {
	return data.BasicLiteral{
		Node:     BuildNode(basicLiteral),
		ValuePos: int(basicLiteral.ValuePos),
		Kind:     int(basicLiteral.Kind),
		Value:    basicLiteral.Value,
	}
}

func BuildEllipsis(ellipsis *ast.Ellipsis) data.Ellipsis {
	return data.Ellipsis{
		Node:     BuildNode(ellipsis),
		Ellipsis: int(ellipsis.Ellipsis),
	}
}

func BuildFunctionLiteral(functionLiteral *ast.FuncLit) data.FunctionLiteral {
	return data.FunctionLiteral{
		Node: BuildNode(functionLiteral),
	}
}

func BuildCompositeLiteral(compositeLiteral *ast.CompositeLit) data.CompositeLiteral {
	return data.CompositeLiteral{
		Node:       BuildNode(compositeLiteral),
		Lbrace:     int(compositeLiteral.Lbrace),
		Rbrace:     int(compositeLiteral.Rbrace),
		Elements:   []data.Expression{},
		Incomplete: compositeLiteral.Incomplete,
	}
}

func BuildParenthesizedExpression(expr *ast.ParenExpr) data.ParenthesizedExpression {
	return data.ParenthesizedExpression{
		Node:   BuildNode(expr),
		Lparen: int(expr.Lparen),
		Rparen: int(expr.Rparen),
	}
}

func BuildNode(n ast.Node) data.Node {

	var nodeType string
	switch n.(type) {
	// Comments and fields
	case *ast.Comment:
		nodeType = "Comment"
	case *ast.CommentGroup:
		nodeType = "CommentGroup"
	case *ast.Field:
		nodeType = "Field"
	case *ast.FieldList:
		nodeType = "FieldList"
	case *ast.BadExpr:
		nodeType = "BadExpression"
	case *ast.Ident:
		nodeType = "Identifier"
	case *ast.BasicLit:
		nodeType = "BasicLiteral"
	case *ast.Ellipsis:
		nodeType = "Ellipsis"
	case *ast.FuncLit:
		nodeType = "FunctionLiteral"
	case *ast.CompositeLit:
		nodeType = "CompositeLiteral"
	case *ast.ParenExpr:
		nodeType = "ParenthesizedExpression"
	case *ast.SelectorExpr:
		nodeType = "SelectorExpression"
	case *ast.IndexExpr:
		nodeType = "IndexExpression"
	case *ast.IndexListExpr:
		nodeType = "IndexListExpression"
	case *ast.SliceExpr:
		nodeType = "SliceExpression"
	case *ast.TypeAssertExpr:
		nodeType = "TypeAssertExpression"
	case *ast.CallExpr:
		nodeType = "CallExpression"
	case *ast.StarExpr:
		nodeType = "StarExpression"
	case *ast.UnaryExpr:
		nodeType = "UnaryExpression"
	case *ast.BinaryExpr:
		nodeType = "BinaryExpression"
	case *ast.KeyValueExpr:
		nodeType = "KeyValueExpression"
	case *ast.ArrayType:
		nodeType = "ArrayType"
	case *ast.StructType:
		nodeType = "StructType"
	case *ast.FuncType:
		nodeType = "FunctionType"
	case *ast.InterfaceType:
		nodeType = "InterfaceType"
	case *ast.MapType:
		nodeType = "MapType"
	case *ast.ChanType:
		nodeType = "ChanelType"
	case *ast.BadStmt:
		nodeType = "BadStatement"
	case *ast.DeclStmt:
		nodeType = "DeclarationStatement"
	case *ast.EmptyStmt:
		nodeType = "EmptyStatement"
	case *ast.LabeledStmt:
		nodeType = "LabeledStatement"
	case *ast.ExprStmt:
		nodeType = "ExpressionStatement"
	case *ast.SendStmt:
		nodeType = "SendStatement"
	case *ast.IncDecStmt:
		nodeType = "IncrementDecrementStatement"
	case *ast.AssignStmt:
		nodeType = "AssignStatement"
	case *ast.GoStmt:
		nodeType = "GoStatement"
	case *ast.DeferStmt:
		nodeType = "DeferStatement"
	case *ast.ReturnStmt:
		nodeType = "ReturnStatement"
	case *ast.BranchStmt:
		nodeType = "BranchStatement"
	case *ast.BlockStmt:
		nodeType = "BlockStatement"
	case *ast.IfStmt:
		nodeType = "IfStatement"
	case *ast.CaseClause:
		nodeType = "CaseClause"
	case *ast.SwitchStmt:
		nodeType = "SwitchStatement"
	case *ast.TypeSwitchStmt:
		nodeType = "TypeSwitchStatement"
	case *ast.CommClause:
		nodeType = "CommClause"
	case *ast.SelectStmt:
		nodeType = "SelectStatement"
	case *ast.ForStmt:
		nodeType = "ForStatement"
	case *ast.RangeStmt:
		nodeType = "RangeStatement"
	// Declarations
	case *ast.ImportSpec:
		nodeType = "ImportSpecification"
	case *ast.ValueSpec:
		nodeType = "ValueSpecification"
	case *ast.TypeSpec:
		nodeType = "TypeSpecification"
	case *ast.BadDecl:
		nodeType = "BadDeclaration"
	case *ast.GenDecl:
		nodeType = "GenericDeclaration"
	case *ast.FuncDecl:
		nodeType = "FunctionDeclaration"
	case *ast.File:
		nodeType = "File"
	case *ast.Package:
		nodeType = "Package"
	}

	stringWriter := bytes.NewBufferString("")
	err := printer.Fprint(stringWriter, token.NewFileSet(), n)

	var id, uuidErr = uuid.NewUUID()
	var idAsString = fmt.Sprintf("%v", time.Now().Unix())

	if uuidErr == nil {
		idAsString = id.String()
	}

	if err == nil {
		return data.Node{
			Id:            idAsString,
			Children:      []data.INode{},
			StartPosition: int(n.Pos()),
			EndPosition:   int(n.End()),
			Code:          stringWriter.String(),
			NodeType:      nodeType,
		}
	}

	return data.Node{
		Id:            idAsString,
		Children:      []data.INode{},
		StartPosition: int(n.Pos()),
		EndPosition:   int(n.End()),
		NodeType:      nodeType,
	}
}
