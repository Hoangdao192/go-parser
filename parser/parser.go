package parser

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"joern-go/util"
	"log"
	"os"
	"path"
	"path/filepath"

	//"go/parser"
	//"go/token"
	"io/fs"
	data "joern-go/parser/ast"
	//"joern-go/util"
	//"log"
	//"os"
	//"path"
	//"path/filepath"
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
			os.MkdirAll(destDir+"/"+path.Dir(filepath)[len(sourceDir):], os.ModePerm)

			fileContent, readErr := util.ReadFile(filepath)
			if readErr == nil {
				parsedFile, parseErr := parser.ParseFile(token.NewFileSet(), info.Name(),
					fileContent, parser.AllErrors)
				if parseErr == nil {
					fileNode := data.File{}
					jsonData, jsonErr := json.Marshal(parse(&fileNode, parsedFile))
					if jsonErr == nil {
						saveFilePath := destDir + "/" + filepath[len(sourceDir):] + ".json"
						saveFile, openFileErr := os.OpenFile(saveFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
							os.ModePerm)
						if openFileErr != nil {
							log.Fatal(openFileErr)
						}
						saveFile.Write(jsonData)
					} else {
						log.Fatal(jsonErr)
					}
				} else {
					log.Fatal(parseErr)
				}

				var v visitor
				ast.Walk(v, parsedFile)
			}
		}
		return nil
	})

	if err != nil {
		fmt.Println(err)
	}
}

type visitor int

