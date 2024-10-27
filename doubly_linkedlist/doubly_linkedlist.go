package doublylinkedlist

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
		d.count++
	} else {
		newNode.next = d.head
		d.head.prev = newNode
		d.head = newNode
		d.count++
	}
}

func (d *DoublyLinkedList) InsertAtTheEnd(data interface{}) {
	newNode := NewNode(data)
	if d.tail == nil {
		d.head = newNode
		d.tail = newNode
		d.count++
	} else {
		newNode.prev = d.tail
		d.tail.next = newNode
		d.tail = newNode
		d.count++
	}
}
