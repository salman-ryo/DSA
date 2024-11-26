package main

import "fmt"

type Stack struct {
	elements []int
}

// define push method
func (s *Stack) Push(value int) {
	s.elements = append(s.elements, value)
}

// pop method
func (s *Stack) Pop() int {
	if len(s.elements) == 0 {
		return -1
	}
	latestItem := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return latestItem
}

// print method
func (s *Stack) Print() {
	if len(s.elements) == 0 {
		fmt.Print("Stack empty")
	}

	for i := 0; i < len(s.elements); i++ {
		fmt.Print(s.elements[i])

	}
}

func main() {
	myStack := &Stack{elements: []int{1, 2}}
	myStack.Push(3)
	myStack.Push(4)
	fmt.Println("Before pop:")
	myStack.Print()
	myStack.Pop()
	fmt.Println("After pop:")
	myStack.Print()
}
