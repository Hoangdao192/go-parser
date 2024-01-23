package main

import "fmt"

type INode interface {
	AddChild(INode)
}

type Node struct {
	Children []*Node
}

func (n Node) AddChild(child INode) {
	n.Children = append(n.Children, child.(*Node))
}

type Expression struct {
}

func main() {
	var child INode = &Node{}

	var n Node = Node{}
	n.Children = append(n.Children, child.(*Node))
	fmt.Println(b.Name)

}
