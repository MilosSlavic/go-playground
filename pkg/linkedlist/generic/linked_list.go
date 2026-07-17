// Package generic implements a type-safe singly linked list using Go
// generics. See the sibling package (pkg/linkedlist) for the any-based
// equivalent.
package generic

import "iter"

type linkedNode[T comparable] struct {
	next *linkedNode[T]
	Item T
}

// LinkedList is a singly linked list holding values of type T.
type LinkedList[T comparable] struct {
	head  *linkedNode[T]
	tail  *linkedNode[T]
	count int
}

func newLinkedNode[T comparable](item T) *linkedNode[T] {
	return &linkedNode[T]{
		Item: item,
	}
}

// NewLinkedList returns an empty LinkedList.
func NewLinkedList[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{}
}

// NewLinkedListUsingSlice returns a LinkedList populated with the elements
// of slice, in order.
func NewLinkedListUsingSlice[T comparable](slice *[]T) *LinkedList[T] {
	ll := LinkedList[T]{}
	for _, value := range *slice {
		ll.Add(value)
	}

	return &ll
}

// zeroValue returns the zero value of T.
func zeroValue[T comparable]() T {
	var zero T
	return zero
}

// Add appends item to the end of the list.
func (ll *LinkedList[T]) Add(item T) {
	ll.count++
	if ll.head == nil {
		ll.head = newLinkedNode(item)
		ll.tail = ll.head
		return
	}

	ll.tail.next = newLinkedNode(item)
	ll.tail = ll.tail.next
}

// AddFirst inserts item at the head of the list.
func (ll *LinkedList[T]) AddFirst(item T) {
	ll.count++
	first := newLinkedNode(item)
	first.next = ll.head
	ll.head = first
	if ll.tail == nil {
		ll.tail = ll.head
	}
}

// Delete removes all nodes whose Item equals item.
func (ll *LinkedList[T]) Delete(item T) {
	var prev *linkedNode[T]
	node := ll.head

	for node != nil {
		if node.Item == item {
			if prev == nil {
				ll.head = node.next
			} else {
				prev.next = node.next
			}
			if node == ll.tail {
				ll.tail = prev
			}
			ll.count--
			node = node.next
			continue
		}
		prev = node
		node = node.next
	}
}

// DeleteAt removes the node at index. Out-of-range indexes are a no-op.
func (ll *LinkedList[T]) DeleteAt(index int) {
	if index < 0 || index >= ll.count {
		return
	}

	if index == 0 {
		ll.head = ll.head.next
		if ll.head == nil {
			ll.tail = nil
		}
		ll.count--
		return
	}

	prev := ll.head
	for range index - 1 {
		prev = prev.next
	}

	target := prev.next
	prev.next = target.next
	if target == ll.tail {
		ll.tail = prev
	}
	ll.count--
}

// DeleteLast removes the tail node. O(n): the list is singly linked, so
// finding the new tail requires a full traversal from head.
func (ll *LinkedList[T]) DeleteLast() {
	ll.DeleteAt(ll.count - 1)
}

// GetFirst returns the head item, or the zero value of T if the list is empty.
func (ll *LinkedList[T]) GetFirst() T {
	if ll.head == nil {
		return zeroValue[T]()
	}
	return ll.head.Item
}

// GetLast returns the tail item, or the zero value of T if the list is empty.
func (ll *LinkedList[T]) GetLast() T {
	if ll.tail == nil {
		return zeroValue[T]()
	}
	return ll.tail.Item
}

// GetAt returns the item at index, or the zero value of T if index is out of range.
func (ll *LinkedList[T]) GetAt(index int) T {
	if index < 0 || index >= ll.count {
		return zeroValue[T]()
	}

	if index == 0 {
		return ll.head.Item
	}

	node := ll.head
	for range index {
		node = node.next
	}

	return node.Item
}

// Exists reports whether item is present in the list.
func (ll *LinkedList[T]) Exists(item T) bool {
	var i int = 0
	node := ll.head
	for i < ll.count {
		if node.Item == item {
			return true
		}

		node = node.next
		i++
	}

	return false
}

// Clear removes all items from the list.
func (ll *LinkedList[T]) Clear() {
	ll.head = nil
	ll.tail = nil
	ll.count = 0
}

// All returns an iterator over the list's items, from head to tail.
func (ll *LinkedList[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		node := ll.head
		for {
			if node == nil {
				return
			}

			if !yield(node.Item) {
				return
			}

			node = node.next
		}
	}
}

// GetCount returns the number of items in the list.
func (ll *LinkedList[T]) GetCount() int {
	return ll.count
}
