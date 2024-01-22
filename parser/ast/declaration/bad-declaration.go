package declaration

// A BadDeclaration node is a placeholder for a declaration containing
// syntax errors for which a correct declaration node cannot be
// created.
type BadDeclaration struct {
	From int `json:"from"` // position range of bad declaration
	To   int `json:"to"`
}

func (d *BadDeclaration) Position() int {
	return d.From
}

func (d *BadDeclaration) End() int {
	return d.To
}

func (*BadDeclaration) declarationNode() {}
