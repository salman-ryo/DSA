package main

import (
	"errors"
	"fmt"
)

// ARRAY
type AStack struct {
	elements []int
}

// define push method
func (s *AStack) Push(value int) {
	s.elements = append(s.elements, value)
}

// pop method
func (s *AStack) Pop() int {
	if len(s.elements) == 0 {
		return -1
	}
	latestItem := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return latestItem
}

// print method
func (s *AStack) Print() {
	if len(s.elements) == 0 {
		fmt.Print("AStack empty")
	}

	for i := 0; i < len(s.elements); i++ {
		fmt.Print(s.elements[i])

	}
}

//------------------------------------------------------------------------------------------------------------------------------------------------

// LINKED LIST

type Node struct {
	data int
	next *Node
}

// its just a linked list with pop and push method
type LStack struct {
	head *Node
}

func (ls *LStack) Push(value int) {
	newNode := &Node{data: value}
	newNode.next = ls.head
	ls.head = newNode
}

func (ls *LStack) Pop() (int, error) {
	if ls.head == nil {
		return 0, errors.New("stack is empty")
	} else {
		current := ls.head
		ls.head = current.next
		return current.data, nil
	}
}

func main() {
	myAStack := &AStack{elements: []int{1, 2}}
	myAStack.Push(3)
	myAStack.Push(4)
	fmt.Println("Before pop:")
	myAStack.Print()
	myAStack.Pop()
	fmt.Println("After pop:")
	myAStack.Print()
}
