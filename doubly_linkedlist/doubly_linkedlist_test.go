package doublylinkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoublyLinkedList(t *testing.T) {

	t.Run("test new node creation", func(t *testing.T) {
		newNode := NewNode(4)
		assert.Equal(t, newNode.data, 4)
		assert.Nil(t, newNode.next)
		assert.Nil(t, newNode.prev)
	})

	t.Run("Test insert at the start of a dll", func(t *testing.T) {
		dll := NewDoublyLinkedList()
		dll.InsertAtTheStart(4)
		assert.Equal(t, dll.head.data, 4)
		assert.Nil(t, dll.head.next)
		assert.Equal(t, dll.count, 1)
	})

	t.Run("test insert node at the end of a linkelist", func(t *testing.T) {
		dll := NewDoublyLinkedList()
		dll.InsertAtTheStart(2)
		dll.InsertAtTheEnd(5)
		assert.Equal(t, dll.count, 2)
		assert.Equal(t, dll.head.data, 2)
		assert.Equal(t, dll.head.next.data, 5)
	})

	t.Run("Test delete a node", func(t *testing.T) {
		dll := NewDoublyLinkedList()
		dll.InsertAtTheStart(4)
		dll.InsertAtTheStart(2)
		dll.InsertAtTheStart(6)
		dll.DeleteANode(2)
		assert.Equal(t, dll.head.data, 6)
		assert.Equal(t, dll.head.next.data, 4)
	})

	t.Run("Test delete the tail node", func(t *testing.T) {
		dll := NewDoublyLinkedList()
		dll.InsertAtTheStart(3)
		dll.InsertAtTheStart(2)
		dll.InsertAtTheStart(1)
		dll.DeleteTail()
		assert.Equal(t, dll.count, 2)
		assert.Equal(t, dll.head.data, 1)
		assert.Equal(t, dll.head.next.data, 2)
	})

	t.Run("Test Delete Head", func(t *testing.T) {
		dll := NewDoublyLinkedList()
		dll.InsertAtTheStart(3)
		dll.InsertAtTheStart(2)
		dll.InsertAtTheStart(1)
		dll.DeleteHead()
		assert.Equal(t, dll.count, 2)
		assert.Equal(t, dll.head.data, 2)

	})

	// test lookup
	t.Run("Test lookup", func(t *testing.T) {
		dll := NewDoublyLinkedList()
		dll.InsertAtTheStart(2)
		dll.InsertAtTheStart(4)
		dll.InsertAtTheStart(5)
		data, err := dll.LookUp(4)
		assert.Equal(t, data, 4)
		assert.Nil(t, err)
	})

	t.Run("Test lookup not found", func(t *testing.T) {
		dll := NewDoublyLinkedList()
		dll.InsertAtTheStart(2)
		dll.InsertAtTheStart(4)
		dll.InsertAtTheStart(5)
		data, err := dll.LookUp(1)
		assert.Nil(t, data)
		assert.Contains(t, err.Error(), "data not found!")

	})

	// test dll empty
	t.Run("Test dll empty", func(t *testing.T) {
		dll := NewDoublyLinkedList()
		data, err := dll.LookUp(3)
		assert.Nil(t, data)
		assert.Contains(t, err.Error(), "linkedlist is empty")
	})

}
