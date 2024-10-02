package circularlinklist

import (
	"errors"
	"fmt"

	"github.com/Ayobami6/linkedlist_implementation/linkedlist"
)

type CircularLinkedList struct {
	head  *linkedlist.Node
	count int
	tail  *linkedlist.Node
}

func NewCircularLinkeList() *CircularLinkedList {
	return &CircularLinkedList{
		head:  nil,
		count: 0,
		tail:  nil,
	}
}

func (c *CircularLinkedList) InsertAtTheStart(value int) {
	node := linkedlist.NewNode(value)
	// if ll is not empty
	if c.head != nil {
		node.SetNext(c.head)
		c.head = node
		c.tail.SetNext(node)

	} else {
		c.head = node
		c.tail = c.head
		c.head.SetNext(c.tail)
		c.tail.SetNext(node)
	}
	c.count++

}

func (c *CircularLinkedList) InsertAtTheEnd(value int) {
	// set the new node to the tail next
	node := linkedlist.NewNode(value)
	// head next of the tail
	tmp := c.tail.GetNext()
	node.SetNext(tmp)
	c.tail.SetNext(node)
	c.count++

}

func (c *CircularLinkedList) LookUp(value int) any {
	node := c.head
	if node == nil {
		return nil
	}
	for node != nil {
		if node.GetVal() == value {
			return node.GetVal()
		}
		node = node.GetNext()
	}

	return nil

}

func (c *CircularLinkedList) GetCount() int {
	return c.count

}

func (c *CircularLinkedList) DumpList() {
	node := c.head
	for node != nil {
		fmt.Printf("Node val is %d\n", node.GetVal())
		node = node.GetNext()
	}
}

// delete head more like pop
func (c *CircularLinkedList) DeleteHead() error {
	current := c.head
	if current == nil {
		return errors.New("circular linkedlist is empty")
	}
	// get head next
	next := c.head.GetNext()
	c.head = next
	c.tail.SetNext(next)
	c.count--

	return nil
}

func (c *CircularLinkedList) DeleteTail() error {
	current := c.head
	prev := c.head

	if current == nil {
		return errors.New("linkedlist is empty")
	}
	if c.count < 2 {
		c.DeleteHead()
		return nil
	}

	// traverse the linkedlist
	for current != nil {
		prev = current
		current = current.GetNext()
	}
	prev.SetNext(c.head)
	c.tail.SetNext(prev)
	c.head = prev
	c.count--
	return nil

}

func (c *CircularLinkedList) DeleteANode(position int) error {
	current := c.head
	count := 1
	prev := current
	if c.count < position {
		return errors.New("position out of range")
	}

	for current != nil {
		if count == position {
			tmp := current.GetNext()
			prev.SetNext(tmp)
			c.head = prev
			c.count--
			return nil
		}
		prev = current
		count++
		current = current.GetNext()
	}
	return nil

}
