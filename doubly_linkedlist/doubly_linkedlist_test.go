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

}
