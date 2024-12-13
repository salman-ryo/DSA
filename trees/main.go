package main

import (
	"errors"
	"fmt"
)

type TNode struct {
	Value int
	Left  *TNode
	Right *TNode
}

func (tn *TNode) Insert(value int) error {
	if tn == nil {
		return errors.New("root node not initialized")
	}
	if value < tn.Value {
		if tn.Left == nil {
			tn.Left = &TNode{Value: value}
		} else {
			return tn.Left.Insert(value)
		}
	} else {
		if tn.Right == nil {
			tn.Right = &TNode{Value: value}
		} else {
			return tn.Right.Insert(value)
		}
	}
	return nil
}

// root,left,right
func (tn *TNode) PreOrderTraversal(prefix string) {
	if tn == nil {
		return
	}
	// Print the current node
	fmt.Println(prefix + fmt.Sprintf("%d", tn.Value))
	// Traverse the left subtree with additional prefix
	if tn.Left != nil {
		tn.Left.PreOrderTraversal(prefix + "  ")
	} else {
		fmt.Println(prefix + "  /") // Indicate no left child
	}
	// Traverse the right subtree with additional prefix
	if tn.Right != nil {
		tn.Right.PreOrderTraversal(prefix + "  ")
	} else {
		fmt.Println(prefix + "  \\") // Indicate no right child
	}
}

// left,root,right
func (tn *TNode) InOrderTraversal(prefix string) {
	if tn == nil {
		return
	}

	// Traverse the left subtree with additional prefix
	if tn.Left != nil {
		tn.Left.InOrderTraversal(prefix + "  ")
	} else {
		fmt.Println(prefix + "  /") // Indicate no left child
	}

	// Print the current node
	fmt.Println(prefix + fmt.Sprintf("%d", tn.Value))

	// Traverse the right subtree with additional prefix
	if tn.Right != nil {
		tn.Right.InOrderTraversal(prefix + "  ")
	} else {
		fmt.Println(prefix + "  \\") // Indicate no right child
	}
}

// left,right,root
func (tn *TNode) PostOrderTraversal(prefix string) {
	if tn == nil {
		return
	}

	// Traverse the left subtree with additional prefix
	if tn.Left != nil {
		tn.Left.PostOrderTraversal(prefix + "  ")
	} else {
		fmt.Println(prefix + "  /") // Indicate no left child
	}

	// Traverse the right subtree with additional prefix
	if tn.Right != nil {
		tn.Right.PostOrderTraversal(prefix + "  ")
	} else {
		fmt.Println(prefix + "  \\") // Indicate no right child
	}

	// Print the current node
	fmt.Println(prefix + fmt.Sprintf("%d", tn.Value))
}

func main() {
	rootNode := &TNode{Value: 2}
	rootNode.Insert(1)
	rootNode.Insert(1)
	rootNode.Insert(0)
	rootNode.Insert(3)
	rootNode.Insert(3)
	rootNode.Insert(4)

	// Display the tree in a readable structure
	fmt.Println("Tree structure:")
	rootNode.PostOrderTraversal("\t")
}
