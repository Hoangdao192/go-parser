package parser

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	data "joern-go/parser/ast"
	"joern-go/parser/ast/expression"
	"joern-go/parser/ast/statement"
	"joern-go/parser/ast/types"
	"joern-go/util"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func Parse(sourceDir string, destDir string) {
	_, openFileErr := os.OpenFile(destDir, os.O_RDONLY, os.ModePerm)
	if openFileErr != nil {
		if mkdirErr := os.MkdirAll(destDir, os.ModePerm); mkdirErr != nil {
			log.Fatal(mkdirErr)
		}
	}

	err := filepath.Walk(sourceDir, func(filepath string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if isGoFile(info) {
			os.MkdirAll(path.Dir(filepath), os.ModePerm)

			fileContent, readErr := util.ReadFile(filepath)
			if readErr == nil {
				parsedFile, parseErr := parser.ParseFile(token.NewFileSet(), info.Name(),
					fileContent, parser.AllErrors)
				if parseErr == nil {
					jsonData, jsonErr := json.Marshal(parsedFile)
					if jsonErr == nil {
						saveFilePath := outDir + "/" + filepath[strings.Index(
							filepath, projectDir):] + ".json"
						saveFile, openFileErr := os.OpenFile(saveFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
							os.ModePerm)
						if openFileErr != nil {
							log.Fatal(openFileErr)
						}
						saveFile.Write(jsonData)
					}
				} else {
					log.Fatal(parseErr)
				}
			}
		}
		return nil
	})

	if err != nil {
		fmt.Println(err)
	}

	fs := token.NewFileSet()
	f, err := parser.ParseFile(fs, "", "package main; var a = 3", parser.DeclarationErrors)
	if err != nil {
		log.Fatal(err)
	}
	var v visitor
	ast.Walk(v, f)
}

func isGoFile(info fs.FileInfo) bool {
	return !info.IsDir() &&
		strings.LastIndex(info.Name(), ".go")+len(".go") == len(info.Name())
}

