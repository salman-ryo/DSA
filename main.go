package main

import "errors"

type SQueue struct {
	elements []int
}

func (sq *SQueue) Enqueue(value int) {
	sq.elements = append(sq.elements, value)
}

func (sq *SQueue) Dequeue() (int, error) {
	if len(sq.elements) == 0 {
		return 0, errors.New("queue is empty")
	}
	headValue := sq.elements[0]
	sq.elements = sq.elements[1:]
	return headValue, nil
}

func (sq *SQueue) Peek() (int, error) {
	if len(sq.elements) == 0 {
		return 0, errors.New("queue is empty")
	}
	return sq.elements[0], nil
}

type Node struct {
	value int
	next  *Node
}

type LLQueue struct {
	head *Node
	tail *Node
}

func (llq *LLQueue) Enqueu(value int) {
	newNode := &Node{value: value}
	// if empty
	if llq.tail == nil {
		llq.head = newNode
		llq.tail = newNode
	} else {
		// add value to tail
		llq.tail.next = newNode
		// update tail to next value
		llq.tail = llq.tail.next
	}
}

func (llq *LLQueue) Dequeue() (int, error) {
	if llq.head == nil {
		return 0, errors.New("list queue is emtpy")
	}

	headValue := llq.head.value
	llq.head = llq.head.next
	// if q becomes emtpy, set tail to nil too so it doesnt point to old node
	if llq.head == nil {
		llq.tail = nil
	}
	return headValue, nil
}
