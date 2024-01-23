package ast

// A Package node represents a set of source files
// collectively building a Go package.
type Package struct {
	Name    string             `json:"name"`    // package name
	Scope   *Scope             `json:"scope"`   // package scope across all files
	Imports map[string]*Object `json:"imports"` // map of package id -> package object
	Files   map[string]*File   `json:"files"`   // Go source files by filename
}

func (p *Package) Start() int {
	return 0
}
func (p *Package) End() int {
	return 0
}
