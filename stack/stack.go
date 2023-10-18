package stack

import (
	"fmt"
	"github.com/elias8/go-gather/base"
	"github.com/elias8/go-gather/list"
)

// Stack is a LIFO (last in, first out) data structure.
type Stack[T any] interface {
	base.Collection[T]

	// Push adds an element to the top of the stack.
	Push(T)

	// Pop removes and returns the top element of the stack.
	Pop() (*T, bool)

	// Peek returns the top element of the stack without removing it.
	Peek() (*T, bool)
}

type stack[T any] struct {
	linkedList list.LinkedList[T]
}

func New[T any]() Stack[T] {
	return &stack[T]{linkedList: list.NewLinkedList[T]()}
}

func (s stack[T]) Size() int {
	return s.linkedList.Size()
}

func (s stack[T]) IsEmpty() bool {
	return s.linkedList.IsEmpty()

}

func (s stack[T]) Contains(element T) bool {
	return s.linkedList.Contains(element)
}

func (s stack[T]) Values() []T {
	return s.linkedList.Values()
}

func (s stack[T]) Clear() {
	s.linkedList.Clear()
}

func (s stack[T]) String() string {
	str := "Stack(["
	for i, v := range s.linkedList.Values() {
		str += fmt.Sprintf("%v", v)
		if i != s.linkedList.Size()-1 {
			str += ", "
		}
	}
	return str + "])"
}

func (s stack[T]) Push(element T) {
	s.linkedList.Add(element)
}

func (s stack[T]) Pop() (*T, bool) {
	return s.linkedList.RemoveLast()
}

func (s stack[T]) Peek() (*T, bool) {
	return s.linkedList.GetLast()
}
