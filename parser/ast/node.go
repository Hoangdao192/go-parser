package ast

type INode interface {
	Start() int // position of first character belonging to the node
	End() int   // position of first character immediately after the node
	//MarshalJson() ([]byte, error)
}

type Node struct {
	Children      []INode `json:"children"`
	StartPosition int     `json:"start"`
	EndPosition   int     `json:"end"`
	Code          string  `json:"code"`
	NodeType      string  `json:"nodeType"`
	Column        int64   `json:"column"`
	Line          int64   `json:"line"`
}

func (n Node) AddChild(child INode) {
	n.Children = append(n.Children, child)
}

func (n Node) Start() int {
	return n.StartPosition
}

func (n Node) End() int {
	return n.EndPosition
}
