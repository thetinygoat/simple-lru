package dlist

import "errors"

// Node represents a dll node
type Node struct {
	data int
	next *Node
	prev *Node
}

// List represents a generic dll
type List struct {
	head     *Node
	tail     *Node
	keyspace map[string]*Node
}

// New instantiates a new dll
func New() *List {
	l := List{}
	l.head = &Node{data: -1}
	l.tail = &Node{data: -1}
	l.head.next = l.tail
	l.tail.prev = l.head
	return &l
}

// RemoveLast removes and returns last node from the list
func (l *List) RemoveLast() (*Node, error) {
	if l.head.next == l.tail {
		return nil, errors.New("Empty List")
	}
	n := l.tail.prev
	n.prev.next = l.tail
	l.tail.prev = n.prev
	n.next = nil
	n.prev = nil
	return n, nil
}

// AddFirst adds an item to the front of the list
func (l *List) AddFirst(data int) *Node {
	n := Node{data: data}
	n.next = l.head.next
	l.head.next = &n
	n.prev = l.head
	n.next.prev = &n
	return &n
}

// Update updates data in the node
func (n *Node) Update(val int) {
	n.data = val
}

// Data returns data present in the node
func (n *Node) Data() int {
	return n.data
}

// Detach removes node from the list
func (n *Node) Detach() {
	n.prev.next = n.next
	n.next.prev = n.prev
}
