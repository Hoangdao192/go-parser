// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package ast declares the types used to represent syntax trees for Go
// packages.
package ast

import (
	"go/token"
	"strings"
)

// ----------------------------------------------------------------------------
// Interfaces
//
// There are 3 main classes of nodes: Expressions and type nodes,
// statement nodes, and declaration nodes. The node names usually
// match the corresponding Go spec production names to which they
// correspond. The node fields correspond to the individual parts
// of the respective productions.
//
// All nodes contain position information marking the beginning of
// the corresponding source text segment; it is accessible via the
// Pos accessor method. Nodes may contain additional position info
// for language constructs where comments may be found between parts
// of the construct (typically any larger, parenthesized subpart).
// That position information is needed to properly position comments
// when printing the construct.

// All node types implement the Node interface.
type Node interface {
	Position() int // position of first character belonging to the node
	End() int      // position of first character immediately after the node
}

// All expression nodes implement the Expr interface.
type Expr interface {
	Node
	exprNode()
}

// All statement nodes implement the Statement interface.
type Statement interface {
	Node
	StatementNode()
}

// All declaration nodes implement the Decl interface.
type Decl interface {
	Node
	declNode()
}

// ----------------------------------------------------------------------------
// Comments

// A Comment node represents a single //-style or /*-style comment.
//
// The Text field contains the comment text without carriage returns (\r) that
// may have been present in the source. Because a comment's end position is
// computed using len(Text), the position reported by End() does not match the
// true source end position for comments containing carriage returns.
type Comment struct {
	Slash int    // position of "/" starting the comment
	Text  string // comment text (excluding '\n' for //-style comments)
}

func (c *Comment) Position() int { return c.Slash }
func (c *Comment) End() int      { return int(int(c.Slash) + len(c.Text)) }

// A CommentGroup represents a sequence of comments
// with no other tokens and no empty lines between.
type CommentGroup struct {
	List []*Comment // len(List) > 0
}

func (g *CommentGroup) Position() int { return g.List[0].Position() }
func (g *CommentGroup) End() int      { return g.List[len(g.List)-1].End() }

