package circularlinklist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCircularLinkedList(t *testing.T) {

	// test new circular linkedlist creates sucessfully
	t.Run("test create new instance of circular linkedlist", func(t *testing.T) {
		circularLinkedList := NewCircularLinkeList()
		assert.Nil(t, circularLinkedList.head)
		assert.Nil(t, circularLinkedList.tail)
		assert.Equal(t, circularLinkedList.count, 0)
	})

	// test insert to the start of a linkedlist
	t.Run("test insert a new node at the start the circular linkedlist", func(t *testing.T) {
		c := NewCircularLinkeList()
		c.InsertAtTheStart(3)
		assert.Equal(t, c.head.GetVal(), 3)
		assert.Equal(t, c.tail, c.head.GetNext())
		assert.Equal(t, c.tail.GetVal(), 3)
	})

	// test multiple insert
	t.Run("test multipe insertStart correctly inserted", func(t *testing.T) {
		c := NewCircularLinkeList()
		c.InsertAtTheStart(4)
		c.InsertAtTheStart(5)
		assert.Equal(t, c.count, 2)
		assert.Equal(t, c.head.GetVal(), c.tail.GetNext().GetVal())
		assert.Equal(t, c.head.GetVal(), 5)

	})

	// test look up value in a circular linkedlist
	t.Run("test look up value found", func(t *testing.T) {
		c := NewCircularLinkeList()
		c.InsertAtTheStart(6)
		c.InsertAtTheStart(7)
		val := c.LookUp(7)
		assert.Equal(t, val, 7)
	})

	// test look up mot found
	t.Run("test look not found", func(t *testing.T) {
		c := NewCircularLinkeList()
		val := c.LookUp(6)
		assert.Nil(t, val)
	})

	// test insert at the end
	t.Run("test insert at the end of a circular linkedlist", func(t *testing.T) {
		c := NewCircularLinkeList()
		c.InsertAtTheStart(5)
		c.InsertAtTheEnd(2)
		assert.Equal(t, c.head.GetVal(), 5)
		assert.Equal(t, c.GetCount(), 2)
	})

	// test delete head
	t.Run("delete a node from the head", func(t *testing.T) {
		// arrange
		c := NewCircularLinkeList()
		c.InsertAtTheStart(4)
		c.InsertAtTheStart(6)
		// act
		c.DeleteHead()

		// assert
		assert.Equal(t, c.head.GetVal(), 4)

	})

	// test delete tail
	t.Run("delete the tail node from a cricular linkedlist", func(t *testing.T) {
		c := NewCircularLinkeList()
		c.InsertAtTheStart(3)
		c.InsertAtTheStart(2)
		err := c.DeleteTail()
		assert.Nil(t, err)
		assert.Equal(t, c.GetCount(), 1)
		assert.Equal(t, c.tail.GetVal(), 2)
	})

	// test delete a specific node
	t.Run("delete a specific node", func(t *testing.T) {
		c := NewCircularLinkeList()
		c.InsertAtTheStart(3)
		c.InsertAtTheEnd(2)
		c.InsertAtTheStart(8)
		err := c.DeleteANode(2)
		assert.Equal(t, c.GetCount(), 2)
		assert.Nil(t, err)
	})
}
