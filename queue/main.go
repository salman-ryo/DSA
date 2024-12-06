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

// A ring buffer is a fixed-size array that reuses space by wrapping around the indices. This avoids the need for resizing or memory inefficiencies.

type RingBuffQueue struct {
	elements []int
	head     int //head index  where deque takes place
	tail     int //tail index  where deque takes place
	size     int //current num of elements
	capacity int //max cap
}

// this func creates and returns a circular ring buff queu
func GenerateRingBuffQ(capacity int) *RingBuffQueue {
	return &RingBuffQueue{
		elements: make([]int, capacity),
		head:     0,
		tail:     0,
		size:     0,
		capacity: capacity,
	}
}

// Enqueu
func (rbq *RingBuffQueue) Enque(value int) bool {
	// if queue is full, return false
	if rbq.size == rbq.capacity {
		return false
	}
	// else add the element to the tail index
	rbq.elements[rbq.tail] = value
	// update tail index to next
	rbq.tail = (rbq.tail + 1) % rbq.capacity //% wraps the value so the index never goes more than 0-cap-1 as that would throw error, therfore makes this que circular
	rbq.size++
	return true
}

func (rbq *RingBuffQueue) Dequeue() (int, bool) {
	if rbq.size == 0 {
		return 0, false
	}
	headItem := rbq.elements[rbq.head]
	rbq.elements[rbq.head] = 0
	rbq.head = (rbq.head + 1) % rbq.capacity
	rbq.size--
	return headItem, true
}

func (rbq *RingBuffQueue) Print() {
	if rbq.size == 0 {
		fmt.Println("Queue Empty")
		return // Ensure no further execution
	}
	fmt.Print("head --> ")
	for i := 0; i < rbq.size; i++ { // Iterate over the valid size
		index := (rbq.head + i) % rbq.capacity
		fmt.Print(rbq.elements[index], " --> ")
	}
	fmt.Println("tail")
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