func isWhitespace(ch byte) bool { return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r' }

func stripTrailingWhitespace(s string) string {
	i := len(s)
	for i > 0 && isWhitespace(s[i-1]) {
		i--
	}
	return s[0:i]
}

// Text returns the text of the comment.
// Comment markers (//, /*, and */), the first space of a line comment, and
// leading and trailing empty lines are removed.
// Comment directives like "//line" and "//go:noinline" are also removed.
// Multiple empty lines are reduced to one, and trailing space on lines is trimmed.
// Unless the result is empty, it is newline-terminated.
func (g *CommentGroup) Text() string {
	if g == nil {
		return ""
	}
	comments := make([]string, len(g.List))
	for i, c := range g.List {
		comments[i] = c.Text
	}

	lines := make([]string, 0, 10) // most comments are less than 10 lines
	for _, c := range comments {
		// Remove comment markers.
		// The parser has given us exactly the comment text.
		switch c[1] {
		case '/':
			//-style comment (no newline at the end)
			c = c[2:]
			if len(c) == 0 {
				// empty line
				break
			}
			if c[0] == ' ' {
				// strip first space - required for Example tests
				c = c[1:]
				break
			}
			if isDirective(c) {
				// Ignore //go:noinline, //line, and so on.
				continue
			}
		case '*':
			/*-style comment */
			c = c[2 : len(c)-2]
		}

		// Split on newlines.
		cl := strings.Split(c, "\n")

		// Walk lines, stripping trailing white space and adding to list.
		for _, l := range cl {
			lines = append(lines, stripTrailingWhitespace(l))
		}
	}

	// Remove leading blank lines; convert runs of
	// interior blank lines to a single blank line.
	n := 0
	for _, line := range lines {
		if line != "" || n > 0 && lines[n-1] != "" {
			lines[n] = line
			n++
		}
	}
	lines = lines[0:n]

	// Add final "" entry to get trailing newline from Join.
	if n > 0 && lines[n-1] != "" {
		lines = append(lines, "")
	}

	return strings.Join(lines, "\n")
}

// isDirective reports whether c is a comment directive.
// This code is also in go/printer.
func isDirective(c string) bool {
	// "//line " is a line directive.
	// "//extern " is for gccgo.
	// "//export " is for cgo.
	// (The // has been removed.)
	if strings.HasPrefix(c, "line ") || strings.HasPrefix(c, "extern ") || strings.HasPrefix(c, "export ") {
		return true
	}

	// "//[a-z0-9]+:[a-z0-9]"
	// (The // has been removed.)
	colon := strings.Index(c, ":")
	if colon <= 0 || colon+1 >= len(c) {
		return false
	}
	for i := 0; i <= colon+1; i++ {
		if i == colon {
			continue
		}
		b := c[i]
		if !('a' <= b && b <= 'z' || '0' <= b && b <= '9') {
			return false
		}
	}
	return true
}

// ----------------------------------------------------------------------------
// Expressions and types

// A Field represents a Field declaration list in a struct type,
// a method list in an interface type, or a parameter/result declaration
// in a signature.
// Field.Names is nil for unnamed parameters (parameter lists which only contain types)
// and embedded struct fields. In the latter case, the field name is the type name.
type Field struct {
	Doc     *CommentGroup // associated documentation; or nil
	Names   []*Ident      // field/method/(type) parameter names; or nil
	Type    Expr          // field/method/parameter type; or nil
	Tag     *BasicLit     // field tag; or nil
	Comment *CommentGroup // line comments; or nil
}

func (f *Field) Position() int {
	if len(f.Names) > 0 {
		return f.Names[0].Position()
	}
	if f.Type != nil {
		return f.Type.Position()
	}
	return token.NoPos
}

func (f *Field) End() int {
	if f.Tag != nil {
		return f.Tag.End()
	}
	if f.Type != nil {
		return f.Type.End()
	}
	if len(f.Names) > 0 {
		return f.Names[len(f.Names)-1].End()
	}
	return token.NoPos
}

// A FieldList represents a list of Fields, enclosed by parentheses,
// curly braces, or square brackets.
type FieldList struct {
	Opening int      // position of opening parenthesis/brace/bracket, if any
	List    []*Field // field list; or nil
	Closing int      // position of closing parenthesis/brace/bracket, if any
}

func (f *FieldList) Position() int {
	if f.Opening.IsValid() {
		return f.Opening
	}
	// the list should not be empty in this case;
	// be conservative and guard against bad ASTs
	if len(f.List) > 0 {
		return f.List[0].Position()
	}
	return token.NoPos
}

func (f *FieldList) End() int {
	if f.Closing.IsValid() {
		return f.Closing + 1
	}
	// the list should not be empty in this case;
	// be conservative and guard against bad ASTs
	if n := len(f.List); n > 0 {
		return f.List[n-1].End()
	}
	return token.NoPos
}

// NumFields returns the number of parameters or struct fields represented by a FieldList.
func (f *FieldList) NumFields() int {
	n := 0
	if f != nil {
		for _, g := range f.List {
			m := len(g.Names)
			if m == 0 {
				m = 1
			}
			n += m
		}
	}
	return n
}

// An expression is represented by a tree consisting of one
// or more of the following concrete expression nodes.
type (
	// A BadExpr node is a placeholder for an expression containing
	// syntax errors for which a correct expression node cannot be
	// created.
	//
	BadExpr struct {
		From, To int // position range of bad expression
	}

	// An Ellipsis node stands for the "..." type in a
	// parameter list or the "..." length in an array type.
	//
	Ellipsis struct {
		Ellipsis int  // position of "..."
		Elt      Expr // ellipsis element type (parameter lists only); or nil
	}

	// A BasicLit node represents a literal of basic type.
	BasicLit struct {
		ValuePos int         // literal position
		Kind     token.Token // token.INT, token.FLOAT, token.IMAG, token.CHAR, or token.STRING
		Value    string      // literal string; e.g. 42, 0x7f, 3.14, 1e-9, 2.4i, 'a', '\x7f', "foo" or `\m\n\o`
	}

	// A FuncLit node represents a function literal.
	FuncLit struct {
		Type *FuncType       // function type
		Body *BlockStatement // function body
	}

	// A CompositeLit node represents a composite literal.
	CompositeLit struct {
		Type       Expr   // literal type; or nil
		Lbrace     int    // position of "{"
		Elts       []Expr // list of composite elements; or nil
		Rbrace     int    // position of "}"
		Incomplete bool   // true if (source) expressions are missing in the Elts list
	}

	// A ParenExpr node represents a parenthesized expression.
	ParenExpr struct {
		Lparen int  // position of "("
		X      Expr // parenthesized expression
		Rparen int  // position of ")"
	}

	// A SelectorExpr node represents an expression followed by a selector.
	SelectorExpr struct {
		X   Expr   // expression
		Sel *Ident // field selector
	}

	// An IndexExpr node represents an expression followed by an index.
	IndexExpr struct {
		X      Expr // expression
		Lbrack int  // position of "["
		Index  Expr // index expression
		Rbrack int  // position of "]"
	}

	// An IndexListExpr node represents an expression followed by multiple
	// indices.
	IndexListExpr struct {
		X       Expr   // expression
		Lbrack  int    // position of "["
		Indices []Expr // index expressions
		Rbrack  int    // position of "]"
	}

	// A SliceExpr node represents an expression followed by slice indices.
	SliceExpr struct {
		X      Expr // expression
		Lbrack int  // position of "["
		Low    Expr // begin of slice range; or nil
		High   Expr // end of slice range; or nil
		Max    Expr // maximum capacity of slice; or nil
		Slice3 bool // true if 3-index slice (2 colons present)
		Rbrack int  // position of "]"
	}

	// A TypeAssertExpr node represents an expression followed by a
	// type assertion.
	//
	TypeAssertExpr struct {
		X      Expr // expression
		Lparen int  // position of "("
		Type   Expr // asserted type; nil means type switch X.(type)
		Rparen int  // position of ")"
	}

	// A CallExpr node represents an expression followed by an argument list.
	CallExpr struct {
		Fun      Expr   // function expression
		Lparen   int    // position of "("
		Args     []Expr // function arguments; or nil
		Ellipsis int    // position of "..." (token.NoPos if there is no "...")
		Rparen   int    // position of ")"
	}

	// A StarExpr node represents an expression of the form "*" Expression.
	// Semantically it could be a unary "*" expression, or a pointer type.
	//
	StarExpr struct {
		Star int  // position of "*"
		X    Expr // operand
	}

	// A UnaryExpr node represents a unary expression.
	// Unary "*" expressions are represented via StarExpr nodes.
	//
	UnaryExpr struct {
		OpPos int         // position of Op
		Op    token.Token // operator
		X     Expr        // operand
	}

	// A BinaryExpr node represents a binary expression.
	BinaryExpr struct {
		X     Expr        // left operand
		OpPos int         // position of Op
		Op    token.Token // operator
		Y     Expr        // right operand
	}

	// A KeyValueExpr node represents (key : value) pairs
	// in composite literals.
	//
	KeyValueExpr struct {
		Key   Expr
		Colon int // position of ":"
		Value Expr
	}
)

// The direction of a channel type is indicated by a bit
// mask including one or both of the following constants.
type ChanDir int

const (
	SEND ChanDir = 1 << iota
	RECV
)

// A type is represented by a tree consisting of one
// or more of the following type-specific expression
// nodes.
type (
	// An ArrayType node represents an array or slice type.
	ArrayType struct {
		Lbrack int  // position of "["
		Len    Expr // Ellipsis node for [...]T array types, nil for slice types
		Elt    Expr // element type
	}

	// A StructType node represents a struct type.
	StructType struct {
		Struct     int        // position of "struct" keyword
		Fields     *FieldList // list of field declarations
		Incomplete bool       // true if (source) fields are missing in the Fields list
	}

	// Pointer types are represented via StarExpr nodes.

	// A FuncType node represents a function type.
	FuncType struct {
		Func       int        // position of "func" keyword (token.NoPos if there is no "func")
		TypeParams *FieldList // type parameters; or nil
		Params     *FieldList // (incoming) parameters; non-nil
		Results    *FieldList // (outgoing) results; or nil
	}

	// An InterfaceType node represents an interface type.
	InterfaceType struct {
		Interface  int        // position of "interface" keyword
		Methods    *FieldList // list of embedded interfaces, methods, or types
		Incomplete bool       // true if (source) methods or types are missing in the Methods list
	}

	// A MapType node represents a map type.
	MapType struct {
		Map   int // position of "map" keyword
		Key   Expr
		Value Expr
	}

	// A ChanType node represents a channel type.
	ChanType struct {
		Begin int     // position of "chan" keyword or "<-" (whichever comes first)
		Arrow int     // position of "<-" (token.NoPos if there is no "<-")
		Dir   ChanDir // channel direction
		Value Expr    // value type
	}
)

// Pos and End implementations for expression/type nodes.

func (x *BadExpr) Position() int  { return x.From }
func (x *Ident) Position() int    { return x.NamePos }
func (x *Ellipsis) Position() int { return x.Ellipsis }
func (x *BasicLit) Position() int { return x.ValuePos }
func (x *FuncLit) Position() int  { return x.Type.Position() }
func (x *CompositeLit) Position() int {
	if x.Type != nil {
		return x.Type.Position()
	}
	return x.Lbrace
}
func (x *ParenExpr) Position() int      { return x.Lparen }
func (x *SelectorExpr) Position() int   { return x.X.Position() }
func (x *IndexExpr) Position() int      { return x.X.Position() }
func (x *IndexListExpr) Position() int  { return x.X.Position() }
func (x *SliceExpr) Position() int      { return x.X.Position() }
func (x *TypeAssertExpr) Position() int { return x.X.Position() }
func (x *CallExpr) Position() int       { return x.Fun.Position() }
func (x *StarExpr) Position() int       { return x.Star }
func (x *UnaryExpr) Position() int      { return x.OpPos }
func (x *BinaryExpr) Position() int     { return x.X.Position() }
func (x *KeyValueExpr) Position() int   { return x.Key.Position() }
func (x *ArrayType) Position() int      { return x.Lbrack }
func (x *StructType) Position() int     { return x.Struct }
func (x *FuncType) Position() int {
	if x.Func.IsValid() || x.Params == nil { // see issue 3870
		return x.Func
	}
	return x.Params.Position() // interface method declarations have no "func" keyword
}
func (x *InterfaceType) Position() int { return x.Interface }
func (x *MapType) Position() int       { return x.Map }
func (x *ChanType) Position() int      { return x.Begin }

func (x *BadExpr) End() int { return x.To }
func (x *Ident) End() int   { return int(int(x.NamePos) + len(x.Name)) }
func (x *Ellipsis) End() int {
	if x.Elt != nil {
		return x.Elt.End()
	}
	return x.Ellipsis + 3 // len("...")
}
func (x *BasicLit) End() int       { return int(int(x.ValuePos) + len(x.Value)) }
func (x *FuncLit) End() int        { return x.Body.End() }
func (x *CompositeLit) End() int   { return x.Rbrace + 1 }
func (x *ParenExpr) End() int      { return x.Rparen + 1 }
func (x *SelectorExpr) End() int   { return x.Sel.End() }
func (x *IndexExpr) End() int      { return x.Rbrack + 1 }
func (x *IndexListExpr) End() int  { return x.Rbrack + 1 }
func (x *SliceExpr) End() int      { return x.Rbrack + 1 }
func (x *TypeAssertExpr) End() int { return x.Rparen + 1 }
func (x *CallExpr) End() int       { return x.Rparen + 1 }
func (x *StarExpr) End() int       { return x.X.End() }
func (x *UnaryExpr) End() int      { return x.X.End() }
func (x *BinaryExpr) End() int     { return x.Y.End() }
func (x *KeyValueExpr) End() int   { return x.Value.End() }
func (x *ArrayType) End() int      { return x.Elt.End() }
func (x *StructType) End() int     { return x.Fields.End() }
func (x *FuncType) End() int {
	if x.Results != nil {
		return x.Results.End()
	}
	return x.Params.End()
}
func (x *InterfaceType) End() int { return x.Methods.End() }
func (x *MapType) End() int       { return x.Value.End() }
func (x *ChanType) End() int      { return x.Value.End() }

// exprNode() ensures that only expression/type nodes can be
// assigned to an Expr.
func (*BadExpr) exprNode()        {}
func (*Ident) exprNode()          {}
func (*Ellipsis) exprNode()       {}
func (*BasicLit) exprNode()       {}
func (*FuncLit) exprNode()        {}
func (*CompositeLit) exprNode()   {}
func (*ParenExpr) exprNode()      {}
func (*SelectorExpr) exprNode()   {}
func (*IndexExpr) exprNode()      {}
func (*IndexListExpr) exprNode()  {}
func (*SliceExpr) exprNode()      {}
func (*TypeAssertExpr) exprNode() {}
func (*CallExpr) exprNode()       {}
func (*StarExpr) exprNode()       {}
func (*UnaryExpr) exprNode()      {}
func (*BinaryExpr) exprNode()     {}
func (*KeyValueExpr) exprNode()   {}

func (*ArrayType) exprNode()     {}
func (*StructType) exprNode()    {}
func (*FuncType) exprNode()      {}
func (*InterfaceType) exprNode() {}
func (*MapType) exprNode()       {}
func (*ChanType) exprNode()      {}

// ----------------------------------------------------------------------------
// Convenience functions for Idents

// NewIdent creates a new Ident without position.
// Useful for ASTs generated by code other than the Go parser.
func NewIdent(name string) *Ident { return &Ident{token.NoPos, name, nil} }

// IsExported reports whether name starts with an upper-case letter.
func IsExported(name string) bool { return token.IsExported(name) }

// IsExported reports whether id starts with an upper-case letter.
func (id *Ident) IsExported() bool { return token.IsExported(id.Name) }

func (id *Ident) String() string {
	if id != nil {
		return id.Name
	}
	return "<nil>"
}

// Pos and End implementations for statement nodes.

// StatementNode() ensures that only statement nodes can be
// assigned to a Statement.

// ----------------------------------------------------------------------------
// Declarations

// A Spec node represents a single (non-parenthesized) import,
// constant, type, or variable declaration.
type (
	// The Spec type stands for any of *ImportSpec, *ValueSpec, and *TypeSpec.
	Spec interface {
		Node
		specNode()
	}

	// An ImportSpec node represents a single package import.
	ImportSpec struct {
		Doc     *CommentGroup // associated documentation; or nil
		Name    *Ident        // local package name (including "."); or nil
		Path    *BasicLit     // import path
		Comment *CommentGroup // line comments; or nil
		EndPos  int           // end of spec (overrides Path.Pos if nonzero)
	}

	// A ValueSpec node represents a constant or variable declaration
	// (ConstSpec or VarSpec production).
	//
	ValueSpec struct {
		Doc     *CommentGroup // associated documentation; or nil
		Names   []*Ident      // value names (len(Names) > 0)
		Type    Expr          // value type; or nil
		Values  []Expr        // initial values; or nil
		Comment *CommentGroup // line comments; or nil
	}

	// A TypeSpec node represents a type declaration (TypeSpec production).
	TypeSpec struct {
		Doc        *CommentGroup // associated documentation; or nil
		Name       *Ident        // type name
		TypeParams *FieldList    // type parameters; or nil
		Assign     int           // position of '=', if any
		Type       Expr          // *Ident, *ParenExpr, *SelectorExpr, *StarExpr, or any of the *XxxTypes
		Comment    *CommentGroup // line comments; or nil
	}
)

// Pos and End implementations for spec nodes.

func (s *ImportSpec) Position() int {
	if s.Name != nil {
		return s.Name.Position()
	}
	return s.Path.Position()
}
func (s *ValueSpec) Position() int { return s.Names[0].Position() }
func (s *TypeSpec) Position() int  { return s.Name.Position() }

func (s *ImportSpec) End() int {
	if s.EndPos != 0 {
		return s.EndPos
	}
	return s.Path.End()
}

func (s *ValueSpec) End() int {
	if n := len(s.Values); n > 0 {
		return s.Values[n-1].End()
	}
	if s.Type != nil {
		return s.Type.End()
	}
	return s.Names[len(s.Names)-1].End()
}
func (s *TypeSpec) End() int { return s.Type.End() }

// specNode() ensures that only spec nodes can be
// assigned to a Spec.
func (*ImportSpec) specNode() {}
func (*ValueSpec) specNode()  {}
func (*TypeSpec) specNode()   {}

// A declaration is represented by one of the following declaration nodes.
type (
	// A BadDecl node is a placeholder for a declaration containing
	// syntax errors for which a correct declaration node cannot be
	// created.
	//
	BadDecl struct {
		From, To int // position range of bad declaration
	}

	// A GenDecl node (generic declaration node) represents an import,
	// constant, type or variable declaration. A valid Lparen position
	// (Lparen.IsValid()) indicates a parenthesized declaration.
	//
	// Relationship between Tok value and Specs element type:
	//
	//	token.IMPORT  *ImportSpec
	//	token.CONST   *ValueSpec
	//	token.TYPE    *TypeSpec
	//	token.VAR     *ValueSpec
	//
	GenDecl struct {
		Doc    *CommentGroup // associated documentation; or nil
		TokPos int           // position of Tok
		Tok    token.Token   // IMPORT, CONST, TYPE, or VAR
		Lparen int           // position of '(', if any
		Specs  []Spec
		Rparen int // position of ')', if any
	}

	// A FuncDecl node represents a function declaration.
	FuncDecl struct {
		Doc  *CommentGroup   // associated documentation; or nil
		Recv *FieldList      // receiver (methods); or nil (functions)
		Name *Ident          // function/method name
		Type *FuncType       // function signature: type and value parameters, results, and position of "func" keyword
		Body *BlockStatement // function body; or nil for external (non-Go) function
	}
)

// Pos and End implementations for declaration nodes.

func (d *BadDecl) Position() int  { return d.From }
func (d *GenDecl) Position() int  { return d.TokPos }
func (d *FuncDecl) Position() int { return d.Type.Position() }

func (d *BadDecl) End() int { return d.To }
func (d *GenDecl) End() int {
	if d.Rparen.IsValid() {
		return d.Rparen + 1
	}
	return d.Specs[0].End()
}
func (d *FuncDecl) End() int {
	if d.Body != nil {
		return d.Body.End()
	}
	return d.Type.End()
}

// declNode() ensures that only declaration nodes can be
// assigned to a Decl.
func (*BadDecl) declNode()  {}
func (*GenDecl) declNode()  {}
func (*FuncDecl) declNode() {}

// ----------------------------------------------------------------------------
// Files and packages

// A File node represents a Go source file.
//
// The Comments list contains all comments in the source file in order of
// appearance, including the comments that are pointed to from other nodes
// via Doc and Comment fields.
//
// For correct printing of source code containing comments (using packages
// go/format and go/printer), special care must be taken to update comments
// when a File's syntax tree is modified: For printing, comments are interspersed
// between tokens based on their position. If syntax tree nodes are
// removed or moved, relevant comments in their vicinity must also be removed
// (from the File.Comments list) or moved accordingly (by updating their
// positions). A CommentMap may be used to facilitate some of these operations.
//
// Whether and how a comment is associated with a node depends on the
// interpretation of the syntax tree by the manipulating program: Except for Doc
// and Comment comments directly associated with nodes, the remaining comments
// are "free-floating" (see also issues #18593, #20744).
type File struct {
	Doc     *CommentGroup // associated documentation; or nil
	Package int           // position of "package" keyword
	Name    *Ident        // package name
	Decls   []Decl        // top-level declarations; or nil

	FileStart, FileEnd int             // start and end of entire file
	Scope              *Scope          // package scope (this file only)
	Imports            []*ImportSpec   // imports in this file
	Unresolved         []*Ident        // unresolved identifiers in this file
	Comments           []*CommentGroup // list of all comments in the source file
	GoVersion          string          // minimum Go version required by //go:build or // +build directives
}

// Pos returns the position of the package declaration.
// (Use FileStart for the start of the entire file.)
func (f *File) Position() int { return f.Package }

// End returns the end of the last declaration in the file.
// (Use FileEnd for the end of the entire file.)
func (f *File) End() int {
	if n := len(f.Decls); n > 0 {
		return f.Decls[n-1].End()
	}
	return f.Name.End()
}

// A Package node represents a set of source files
// collectively building a Go package.
type Package struct {
	Name    string             // package name
	Scope   *Scope             // package scope across all files
	Imports map[string]*Object // map of package id -> package object
	Files   map[string]*File   // Go source files by filename
}

func (p *Package) Position() int { return token.NoPos }
func (p *Package) End() int      { return token.NoPos }

// IsGenerated reports whether the file was generated by a program,
// not handwritten, by detecting the special comment described
// at https://go.dev/s/generatedcode.
//
// The syntax tree must have been parsed with the ParseComments flag.
// Example:
//
//	f, err := parser.ParseFile(fset, filename, src, parser.ParseComments|parser.PackageClauseOnly)
//	if err != nil { ... }
//	gen := ast.IsGenerated(f)
func IsGenerated(file *File) bool {
	_, ok := generator(file)
	return ok
}

func generator(file *File) (string, bool) {
	for _, group := range file.Comments {
		for _, comment := range group.List {
			if comment.Position() > file.Package {
				break // after package declaration
			}
			// opt: check Contains first to avoid unnecessary array allocation in Split.
			const prefix = "// Code generated "
			if strings.Contains(comment.Text, prefix) {
				for _, line := range strings.Split(comment.Text, "\n") {
					if rest, ok := strings.CutPrefix(line, prefix); ok {
						if gen, ok := strings.CutSuffix(rest, " DO NOT EDIT."); ok {
							return gen, true
						}
					}
				}
			}
		}
	}
	return "", false
}
