package linkedlist

import "fmt"

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

func (n *Node) setNext(node *Node) {
	n.next = node
}

func (n *Node) getNext() *Node {
	return n.next
}

func (n *Node) getVal() int {
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
		newNode.setNext(l.head)
	}
	l.head = newNode
	l.count++
}

func (l *LinkedList) LookUp(value int) int {
	node := l.head
	var val int
	for node != nil {
		if node.getVal() == value {
			val = node.getVal()
		}
		node = node.getNext()
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
		node = node.getNext()
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
		node = node.getNext()
	}
}
