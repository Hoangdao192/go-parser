package ast

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
	Node
	Doc          *CommentGroup          `json:"doc"`          // associated documentation; or nil
	Package      int                    `json:"package"`      // position of "package" keyword
	Name         *Identifier            `json:"name"`         // package name
	Declarations []Declaration          `json:"declarations"` // top-level declarations; or nil
	FileStart    int                    `json:"fileStart"`    // start of entire file
	FileEnd      int                    `json:"fileEnd"`      // end of entire file
	Scope        Scope                  `json:"scope"`        // package scope (this file only)
	Imports      []*ImportSpecification `json:"imports"`      // imports in this file
	Unresolved   []*Identifier          `json:"unresolved"`   // unresolved identifiers in this file
	Comments     []*CommentGroup        `json:"comments"`     // list of all comments in the source file
	GoVersion    string                 `json:"goVersion"`    // minimum Go version required by //go:build or // +build directives
}

// Start returns the position of the package declaration.
// (Use FileStart for the start of the entire file.)
func (f *File) Start() int {
	return f.Package
}

// End returns the end of the last declaration in the file.
// (Use FileEnd for the end of the entire file.)
func (f *File) End() int {
	if n := len(f.Declarations); n > 0 {
		return f.Declarations[n-1].End()
	}
	return f.Name.End()
}
