package main

import (
	"errors"
	"fmt"
)

type SQueue struct {
	elements []int
}

func (q *SQueue) Enque(value int) {
	q.elements = append(q.elements, value)
}

func (q *SQueue) Dequeue() (int, error) {
	if len(q.elements) == 0 {
		return 0, errors.New("queue is empty")
	}
	frontItem := q.elements[0]
	newSlice := make([]int, len(q.elements)-1)
	copy(newSlice, q.elements[1:])
	q.elements = newSlice
	return frontItem, nil
}

func (q *SQueue) Front() (int, error) {
	if len(q.elements) == 0 {
		return 0, errors.New("queue is empty")
	}
	return q.elements[0], nil
}

func (q *SQueue) Print() {
	if len(q.elements) == 0 {
		fmt.Println("Stack is empty")
	}
	for i := 0; i < len(q.elements); i++ {
		fmt.Print(q.elements[i], "<-")
	}
}

func main() {
	myQ := &SQueue{}
	myQ.Enque(1)
	myQ.Enque(2)
	myQ.Enque(3)
	myQ.Print()
	testDeq, err := myQ.Dequeue()
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println("We deqed: ", testDeq)
	myQ.Print()
}
