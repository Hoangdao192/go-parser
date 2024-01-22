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
	Kind int    `json:"kind"`
	Name string `json:"name"` // declared name
	// corresponding Field, XxxSpecification, FunctionDeclaration, LabeledStatement, AssignStatement, Scope; or nil
	Declaration any `json:"declaration"`
	Data        any `json:"data"` // object-specific data; or nil
	Type        any `json:"type"` // placeholder for type information; may be nil
}
