package ast

// A Scope maintains the set of named language entities declared
// in the scope and a link to the immediately surrounding (outer)
// scope.
type Scope struct {
	Outer   *Scope             `json:"outer"`
	Objects map[string]*Object `json:"objects"`
}
