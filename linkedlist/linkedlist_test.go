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
		assert.Equal(t, node.getVal(), 4)
		assert.Nil(t, node.getNext())
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

}
