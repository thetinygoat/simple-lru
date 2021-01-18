package lru

import (
	"log"

	"github.com/thetinygoat/lld/design/pkg/dlist"
)

// Lru reprents an lru object
type Lru struct {
	list     *dlist.List
	size     int
	capacity int
	keyspace map[string]*dlist.Node
}

// New instantiates a new Lru object
func New(capacity int) *Lru {
	return &Lru{capacity: capacity, list: dlist.New(), keyspace: make(map[string]*dlist.Node)}
}

// Set inserts a key value pair
func (l *Lru) Set(key string, value int) {
	if n, ok := l.keyspace[key]; ok {
		n.Update(value)
		return
	}
	newNode := l.list.AddFirst(value)
	l.keyspace[key] = newNode
	l.size++
	if l.size > l.capacity {
		evicted, _ := l.list.RemoveLast()
		log.Printf("evicted key with value %d\n", evicted.Data())
	}
}

// Get returns the value for a particular key
func (l *Lru) Get(key string) int {
	if _, ok := l.keyspace[key]; !ok {
		return -1
	}
	node := l.keyspace[key]
	data := node.Data()
	node.Detach()
	l.list.AddFirst(node.Data())
	node = nil
	return data
}
