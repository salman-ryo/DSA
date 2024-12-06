package main

import "errors"

type QNode struct {
	value int
	next  *QNode
}

type LLQueue struct {
	head *QNode
	tail *QNode
}

func (llq *LLQueue) Enqueu(value int) {
	newQNode := &QNode{value: value}
	// if empty
	if llq.tail == nil {
		llq.head = newQNode
		llq.tail = newQNode
	} else {
		// add value to tail
		llq.tail.next = newQNode
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
	// if q becomes emtpy, set tail to nil too so it doesnt point to old QNode
	if llq.head == nil {
		llq.tail = nil
	}
	return headValue, nil
}
