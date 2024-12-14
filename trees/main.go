package main

import (
	"errors"
	"fmt"
)

type BTNode struct {
	Value int
	Left  *BTNode
	Right *BTNode
}

func (tn *BTNode) Insert(value int) error {
	if tn == nil {
		return errors.New("root node not initialized")
	}
	if value < tn.Value {
		if tn.Left == nil {
			tn.Left = &BTNode{Value: value}
		} else {
			return tn.Left.Insert(value)
		}
	} else {
		if tn.Right == nil {
			tn.Right = &BTNode{Value: value}
		} else {
			return tn.Right.Insert(value)
		}
	}
	return nil
}

// root,left,right
func (tn *BTNode) PreOrderTraversal(prefix string) {
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
func (tn *BTNode) InOrderTraversal(prefix string) {
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
func (tn *BTNode) PostOrderTraversal(prefix string) {
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

// Search if a value exists
func (btn *BTNode) Search(value int) bool {
	// if end of node, return false
	if btn == nil {
		return false
	}
	// if value matches current node, return true
	if value == btn.Value {
		return true
	}

	// if less, search in left node
	if value < btn.Value {
		return btn.Left.Search(value) //return the result from Search
	} else {
		return btn.Right.Search(value) //return the result from Search
	}
}

func main() {
	rooBTNode := &BTNode{Value: 3}
	rooBTNode.Insert(1)
	rooBTNode.Insert(2)
	rooBTNode.Insert(5)
	rooBTNode.Insert(4)

	// Display the tree in a readable structure
	fmt.Println("Tree structure:")
	rooBTNode.PreOrderTraversal("\t")

	exists := rooBTNode.Search(4)
	fmt.Println(exists)

}
