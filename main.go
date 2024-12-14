package main

import "fmt"

type TNode struct {
	Value int
	Left  *TNode
	Right *TNode
}

func (tn *TNode) Insert(value int) {
	if tn == nil {
		return
	}
	if value < tn.Value {
		if tn.Left == nil {
			tn.Left = &TNode{Value: value}
		} else {
			tn.Left.Insert(value)
		}
	} else {
		if tn.Right == nil {
			tn.Right = &TNode{Value: value}
		} else {
			tn.Right.Insert(value)
		}
	}
}

//root, left, right
func (tn *TNode) PreOrderTraversal() {
	if tn == nil {
		return
	}
	fmt.Println(tn.Value, " ")
	tn.Left.PreOrderTraversal()
	tn.Right.PreOrderTraversal()
}

//left,root, right
func (tn *TNode) InOrderTraversal() {
	if tn == nil {
		return
	}
	tn.Left.InOrderTraversal()
	fmt.Println(tn.Value, " ")
	tn.Right.InOrderTraversal()
}

//left, right, root
func (tn *TNode) PostOrderTraversal() {
	if tn == nil {
		return
	}
	tn.Left.PostOrderTraversal()
	tn.Right.PostOrderTraversal()
	fmt.Println(tn.Value, " ")
}

func main() {
	rootNode := &TNode{Value: 3}
	rootNode.Insert(1)
	rootNode.Insert(2)
	rootNode.Insert(4)
	rootNode.Insert(5)
	rootNode.PreOrderTraversal()
	fmt.Println("===========================")
	rootNode.InOrderTraversal()
	fmt.Println("===========================")
	rootNode.PostOrderTraversal()
}