func (v visitor) Visit(node ast.Node) (w ast.Visitor) {
	if node == nil {
		return nil
	}
	fmt.Printf("%s%T\n", strings.Repeat("\t", int(v)), node)
	return v + 1
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
		parsedNode := BuildComment(node.(*ast.Comment))
		return &parsedNode

	case *ast.CommentGroup:
		parsedNode := BuildCommentGroup(node.(*ast.CommentGroup))
		for _, c := range n.List {
			childNode := parse(&parsedNode, c)
			parsedNode.AddChild(childNode)
		}
		return &parsedNode

	case *ast.Field:
		parsedNode := BuildField(node.(*ast.Field))
		if n.Doc != nil {
			var docNode = parse(&parsedNode, n.Doc)
			parsedNode.Doc = docNode.(*data.CommentGroup)
		}

		for _, x := range n.Names {
			var identifierNode = parse(&parsedNode, x)
			parsedNode.Names = append(parsedNode.Names, identifierNode.(*data.Identifier))
		}

		if n.Type != nil {
			var typeNode = parse(&parsedNode, n.Type)
			parsedNode.Type = typeNode.(data.Expression)
		}
		if n.Tag != nil {
			tagNode := parse(&parsedNode, n.Tag)
			parsedNode.Tag = tagNode.(*data.BasicLiteral)
		}
		if n.Comment != nil {
			commentNode := parse(&parsedNode, n.Comment)
			parsedNode.Comment = commentNode.(*data.CommentGroup)
		}
		return &parsedNode

	case *ast.FieldList:
		parsedNode := BuildFieldList(n)
		for _, f := range n.List {
			fieldNode := parse(&parsedNode, f)
			parsedNode.List = append(parsedNode.List, fieldNode.(*data.Field))
		}
		return &parsedNode

	case *ast.BadExpr:
		parsedNode := BuildBadExpression(n)
		return &parsedNode

	case *ast.Ident:
		parsedNode := BuildIdentifier(n)
		return &parsedNode

	case *ast.BasicLit:
		parsedNode := BuildBasicLiteral(n)
		return &parsedNode

	case *ast.Ellipsis:
		parsedNode := BuildEllipsis(n)
		if n.Elt != nil {
			elementNode := parse(&parsedNode, n.Elt)
			parsedNode.Element = elementNode.(data.Expression)
		}
		return &parsedNode

	case *ast.FuncLit:
		parsedNode := BuildFunctionLiteral(n)
		parsedNode.Type = parse(&parsedNode, n.Type).(*data.FunctionType)
		parsedNode.Body = parse(&parsedNode, n.Body).(*data.BlockStatement)
		return &parsedNode

	case *ast.CompositeLit:
		parsedNode := BuildCompositeLiteral(n)
		if n.Type != nil {
			parsedNode.Type = parse(&parsedNode, n.Type).(data.Expression)
		}
		for _, elt := range n.Elts {
			elementNode := parse(&parsedNode, elt)
			parsedNode.Elements = append(parsedNode.Elements, elementNode.(data.Expression))
		}
		return &parsedNode

	case *ast.ParenExpr:
		parsedNode := BuildParenthesizedExpression(n)
		expressionNode := parse(&parsedNode, n.X)
		parsedNode.Expression = expressionNode.(data.Expression)
		return &parsedNode

	case *ast.SelectorExpr:
		parsedNode := data.SelectorExpression{
			Node: BuildNode(n),
		}
		expressionNode := parse(&parsedNode, n.X)
		parsedNode.Expression = expressionNode.(data.Expression)
		identifierNode := parse(&parsedNode, n.Sel)
		parsedNode.Sel = identifierNode.(*data.Identifier)
		return &parsedNode

	case *ast.IndexExpr:
		parsedNode := data.IndexExpression{
			Node:         BuildNode(n),
			Expression:   parse(parsedNode, n.X).(data.Expression),
			LeftBracket:  int(n.Lbrack),
			Index:        parse(parsedNode, n.Index).(data.Expression),
			RightBracket: int(n.Rbrack),
		}
		return &parsedNode

	case *ast.IndexListExpr:
		parsedNode := data.IndexListExpression{
			Node:         BuildNode(n),
			Expression:   parse(parsedNode, n.X).(data.Expression),
			LeftBracket:  int(n.Lbrack),
			Indices:      []data.Expression{},
			RightBracket: int(n.Rbrack),
		}
		for _, index := range n.Indices {
			parsedNode.Indices = append(
				parsedNode.Indices, parse(&parsedNode, index).(data.Expression))
		}
		return &parsedNode

	case *ast.SliceExpr:
		parsedNode := data.SliceExpression{
			Node:         BuildNode(n),
			Expression:   parse(parsedNode, n.X).(data.Expression),
			LeftBracket:  int(n.Lbrack),
			Slice3:       n.Slice3,
			RightBracket: int(n.Rbrack),
		}
		if n.Low != nil {
			parsedNode.Low = parse(&parsedNode, n.Low).(data.Expression)
		}
		if n.High != nil {
			parsedNode.High = parse(&parsedNode, n.High).(data.Expression)
		}
		if n.Max != nil {
			parsedNode.Max = parse(&parsedNode, n.Max).(data.Expression)
		}
		return &parsedNode

	case *ast.TypeAssertExpr:
		parsedNode := data.TypeAssertExpression{
			Node:       BuildNode(n),
			Expression: parse(parsedNode, n.X).(data.Expression),
			Lparen:     int(n.Lparen),
			Rparen:     int(n.Rparen),
		}
		if n.Type != nil {
			parsedNode.Type = parse(&parsedNode, n.Type).(data.Expression)
		}
		return &parsedNode

	case *ast.CallExpr:
		parsedNode := data.CallExpression{
			Node:     BuildNode(n),
			Function: parse(parsedNode, n.Fun).(data.Expression),
			Args:     []data.Expression{},
		}
		for _, arg := range n.Args {
			parsedNode.Args = append(
				parsedNode.Args, parse(&parsedNode, arg).(data.Expression))
		}
		return &parsedNode

	case *ast.StarExpr:
		parsedNode := data.StarExpression{
			Node:       BuildNode(n),
			Star:       int(n.Star),
			Expression: parse(parsedNode, n.X).(data.Expression),
		}
		return &parsedNode

	case *ast.UnaryExpr:
		parsedNode := data.UnaryExpression{
			Node:       BuildNode(n),
			OpPos:      int(n.OpPos),
			Op:         int(n.Op),
			Expression: parse(parsedNode, n.X).(data.Expression),
		}
		return &parsedNode

	case *ast.BinaryExpr:
		parsedNode := data.BinaryExpression{
			Node:            BuildNode(n),
			LeftExpression:  parse(parsedNode, n.X).(data.Expression),
			OpPos:           int(n.OpPos),
			Op:              int(n.Op),
			RightExpression: parse(parsedNode, n.Y).(data.Expression),
		}
		return &parsedNode

	case *ast.KeyValueExpr:
		parsedNode := data.KeyValueExpression{
			Node:  BuildNode(n),
			Key:   parse(parsedNode, n.Key).(data.Expression),
			Value: parse(parsedNode, n.Value).(data.Expression),
		}
		return &parsedNode

	// Types
	case *ast.ArrayType:
		parsedNode := data.ArrayType{
			Node:        BuildNode(n),
			LeftBracket: int(n.Lbrack),
			Element:     parse(parsedNode, n.Elt).(data.Expression),
		}
		if n.Len != nil {
			parsedNode.Length = parse(&parsedNode, n.Len).(data.Expression)
		}
		return &parsedNode

	case *ast.StructType:
		parsedNode := data.StructType{
			Node:       BuildNode(n),
			Struct:     int(n.Struct),
			Fields:     parse(parsedNode, n.Fields).(*data.FieldList),
			Incomplete: n.Incomplete,
		}
		return &parsedNode

	case *ast.FuncType:
		parsedNode := data.FunctionType{
			Node:     BuildNode(n),
			Function: int(n.Func),
		}
		if n.TypeParams != nil {
			parsedNode.TypeParams = parse(&parsedNode, n.TypeParams).(*data.FieldList)
		}
		if n.Params != nil {
			parsedNode.Params = parse(&parsedNode, n.Params).(*data.FieldList)
		}
		if n.Results != nil {
			parsedNode.Results = parse(&parsedNode, n.Results).(*data.FieldList)
		}
		return &parsedNode

	case *ast.InterfaceType:
		parsedNode := data.InterfaceType{
			Node:       BuildNode(n),
			Interface:  int(n.Interface),
			Incomplete: n.Incomplete,
		}
		parsedNode.Methods = parse(&parsedNode, n.Methods).(*data.FieldList)
		return &parsedNode

	case *ast.MapType:
		parsedNode := data.MapType{
			Node: BuildNode(n),
			Map:  int(n.Map),
		}
		parsedNode.Key = parse(&parsedNode, n.Key).(data.Expression)
		parsedNode.Value = parse(&parsedNode, n.Value).(data.Expression)
		return &parsedNode

	case *ast.ChanType:
		parsedNode := data.ChanelType{
			Node:      BuildNode(n),
			Begin:     int(n.Begin),
			Arrow:     int(n.Arrow),
			Direction: int(n.Dir),
		}
		parsedNode.Value = parse(&parsedNode, n.Value).(data.Expression)
		return &parsedNode

	// Statements
	case *ast.BadStmt:
		parsedNode := data.BadStatement{
			Node: BuildNode(n),
			From: int(n.From),
			To:   int(n.To),
		}
		return &parsedNode

	case *ast.DeclStmt:
		parsedNode := data.DeclarationStatement{
			Node: BuildNode(n),
		}
		parsedNode.Declaration = parse(&parsedNode, n.Decl).(data.Declaration)
		return &parsedNode

	case *ast.EmptyStmt:
		parsedNode := data.EmptyStatement{
			Node:      BuildNode(n),
			Semicolon: int(n.Semicolon),
			Implicit:  n.Implicit,
		}
		return &parsedNode

	case *ast.LabeledStmt:
		parsedNode := data.LabeledStatement{
			Node:  BuildNode(n),
			Colon: int(n.Colon),
		}
		parsedNode.Label = parse(&parsedNode, n.Label).(*data.Identifier)
		parsedNode.Statement = parse(&parsedNode, n.Stmt).(data.Statement)
		return &parsedNode

	case *ast.ExprStmt:
		parsedNode := data.ExpressionStatement{
			Node: BuildNode(n),
		}
		parsedNode.Expression = parse(&parsedNode, n.X).(data.Expression)
		return &parsedNode

	case *ast.SendStmt:
		parsedNode := data.SendStatement{
			Node:  BuildNode(n),
			Arrow: int(n.Arrow),
		}
		parsedNode.Chanel = parse(&parsedNode, n.Chan).(data.Expression)
		parsedNode.Value = parse(&parsedNode, n.Value).(data.Expression)
		return &parsedNode

	case *ast.IncDecStmt:
		parsedNode := data.IncrementDecrementStatement{
			Node:          BuildNode(n),
			TokenPosition: int(n.TokPos),
			Token:         int(n.Tok),
		}
		parsedNode.Expression = parse(&parsedNode, n.X).(data.Expression)
		return &parsedNode

	case *ast.AssignStmt:
		parsedNode := data.AssignStatement{
			Node:     BuildNode(n),
			Lhs:      []data.Expression{},
			TokenPos: int(n.TokPos),
			Token:    int(n.Tok),
			Rhs:      []data.Expression{},
		}
		for _, expr := range n.Lhs {
			parsedNode.Lhs = append(parsedNode.Lhs, parse(&parsedNode, expr).(data.Expression))
		}
		for _, expr := range n.Rhs {
			parsedNode.Rhs = append(parsedNode.Rhs, parse(&parsedNode, expr).(data.Expression))
		}
		return &parsedNode

	case *ast.GoStmt:
		parsedNode := data.GoStatement{
			Node: BuildNode(n),
			Go:   int(n.Go),
		}
		parsedNode.Call = parse(&parsedNode, n.Call).(*data.CallExpression)
		return &parsedNode

	case *ast.DeferStmt:
		parsedNode := data.DeferStatement{
			Node:  BuildNode(n),
			Defer: int(n.Defer),
		}
		parsedNode.Call = parse(&parsedNode, n.Call).(*data.CallExpression)
		return &parsedNode

	case *ast.ReturnStmt:
		parsedNode := data.ReturnStatement{
			Node:    BuildNode(n),
			Return:  int(n.Return),
			Results: []data.Expression{},
		}
		for _, expr := range n.Results {
			parsedNode.Results = append(parsedNode.Results, parse(&parsedNode, expr).(data.Expression))
		}
		return &parsedNode

	case *ast.BranchStmt:
		parsedNode := data.BranchStatement{
			Node:          BuildNode(n),
			TokenPosition: int(n.TokPos),
			Token:         int(n.Tok),
		}
		if n.Label != nil {
			parsedNode.Label = parse(&parsedNode, n.Label).(*data.Identifier)
		}
		return &parsedNode

	case *ast.BlockStmt:
		parsedNode := data.BlockStatement{
			Node:   BuildNode(n),
			Lbrace: int(n.Lbrace),
			List:   []data.Statement{},
			Rbrace: int(n.Rbrace),
		}
		for _, stmt := range n.List {
			parsedNode.List = append(parsedNode.List, parse(&parsedNode, stmt).(data.Statement))
		}
		return &parsedNode

	case *ast.IfStmt:
		parsedNode := data.IfStatement{
			Node: BuildNode(n),
			If:   int(n.If),
		}
		if n.Init != nil {
			parsedNode.Initialization = parse(&parsedNode, n.Init).(data.Statement)
		}
		parsedNode.Condition = parse(&parsedNode, n.Cond).(data.Expression)
		parsedNode.Body = parse(&parsedNode, n.Body).(*data.BlockStatement)
		if n.Else != nil {
			parsedNode.Else = parse(&parsedNode, n.Else).(data.Statement)
		}
		return &parsedNode

	case *ast.CaseClause:
		parsedNode := data.CaseClause{
			Node:  BuildNode(n),
			Case:  int(n.Case),
			List:  []data.Expression{},
			Colon: int(n.Colon),
			Body:  []data.Statement{},
		}
		for _, expr := range n.List {
			parsedNode.List = append(parsedNode.List, parse(&parsedNode, expr).(data.Expression))
		}
		for _, stmt := range n.Body {
			parsedNode.Body = append(parsedNode.Body, parse(&parsedNode, stmt).(data.Statement))
		}
		return &parsedNode

	case *ast.SwitchStmt:
		parsedNode := data.SwitchStatement{
			Node:   BuildNode(n),
			Switch: int(n.Switch),
		}
		if n.Init != nil {
			parsedNode.Initialization = parse(&parsedNode, n.Init).(data.Statement)
		}
		if n.Tag != nil {
			parsedNode.Tag = parse(&parsedNode, n.Tag).(data.Expression)
		}
		parsedNode.Body = parse(&parsedNode, n.Body).(*data.BlockStatement)
		return &parsedNode

	case *ast.TypeSwitchStmt:
		parsedNode := data.TypeSwitchStatement{
			Node:   BuildNode(n),
			Switch: int(n.Switch),
		}
		if n.Init != nil {
			parsedNode.Initialization = parse(&parsedNode, n.Init).(data.Statement)
		}
		parsedNode.Assign = parse(&parsedNode, n.Assign).(data.Statement)
		parsedNode.Body = parse(&parsedNode, n.Body).(*data.BlockStatement)
		return &parsedNode

	case *ast.CommClause:
		parsedNode := data.CommClause{
			Node:  BuildNode(n),
			Case:  int(n.Case),
			Colon: int(n.Colon),
			Body:  []data.Statement{},
		}
		if n.Comm != nil {
			parsedNode.Comm = parse(&parsedNode, n.Comm).(data.Statement)
		}
		for _, stmt := range n.Body {
			parsedNode.Body = append(parsedNode.Body, parse(&parsedNode, stmt).(data.Statement))
		}
		return &parsedNode

	case *ast.SelectStmt:
		parsedNode := data.SelectStatement{
			Node:   BuildNode(n),
			Select: int(n.Select),
		}
		parsedNode.Body = parse(&parsedNode, n.Body).(*data.BlockStatement)
		return &parsedNode

	case *ast.ForStmt:
		parsedNode := data.ForStatement{
			Node: BuildNode(n),
			For:  int(n.For),
		}
		if n.Init != nil {
			parsedNode.Initialization = parse(&parsedNode, n.Init).(data.Statement)
		}
		if n.Cond != nil {
			parsedNode.Condition = parse(&parsedNode, n.Cond).(data.Expression)
		}
		if n.Post != nil {
			parsedNode.Post = parse(&parsedNode, n.Post).(data.Statement)
		}
		parsedNode.Body = parse(&parsedNode, n.Body).(*data.BlockStatement)
		return &parsedNode

	case *ast.RangeStmt:
		parsedNode := data.RangeStatement{
			Node:     BuildNode(n),
			For:      int(n.For),
			TokenPos: int(n.TokPos),
			Token:    int(n.Tok),
			Range:    int(n.Range),
		}
		if n.Key != nil {
			parsedNode.Key = parse(&parsedNode, n.Key).(data.Expression)
		}
		if n.Value != nil {
			parsedNode.Value = parse(&parsedNode, n.Value).(data.Expression)
		}
		parsedNode.Expression = parse(&parsedNode, n.X).(data.Expression)
		parsedNode.Body = parse(&parsedNode, n.Body).(*data.BlockStatement)
		return &parsedNode

	// Declarations
	case *ast.ImportSpec:
		parsedNode := data.ImportSpecification{
			Node:        BuildNode(n),
			EndPosition: int(n.EndPos),
		}
		if n.Doc != nil {
			parsedNode.Doc = parse(&parsedNode, n.Doc).(*data.CommentGroup)
		}
		if n.Name != nil {
			parsedNode.Name = parse(&parsedNode, n.Name).(*data.Identifier)
		}
		parsedNode.Path = parse(&parsedNode, n.Path).(*data.BasicLiteral)
		if n.Comment != nil {
			parsedNode.Comment = parse(&parsedNode, n.Comment).(*data.CommentGroup)
		}
		return &parsedNode

	case *ast.ValueSpec:
		parsedNode := data.ValueSpecification{
			Node:   BuildNode(n),
			Names:  []*data.Identifier{},
			Values: []data.Expression{},
		}
		if n.Doc != nil {
			parsedNode.Doc = parse(&parsedNode, n.Doc).(*data.CommentGroup)
		}
		for _, ident := range n.Names {
			parsedNode.Names = append(parsedNode.Names, parse(&parsedNode, ident).(*data.Identifier))
		}
		if n.Type != nil {
			parsedNode.Type = parse(&parsedNode, n.Type).(data.Expression)
		}
		for _, expr := range n.Values {
			parsedNode.Values = append(parsedNode.Values, parse(&parsedNode, expr).(data.Expression))
		}
		if n.Comment != nil {
			parsedNode.Comment = parse(&parsedNode, n.Comment).(*data.CommentGroup)
		}
		return &parsedNode

	case *ast.TypeSpec:
		parsedNode := data.TypeSpecification{
			Node:   BuildNode(n),
			Assign: int(n.Assign),
		}
		if n.Doc != nil {
			parsedNode.Doc = parse(&parsedNode, n.Doc).(*data.CommentGroup)
		}
		parsedNode.Name = parse(&parsedNode, n.Name).(*data.Identifier)
		if n.TypeParams != nil {
			parsedNode.TypeParams = parse(&parsedNode, n.TypeParams).(*data.FieldList)
		}
		parsedNode.Type = parse(&parsedNode, n.Type).(data.Expression)
		if n.Comment != nil {
			parsedNode.Comment = parse(&parsedNode, n.Comment).(*data.CommentGroup)
		}
		return &parsedNode

	case *ast.BadDecl:
		parsedNode := data.BadDeclaration{
			Node: BuildNode(n),
			From: int(n.From),
			To:   int(n.To),
		}
		// nothing to do
		return &parsedNode

	case *ast.GenDecl:
		parsedNode := data.GenericDeclaration{
			Node:           BuildNode(n),
			TokenPosition:  int(n.TokPos),
			Token:          int(n.Tok),
			Lparen:         int(n.Lparen),
			Specifications: []data.Specification{},
			Rparen:         int(n.Rparen),
		}
		if n.Doc != nil {
			parsedNode.Doc = parse(&parsedNode, n.Doc).(*data.CommentGroup)
		}
		for _, s := range n.Specs {
			parsedNode.Specifications = append(parsedNode.Specifications, parse(&parsedNode, s).(data.Specification))
		}
		return &parsedNode

	case *ast.FuncDecl:
		parsedNode := data.FunctionDeclaration{
			Node: BuildNode(n),
		}
		if n.Doc != nil {
			parsedNode.Doc = parse(&parsedNode, n.Doc).(*data.CommentGroup)
		}
		if n.Recv != nil {
			parsedNode.Receiver = parse(&parsedNode, n.Recv).(*data.FieldList)
		}
		parsedNode.Name = parse(&parsedNode, n.Name).(*data.Identifier)
		parsedNode.Type = parse(&parsedNode, n.Type).(*data.FunctionType)
		if n.Body != nil {
			parsedNode.Body = parse(&parsedNode, n.Body).(*data.BlockStatement)
		}
		return &parsedNode

	// Files and packages
	case *ast.File:
		parsedNode := data.File{
			Node:         BuildNode(n),
			Package:      int(n.Package),
			Declarations: []data.Declaration{},
			FileStart:    int(n.FileStart),
			FileEnd:      int(n.FileEnd),
			Imports:      []*data.ImportSpecification{},
			Unresolved:   []*data.Identifier{},
			Comments:     []*data.CommentGroup{},
			GoVersion:    n.GoVersion,
		}
		if n.Doc != nil {
			parsedNode.Doc = parse(&parsedNode, n.Doc).(*data.CommentGroup)
		}
		parsedNode.Name = parse(&parsedNode, n.Name).(*data.Identifier)
		for _, decl := range n.Decls {
			parsedNode.Declarations = append(parsedNode.Declarations, parse(&parsedNode, decl).(data.Declaration))
		}
		// don't walk n.Comments - they have been
		// visited already through the individual
		// nodes
		return &parsedNode

	case *ast.Package:
		parsedNode := data.Package{
			Node:    BuildNode(n),
			Name:    n.Name,
			Imports: map[string]data.Object{},
			Files:   map[string]*data.File{},
		}
		//for _, f := range n.Files {
		//	fileNode = parse()
		//	Walk(v, f)
		//}
		return &parsedNode

	default:
		panic(fmt.Sprintf("ast.Walk: unexpected node type %T", n))
	}

	return parsedNode
}
