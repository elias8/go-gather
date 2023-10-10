package list

import (
	"fmt"
	"reflect"
)

type node[T any] struct {
	value T
	prev  *node[T]
	next  *node[T]
}

func newNode[T any](value T) *node[T] {
	return &node[T]{value: value}
}

// LinkedList represents a doubly linked list.
type linkedList[T any] struct {
	head *node[T]
	tail *node[T]
	size int
}

// NewLinkedList returns an empty LinkedList.
func NewLinkedList[T any]() LinkedList[T] {
	return &linkedList[T]{}
}

func (l *linkedList[T]) Size() int {
	return l.size
}

func (l *linkedList[T]) IsEmpty() bool {
	return l.size == 0
}

func (l *linkedList[T]) Contains(element T) bool {
	for current := l.head; current != nil; current = current.next {
		if reflect.DeepEqual(current.value, element) {
			return true
		}
	}
	return false
}

func (l *linkedList[T]) String() string {
	s := "LinkedList(["
	for current := l.head; current != nil; current = current.next {
		if current.prev != nil {
			s += " <-> "
		}
		s += fmt.Sprintf("%v", current.value)
	}
	s += "])"
	return s
}

func (l *linkedList[T]) Values() []T {
	var slice []T
	for current := l.head; current != nil; current = current.next {
		slice = append(slice, current.value)
	}
	return slice
}

func (l *linkedList[T]) Add(element T) {
	node := newNode(element)
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		prev := l.tail
		l.tail = node
		prev.next = node
		node.prev = prev
	}
	l.size++
}

func (l *linkedList[T]) AddFirst(element T) {
	node := newNode(element)
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		l.head.prev = node
		node.next = l.head
		l.head = node
	}
	l.size++
}

func (l *linkedList[T]) AddLast(element T) {
	l.Add(element)
}

func (l *linkedList[T]) Clear() {
	l.head = nil
	l.tail = nil
	l.size = 0
}

func (l *linkedList[T]) Remove(element T) bool {
	current := l.head
	for current != nil {
		if reflect.DeepEqual(current.value, element) {
			if current.prev != nil {
				current.prev.next = current.next
			} else {
				l.head = current.next
			}
			if current.next != nil {
				current.next.prev = current.prev
			} else {
				l.tail = current.prev
			}
			l.size--
			return true
		}
		current = current.next
	}
	return false
}

func (l *linkedList[T]) RemoveFirst() (*T, bool) {
	if l.head != nil {
		temp := l.head
		l.head = l.head.next
		if l.head != nil {
			l.head.prev = nil
		}
		l.size--
		return &temp.value, true
	}
	return nil, false
}

func (l *linkedList[T]) RemoveLast() (*T, bool) {
	if l.head == nil {
		return nil, false
	} else if l.head == l.tail {
		removed := l.tail.value
		l.head = nil
		l.tail = nil
		l.size--
		return &removed, true
	} else {
		removed := l.tail.value
		l.tail = l.tail.prev
		l.tail.next = nil
		l.size--
		return &removed, true
	}
}

func (l *linkedList[T]) Set(index int, element T) (*T, bool) {
	position := 0
	current := l.head
	for current != nil && position <= l.size && position <= index {
		if position == index {
			replaced := current.value
			current.value = element
			return &replaced, true
		}
		current = current.next
		position++
	}
	return nil, false
}

func (l *linkedList[T]) Get(index int) (*T, bool) {
	position := 0
	current := l.head
	for current != nil && position <= l.size && position <= index {
		if position == index {
			return &current.value, true
		}
		current = current.next
	}
	return nil, false
}

func (l *linkedList[T]) GetFirst() (*T, bool) {
	if l.head == nil {
		return nil, false
	}
	return &l.head.value, true
}

func (l *linkedList[T]) GetLast() (*T, bool) {
	if l.tail == nil {
		return nil, false
	}
	return &l.tail.value, true
}

func (l *linkedList[T]) IndexOf(element T) (int, bool) {
	position := 0
	current := l.head
	for current != nil {
		if reflect.DeepEqual(current.value, element) {
			return position, true
		}
		current = current.next
		position++
	}
	return -1, false
}

func (l *linkedList[T]) LastIndexOf(element T) (int, bool) {
	position := l.size - 1
	current := l.tail
	for current != nil {
		if reflect.DeepEqual(current.value, element) {
			return position, true
		}
		current = current.prev
		position--
	}
	return -1, false
}

func (l *linkedList[T]) Reverse() {
	current := l.head
	for current != nil {
		next := current.next
		current.next = current.prev
		current.prev = next
		current = next
	}
	l.tail, l.head = l.head, l.tail
}
