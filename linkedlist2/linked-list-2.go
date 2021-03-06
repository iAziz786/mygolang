package linkedlist2

import (
	"errors"
)

var (
	ErrNotFound  = errors.New("item not found")
	ErrEmptyList = errors.New("empty list")
)

type node struct {
	value interface{}
	next  *node
	// prev  *node
}

type LinkedList struct {
	node *node
	size uint64
}

func newNode(val interface{}) *node {
	return &node{
		value: val,
	}
}

func New() *LinkedList {
	return &LinkedList{
		node: nil,
	}
}

func (ll LinkedList) Size() uint64 {
	return ll.size
}

func (ll LinkedList) Empty() bool {
	return ll.Size() == 0
}

// PushFront will add an item to the front of the linked list
func (ll *LinkedList) PushFront(val interface{}) {
	newNode := newNode(val)

	newNode.next = ll.node

	ll.node = newNode

	ll.setSize(ll.size + 1)
}

// PopFront will remove an item to the front of the linked list
func (ll *LinkedList) PopFront() error {
	if ll.Size() <= 0 {
		return ErrEmptyList
	}

	ll.node = ll.node.next
	ll.setSize(ll.size - 1)

	return nil
}

// PushBack will add an item to the end of the linked list
func (ll *LinkedList) PushBack(val interface{}) {
	lastNode := &ll.node

	for *lastNode != nil {
		lastNode = &(*lastNode).next
	}

	*lastNode = newNode(val)
	ll.setSize(ll.size + 1)
}

// PopBack will remove an item from the end of the linked list
func (ll *LinkedList) PopBack() {
	lastNode := &ll.node

	for *lastNode != nil && (*lastNode).next != nil {
		lastNode = &(*lastNode).next
	}

	*lastNode = nil

	ll.setSize(ll.size - 1)
}

// Push will add an item to the front of the linked list
func (ll *LinkedList) ValueAt(index uint64) (interface{}, error) {
	var nodeToFind *node

	// Out of range should return immediately
	if index > ll.Size() {
		return 0, ErrNotFound
	}

	for idx := uint64(0); index >= idx; idx++ {
		if nodeToFind == nil {
			nodeToFind = ll.node
		} else {
			nodeToFind = nodeToFind.next
		}
	}

	if nodeToFind == nil {
		return 0, ErrNotFound
	}

	return nodeToFind.value, nil
}

// Insert will add an item to the index of the linked list. It panics if index
// is out of range.
func (ll *LinkedList) Insert(idx int, val interface{}) {
	findingNode := &ll.node

	for i := 0; i < idx; i++ {
		findingNode = &((*findingNode).next)
	}

	prevNode := *findingNode
	*findingNode = newNode(val)
	(*findingNode).next = prevNode

	ll.setSize(ll.size + 1)
}

// Reverse will reverse the linked list
func (ll *LinkedList) Reverse() {
	findingNode := ll.node
	head := &ll.node

	var prevNode *node = nil
	for findingNode != nil {
		next := findingNode.next
		findingNode.next = prevNode
		prevNode = findingNode
		findingNode = next
	}

	*head = prevNode
}

// Erase will remove the item on the provided index, if any
func (ll *LinkedList) Erase(idx uint64) {
	removingNode := &ll.node

	var nextNode **node = &(*removingNode).next
	for i := 0; uint64(i) < idx; i++ {
		removingNode = nextNode
		nextNode = &(*nextNode).next
	}

	*removingNode = *nextNode
	ll.setSize(ll.size - 1)
}

// setSize handles the edge cases in case of setting the size. If you provide
// a non unsigned value it will set to zero since that's the minimal possible
// size can have a linked list
func (ll *LinkedList) setSize(size uint64) {
	ll.size = size
	ll.size += (ll.size >> 63)
}
