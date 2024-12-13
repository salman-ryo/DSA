package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type TNode struct {
	Value int    `json:"value"` // JSON field name
	Left  *TNode `json:"left"`  // Recursive serialization
	Right *TNode `json:"right"`
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

func main() {
	rootNode := &TNode{Value: 2}
	rootNode.Insert(1)
	rootNode.Insert(3)

	// Convert tree to JSON
	jsonRoot, err := json.MarshalIndent(rootNode, "", "  ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(jsonRoot))
}
