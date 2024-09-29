package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedList(t *testing.T) {

	t.Run("Test newnode setValues correctly", func(t *testing.T) {
		node := NewNode(4)
		assert.Equal(t, node.val, 4)
		assert.Nil(t, node.next)
		assert.Equal(t, node.GetVal(), 4)
		assert.Nil(t, node.GetNext())
	})

	t.Run("Test linkedlist inserts at the head", func(t *testing.T) {
		list := NewLinkedList()
		list.InsertAtTheStart(4)
		assert.Equal(t, list.head.val, 4)
		assert.Equal(t, list.count, 1)

	})

	t.Run("test insert at the end of the list", func(t *testing.T) {
		list := NewLinkedList()
		list.InsertAtTheStart(9)
		list.InsertAtTheEnd(4)
		assert.NotEqual(t, list.head.val, 4)
		assert.Equal(t, list.count, 2)
		assert.Equal(t, list.head.next.val, 4)
	})

	t.Run("delete a node from the beginning of an empty linkedlist", func(t *testing.T) {
		list := NewLinkedList()
		err := list.DeleteHead()
		assert.EqualError(t, err, "linkedlist is empty")
	})

	t.Run("delete a node from the beginning of a linkedlist", func(t *testing.T) {
		list := NewLinkedList()
		list.InsertAtTheStart(9)
		_ = list.DeleteHead()
		assert.Nil(t, list.head)

	})

	t.Run("delete the tail node from the linkedlist", func(t *testing.T) {
		list := NewLinkedList()
		list.InsertAtTheStart(8)
		list.InsertAtTheStart(5)
		list.DeleteTail()
		assert.Nil(t, list.head.next)
		assert.Equal(t, list.head.val, 8)

	})

	t.Run("delete the tail node from the linkedlist if head is the tail", func(t *testing.T) {
		list := NewLinkedList()
		list.InsertAtTheStart(5)
		list.DeleteTail()
		assert.Nil(t, list.head)
	})

	t.Run("delete a node from a node position in the linkedlist", func(t *testing.T) {
		list := NewLinkedList()
		list.InsertAtTheStart(9)
		list.InsertAtTheStart(6)
		list.InsertAtTheStart(4)
		list.InsertAtTheStart(2)
		list.InsertAtTheStart(3)
		// delete the second node with value of 6
		list.DeleteANode(2)
		assert.Equal(t, list.head.next.val, 2)
		assert.Equal(t, list.head.val, 3)
	})

}
