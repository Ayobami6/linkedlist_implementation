package linkedlist

import (
	"errors"
	"fmt"
)

// create the Node struct

type Node struct {
	val  int
	next *Node
}

func NewNode(val int) *Node {
	return &Node{
		val: val,
	}
}

func (n *Node) SetNext(node *Node) {
	n.next = node
}

func (n *Node) GetNext() *Node {
	return n.next
}

func (n *Node) GetVal() int {
	return n.val
}

// create the linkedlist struct

type LinkedList struct {
	head  *Node
	count int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		head:  nil,
		count: 0,
	}
}

// insert at the start of the linkedlist

func (l *LinkedList) InsertAtTheStart(value int) {
	// insert
	newNode := NewNode(value)
	if l.head != nil {
		newNode.SetNext(l.head)
	}
	l.head = newNode
	l.count++
}

func (l *LinkedList) LookUp(value int) int {
	node := l.head
	var val int
	for node != nil {
		if node.GetVal() == value {
			val = node.GetVal()
		}
		node = node.GetNext()
	}
	return val
}

func (l *LinkedList) InsertAtTheEnd(value int) {
	node := l.head
	newNode := NewNode(value)
	count := 1
	// traverse to the end node
	for node != nil {
		if count == l.count {
			break
		}
		node = node.GetNext()
		count++
	}
	node.next = newNode
	l.count++

}

func (l *LinkedList) GetCount() int {
	return l.count
}

func (l *LinkedList) DumpList() {
	node := l.head
	for node != nil {
		fmt.Printf("Node val is %d\n", node.val)
		node = node.GetNext()
	}
}

// delete the head of a linkedlist
func (l *LinkedList) DeleteHead() error {
	current := l.head
	if current == nil {
		// if head is empty
		return errors.New("linkedlist is empty")

	} else if current.next == nil {
		// if current has no next
		current = nil
		l.head = current
		l.count--
	} else {
		// if there is a next, just replace the head with next
		tmp := current.next
		current = tmp
		l.head = current
		l.count--
	}
	return nil
}

// delete the tail of a linkedlist
func (l *LinkedList) DeleteTail() error {
	current := l.head
	// prev pointer behind the current
	prev := l.head
	if current == nil {
		return errors.New("linkedlist is empty")
	}
	if l.count < 2 {
		l.DeleteHead()
		return nil
	}
	for current != nil {
		prev = current
		current = current.next

	}
	prev.next = nil
	l.head = prev
	l.count--
	return nil

}

// delete a node from a location
// arg node position
func (l *LinkedList) DeleteANode(position int) error {
	current := l.head
	count := 1
	prev := current
	if l.count < position {
		return errors.New("position out of range")
	}

	for current != nil {
		if count == position-1 {
			tmp := current.next
			prev.next = tmp
			l.head = prev
			l.count--
			return nil
		}
		prev = current
		count++
		current = current.next

	}
	return nil
}

// clear a linkedlist
func (l *LinkedList) Clear() {
	l.head = nil
}
