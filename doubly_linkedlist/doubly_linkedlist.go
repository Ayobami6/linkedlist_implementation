package doublylinkedlist

import (
	"errors"
	"fmt"
)

// Node of Doubly linkedlist

type Node struct {
	data interface{}
	next *Node
	prev *Node
}

func NewNode(data interface{}) *Node {
	return &Node{
		data: data,
		next: nil,
		prev: nil,
	}
}

func (n *Node) GetNext() *Node {
	return n.next
}

func (n *Node) GetPrev() *Node {
	return n.prev
}

// GetData
func (n *Node) GetData() interface{} {
	return n.data
}

type DoublyLinkedList struct {
	head  *Node
	count int
	tail  *Node
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{
		head:  nil,
		count: 0,
		tail:  nil,
	}
}

// Insert at the start
func (d *DoublyLinkedList) InsertAtTheStart(data interface{}) {
	// if head is nil, just put node
	newNode := NewNode(data)
	if d.head == nil {
		d.head = newNode
		d.tail = newNode
	} else {
		newNode.next = d.head
		d.head.prev = newNode
		d.head = newNode
	}
	d.count++
}

func (d *DoublyLinkedList) InsertAtTheEnd(data interface{}) {
	newNode := NewNode(data)
	if d.tail == nil {
		d.head = newNode
		d.tail = newNode
	} else {
		newNode.prev = d.tail
		d.tail.next = newNode
		d.tail = newNode
	}
	d.count++
}

// DeleteANode,  deletes a node from a particular position
func (d *DoublyLinkedList) DeleteANode(position int) error {
	counter := 1
	if d.count < position {
		return errors.New("position out of range")
	}
	head := d.head
	prev := head
	for head != nil {
		if counter == position {
			tmp := head.next
			prev.next = tmp
			d.head = prev
			d.count--
			return nil
		}
		prev = head
		counter++
		head = head.next

	}
	return errors.New("unable to delete node")
}

// DeleteTail, deletes the tail node
func (d *DoublyLinkedList) DeleteTail() error {
	if d.head == nil {
		// if head is nil
		return errors.New("linkedlist is empty")

	}
	if d.head == d.tail {
		d.head = nil
		d.tail = nil
	} else {

		prev := d.head
		current := d.head
		for current != nil {
			prev = current
			current = current.next
		}
		prev.next = nil
		d.tail = prev
	}
	d.count--
	return nil

}

// DeleteHead, deletes the head pf a ll
func (d *DoublyLinkedList) DeleteHead() error {
	if d.head == nil {
		return errors.New("linkedlist is empty")
	}
	tmp := d.head.next
	d.head = tmp
	d.tail = d.head
	d.count--
	return nil

}

func (d *DoublyLinkedList) LookUp(data interface{}) (interface{}, error) {
	if d.head == nil {
		return nil, errors.New("linkedlist is empty")
	}
	head := d.head
	for head != nil {
		if head.data == data {
			return head.data, nil
		}
		head = head.GetNext()
	}
	return nil, errors.New("data not found!")

}

func (d *DoublyLinkedList) DumpList() {
	head := d.head
	for head != nil {
		fmt.Println(head.data)
		head = head.GetNext()
	}
}
