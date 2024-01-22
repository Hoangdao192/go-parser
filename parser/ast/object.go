package ast

// An Object describes a named language entity such as a package,
// constant, type, variable, function (incl. methods), or label.
//
// The Data fields contains object-specific data:
//
//	Kind    Data type         Data value
//	Pkg     *Scope            package scope
//	Con     int               iota for the respective declaration
type Object struct {
	Kind int `json:"kind"`
	// declared name
	Name string `json:"name"`
	// corresponding Field, XxxSpec, FuncDecl, LabeledStmt, AssignStmt, Scope; or nil
	Decl any `json:"decl"`
	// object-specific data; or nil
	Data any `json:"data"`
	// placeholder for type information; may be nil
	Type any `json:"type"`
}