func parse(parent data.INode, node ast.Node) data.INode {
	if node == nil {
		return nil
	}

	var parsedNode data.INode

	// walk children
	// (the order of the cases matches the order
	// of the corresponding node types in ast.go)
	switch n := node.(type) {
	// Comments and fields
	case *ast.Comment:
		parsedNode = BuildComment(node.(*ast.Comment))

	case *ast.CommentGroup:
		parsedNode = BuildCommentGroup(node.(*ast.CommentGroup))
		for _, c := range n.List {
			childNode := parse(parsedNode, c)
			parsedNode.(data.Node).AddChild(childNode)
		}

	case *ast.Field:
		parsedNode := BuildField(node.(*ast.Field))
		if n.Doc != nil {
			var docNode = parse(parsedNode, n.Doc)
			parsedNode.Doc = docNode.(data.CommentGroup)
		}

		for _, x := range n.Names {
			var identifierNode = parse(parsedNode, x)
			parsedNode.Names = append(parsedNode.Names, identifierNode.(expression.Identifier))
		}

		if n.Type != nil {
			var typeNode = parse(parsedNode, n.Type)
			parsedNode.Type = typeNode.(expression.Expression)
		}
		if n.Tag != nil {
			tagNode := parse(parsedNode, n.Tag)
			parsedNode.Tag = tagNode.(expression.BasicLiteral)
		}
		if n.Comment != nil {
			commentNode := parse(parsedNode, n.Comment)
			parsedNode.Comment = commentNode.(data.CommentGroup)
		}

	case *ast.FieldList:
		parsedNode := BuildFieldList(n)
		for _, f := range n.List {
			fieldNode := parse(parsedNode, f)
			parsedNode.List = append(parsedNode.List, fieldNode.(data.Field))
		}

	case *ast.BadExpr:
		parsedNode := BuildBadExpression(n)

	case *ast.Ident:
		parsedNode := BuildIdentifier(n)

	case *ast.BasicLit:
		parsedNode := BuildBasicLiteral(n)

	case *ast.Ellipsis:
		parsedNode := BuildEllipsis(n)
		if n.Elt != nil {
			elementNode := parse(parsedNode, n.Elt)
			parsedNode.Element = elementNode.(expression.Expression)
		}

	case *ast.FuncLit:
		parsedNode := BuildFunctionLiteral(n)
		parsedNode.Type = parse(parsedNode, n.Type).(types.FunctionType)
		parsedNode.Body = parse(parsedNode, n.Body).(statement.BlockStatement)

	case *ast.CompositeLit:
		parsedNode := BuildCompositeLiteral(n)
		if n.Type != nil {
			parsedNode.Type = parse(parsedNode, n.Type).(expression.Expression)
		}
		for _, elt := range n.Elts {
			elementNode := parse(parsedNode, elt)
			parsedNode.Elements = append(parsedNode.Elements, elementNode.(expression.Expression))
		}

	case *ast.ParenExpr:
		parsedNode := BuildParenthesizedExpression(n)
		expressionNode := parse(parsedNode, n.X)
		parsedNode.Expression = expressionNode.(expression.Expression)

	case *ast.SelectorExpr:
		parsedNode := expression.SelectorExpression{
			Node: BuildNode(n),
		}
		expressionNode := parse(parsedNode, n.X)
		parsedNode.Expression = expressionNode.(expression.Expression)
		identifierNode := parse(parsedNode, n.Sel)
		parsedNode.Sel = identifierNode.(expression.Identifier)

	case *ast.IndexExpr:
		parsedNode := expression.IndexExpression{
			Node:         BuildNode(n),
			Expression:   parse(parsedNode, n.X).(expression.Expression),
			LeftBracket:  int(n.Lbrack),
			Index:        parse(parsedNode, n.Index).(expression.Expression),
			RightBracket: int(n.Rbrack),
		}

	case *ast.IndexListExpr:
		parsedNode := expression.IndexListExpression{
			Node:         BuildNode(n),
			Expression:   parse(parsedNode, n.X).(expression.Expression),
			LeftBracket:  int(n.Lbrack),
			Indices:      []expression.Expression{},
			RightBracket: int(n.Rbrack),
		}
		for _, index := range n.Indices {
			parsedNode.Indices = append(
				parsedNode.Indices, parse(parsedNode, index).(expression.Expression))
		}

	case *ast.SliceExpr:
		parsedNode := expression.SliceExpression{
			Node:         BuildNode(n),
			Expression:   parse(parsedNode, n.X).(expression.Expression),
			LeftBracket:  int(n.Lbrack),
			Slice3:       n.Slice3,
			RightBracket: int(n.Rbrack),
		}
		if n.Low != nil {
			parsedNode.Low = parse(parsedNode, n.Low).(expression.Expression)
		}
		if n.High != nil {
			parsedNode.High = parse(parsedNode, n.High).(expression.Expression)
		}
		if n.Max != nil {
			parsedNode.Max = parse(parsedNode, n.Max).(expression.Expression)
		}

	case *ast.TypeAssertExpr:
		parsedNode := expression.TypeAssertExpression{
			Node:       BuildNode(n),
			Expression: parse(parsedNode, n.X).(expression.Expression),
			Lparen:     int(n.Lparen),
			Rparen:     int(n.Rparen),
		}
		if n.Type != nil {
			parsedNode.Type = parse(parsedNode, n.Type).(expression.Expression)
		}

	case *ast.CallExpr:
		parsedNode := expression.CallExpression{
			Node:     BuildNode(n),
			Function: parse(parsedNode, n.Fun).(expression.Expression),
			Args:     []expression.Expression{},
		}
		for _, arg := range n.Args {
			parsedNode.Args = append(
				parsedNode.Args, parse(parsedNode, arg).(expression.Expression))
		}

	case *ast.StarExpr:
		parsedNode := expression.StarExpression{
			Node:       BuildNode(n),
			Star:       int(n.Star),
			Expression: parse(parsedNode, n.X).(expression.Expression),
		}

	case *ast.UnaryExpr:
		parsedNode := expression.UnaryExpression{
			Node:       BuildNode(n),
			OpPos:      int(n.OpPos),
			Op:         int(n.Op),
			Expression: parse(parsedNode, n.X).(expression.Expression),
		}

	case *ast.BinaryExpr:
		parsedNode := expression.BinaryExpression{
			Node:            BuildNode(n),
			LeftExpression:  parse(parsedNode, n.X).(expression.Expression),
			OpPos:           int(n.OpPos),
			Op:              int(n.Op),
			RightExpression: parse(parsedNode, n.Y).(expression.Expression),
		}

	case *ast.KeyValueExpr:
		parsedNode := expression.KeyValueExpression{
			Node:  BuildNode(n),
			Key:   parse(parsedNode, n.Key).(expression.Expression),
			Value: parse(parsedNode, n.Value).(expression.Expression),
		}

	// Types
	case *ast.ArrayType:
		parsedNode := types.ArrayType{
			Node:        BuildNode(n),
			LeftBracket: int(n.Lbrack),
			Element:     parse(parsedNode, n.Elt).(expression.Expression),
		}
		if n.Len != nil {
			parsedNode.Length = parse(parsedNode, n.Len).(expression.Expression)
		}

	case *ast.StructType:
		parsedNode := types.StructType{
			Node:       BuildNode(n),
			Struct:     int(n.Struct),
			Fields:     parse(parsedNode, n.Fields).(data.FieldList),
			Incomplete: n.Incomplete,
		}

	case *ast.FuncType:
		if n.TypeParams != nil {
			Walk(v, n.TypeParams)
		}
		if n.Params != nil {
			Walk(v, n.Params)
		}
		if n.Results != nil {
			Walk(v, n.Results)
		}

	case *ast.InterfaceType:
		Walk(v, n.Methods)

	case *ast.MapType:
		Walk(v, n.Key)
		Walk(v, n.Value)

	case *ast.ChanType:
		Walk(v, n.Value)

	// Statements
	case *ast.BadStmt:
		// nothing to do

	case *ast.DeclStmt:
		Walk(v, n.Decl)

	case *ast.EmptyStmt:
		// nothing to do

	case *ast.LabeledStmt:
		Walk(v, n.Label)
		Walk(v, n.Stmt)

	case *ast.ExprStmt:
		expressionNode := parse(parsedNode, n.X)

	case *ast.SendStmt:
		Walk(v, n.Chan)
		Walk(v, n.Value)

	case *ast.IncDecStmt:
		expressionNode := parse(parsedNode, n.X)

	case *ast.AssignStmt:
		walkExprList(v, n.Lhs)
		walkExprList(v, n.Rhs)

	case *ast.GoStmt:
		Walk(v, n.Call)

	case *ast.DeferStmt:
		Walk(v, n.Call)

	case *ast.ReturnStmt:
		walkExprList(v, n.Results)

	case *ast.BranchStmt:
		if n.Label != nil {
			Walk(v, n.Label)
		}

	case *ast.BlockStmt:
		walkStmtList(v, n.List)

	case *ast.IfStmt:
		if n.Init != nil {
			Walk(v, n.Init)
		}
		Walk(v, n.Cond)
		Walk(v, n.Body)
		if n.Else != nil {
			Walk(v, n.Else)
		}

	case *ast.CaseClause:
		walkExprList(v, n.List)
		walkStmtList(v, n.Body)

	case *ast.SwitchStmt:
		if n.Init != nil {
			Walk(v, n.Init)
		}
		if n.Tag != nil {
			Walk(v, n.Tag)
		}
		Walk(v, n.Body)

	case *ast.TypeSwitchStmt:
		if n.Init != nil {
			Walk(v, n.Init)
		}
		Walk(v, n.Assign)
		Walk(v, n.Body)

	case *ast.CommClause:
		if n.Comm != nil {
			Walk(v, n.Comm)
		}
		walkStmtList(v, n.Body)

	case *ast.SelectStmt:
		Walk(v, n.Body)

	case *ast.ForStmt:
		if n.Init != nil {
			Walk(v, n.Init)
		}
		if n.Cond != nil {
			Walk(v, n.Cond)
		}
		if n.Post != nil {
			Walk(v, n.Post)
		}
		Walk(v, n.Body)

	case *ast.RangeStmt:
		if n.Key != nil {
			Walk(v, n.Key)
		}
		if n.Value != nil {
			Walk(v, n.Value)
		}
		expressionNode := parse(parsedNode, n.X)
		Walk(v, n.Body)

	// Declarations
	case *ast.ImportSpec:
		if n.Doc != nil {
			Walk(v, n.Doc)
		}
		if n.Name != nil {
			Walk(v, n.Name)
		}
		Walk(v, n.Path)
		if n.Comment != nil {
			Walk(v, n.Comment)
		}

	case *ast.ValueSpec:
		if n.Doc != nil {
			Walk(v, n.Doc)
		}
		walkIdentList(v, n.Names)
		if n.Type != nil {
			Walk(v, n.Type)
		}
		walkExprList(v, n.Values)
		if n.Comment != nil {
			Walk(v, n.Comment)
		}

	case *ast.TypeSpec:
		if n.Doc != nil {
			Walk(v, n.Doc)
		}
		Walk(v, n.Name)
		if n.TypeParams != nil {
			Walk(v, n.TypeParams)
		}
		Walk(v, n.Type)
		if n.Comment != nil {
			Walk(v, n.Comment)
		}

	case *ast.BadDecl:
		// nothing to do

	case *ast.GenDecl:
		if n.Doc != nil {
			Walk(v, n.Doc)
		}
		for _, s := range n.Specs {
			Walk(v, s)
		}

	case *ast.FuncDecl:
		if n.Doc != nil {
			Walk(v, n.Doc)
		}
		if n.Recv != nil {
			Walk(v, n.Recv)
		}
		Walk(v, n.Name)
		Walk(v, n.Type)
		if n.Body != nil {
			Walk(v, n.Body)
		}

	// Files and packages
	case *ast.File:
		if n.Doc != nil {
			Walk(v, n.Doc)
		}
		Walk(v, n.Name)
		walkDeclList(v, n.Decls)
		// don't walk n.Comments - they have been
		// visited already through the individual
		// nodes

	case *ast.Package:
		for _, f := range n.Files {
			Walk(v, f)
		}

	default:
		panic(fmt.Sprintf("ast.Walk: unexpected node type %T", n))
	}

	v.Visit(nil)
}
