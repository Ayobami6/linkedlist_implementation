package circularlinklist

import "github.com/Ayobami6/linkedlist_implementation/linkedlist"

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
	node := linkedlist.NewNode(5)
	if c.head != nil {
		node.SetNext(c.head)
	}
	c.tail = c.head
	c.head = node
	c.tail.SetNext(node)

}
