package dsa

import "iter"

type linkedNode struct {
	next *linkedNode
	Item any
}

type LinkedList struct {
	head  *linkedNode
	tail  *linkedNode
	count int
}

func newLinkedNode(item any) *linkedNode {
	return &linkedNode{
		Item: item,
	}
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func NewLinkedListUsingSlice(slice *[]any) *LinkedList {
	ll := LinkedList{}
	for _, value := range *slice {
		ll.Add(value)
	}

	return &ll
}

func (ll *LinkedList) Add(item any) {
	ll.count++
	if ll.head == nil {
		ll.head = newLinkedNode(item)
		ll.tail = ll.head
		return
	}

	ll.tail.next = newLinkedNode(item)
	ll.tail = ll.tail.next
}

func (ll *LinkedList) AddFirst(item any) {
	ll.count++
	first := newLinkedNode(item)
	first.next = ll.head
	ll.head = first
	if ll.tail == nil {
		ll.tail = ll.head
	}
}

func (ll *LinkedList) Delete(item any) {
	var prev *linkedNode
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

func (ll *LinkedList) DeleteAt(index int) {
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

func (ll *LinkedList) DeleteLast() {
	ll.DeleteAt(ll.count - 1) // This is unidirectional linked list (points to next), cannot go backwards
}

func (ll *LinkedList) GetFirst() any {
	if ll.head == nil {
		return nil
	}
	return ll.head.Item
}

func (ll *LinkedList) GetLast() any {
	if ll.tail == nil {
		return nil
	}
	return ll.tail.Item
}

func (ll *LinkedList) GetAt(index int) any {
	if index < 0 || index >= ll.count {
		return nil
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

func (ll *LinkedList) Exists(item any) bool {
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

func (ll *LinkedList) Clear() {
	ll.head = nil
	ll.tail = nil
	ll.count = 0
}

func (ll *LinkedList) All() iter.Seq[any] {
	return func(yield func(any) bool) {
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

func (ll *LinkedList) GetCount() int {
	return ll.count
}
