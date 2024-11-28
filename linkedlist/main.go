package main

import (
	"errors"
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

// add to head
func (ll *LinkedList) Add(value int) {
	newNode := &Node{data: value}
	newNode.next = ll.head
	ll.head = newNode
}

// add to end
func (ll *LinkedList) Append(value int) {
	newNode := &Node{data: value}
	current := ll.head
	// if current is null, ll is empty
	if current == nil {
		ll.head = newNode
	}
	// traverse till last node
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

// remove head
func (ll *LinkedList) Remove() error {
	if ll.head == nil {
		return fmt.Errorf("LL Already mt")
	}
	newHead := ll.head.next
	ll.head = newHead
	return nil
}

// remove last
func (ll *LinkedList) Pop() (int, error) {
	if ll.head == nil {
		return 0, errors.New("LL Already mt")
	}
	// if only element, remove it
	if ll.head.next == nil {
		lastItem := ll.head
		ll.head = nil
		return lastItem.data, nil
	}

	// traverse to last sec element
	current := ll.head
	for current.next.next != nil {
		current = current.next
	}
	// cut off last ele
	finalNode := current.next
	current.next = nil
	return finalNode.data, nil
}

// remove between
func (ll *LinkedList) Delete(value int) error {
	if ll.head == nil {
		return fmt.Errorf("LL Already mt")
	}

	current := ll.head
	if current.data == value {
		ll.head = current.next
		return nil
	}

	// traverse until next node has the item or it's the last one
	for current.next != nil && current.next.data != value {
		current = current.next
	}
	// skip the next node
	current.next = current.next.next
	return nil
}

// print ll
func (ll *LinkedList) Print() {
	if ll.head == nil {
		fmt.Println("LL mt")
	}
	current := ll.head
	for current != nil {
		fmt.Print(current.data, " --> ")
		current = current.next
	}
	fmt.Print("nil")

}

func main() {
	myLL := &LinkedList{}
	myLL.Add(4)
	myLL.Add(5)
	myLL.Add(6)
	myLL.Add(7)
	fmt.Println("Ll: ")
	// myLL.Print()
	// myLL.Delete(4)
	testPop, err := myLL.Pop()
	if err != nil {
		fmt.Println(err)
	}
	myLL.Print()
	fmt.Println("Popped item: ", testPop)

}
