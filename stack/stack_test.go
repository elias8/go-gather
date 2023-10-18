package stack

import (
	"slices"
	"testing"
)

func TestNew(t *testing.T) {
	stack := New[int]()

	if stack == nil {
		t.Fatalf("Expected stack to not be nil")
	}
	if !stack.IsEmpty() {
		t.Fatalf("Expected stack to be empty")
	}
}

func TestStack_Push(t *testing.T) {
	stack := New[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	if stack.Size() != 3 {
		t.Fatalf("Expected stack to have 3 elements, but got %v", stack.Size())
	}
}

func TestStack_Clear(t *testing.T) {
	stack := New[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Clear()

	if !stack.IsEmpty() {
		t.Fatalf("Expected stack to be empty after clearing")
	}
}

func TestStack_Contains(t *testing.T) {
	stack := New[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	if !stack.Contains(1) {
		t.Fatalf("Expected stack to contain 1")
	}
	if !stack.Contains(2) {
		t.Fatalf("Expected stack to contain 2")
	}
	if !stack.Contains(3) {
		t.Fatalf("Expected stack to contain 3")
	}
	if stack.Contains(4) {
		t.Fatalf("Expected stack to not contain 4")
	}
}

func TestStack_IsEmpty(t *testing.T) {
	stack := New[int]()
	if !stack.IsEmpty() {
		t.Fatalf("Expected stack to be empty")
	}

	stack.Push(1)

	if stack.IsEmpty() {
		t.Fatalf("Expected stack to not be empty")
	}
}

func TestStack_Peek(t *testing.T) {
	stack := New[int]()

	if _, found := stack.Peek(); found {
		t.Fatalf("Expected stack.Peek() to not find an element")
	}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	if stack.Size() != 3 {
		t.Fatalf("Expected stack to have 3 elements, but got %v", stack.Size())
	}

	val, found := stack.Peek()
	if !found {
		t.Fatalf("Expected stack.Peek() to find an element")
	}
	if *val != 3 {
		t.Fatalf("Expected stack.Peek() to return 3")
	}
	if stack.Size() != 3 {
		t.Fatalf("Expected stack to have 3 elements after calling stack.Peek(), but got %v", stack.Size())
	}
}

func TestStack_Pop(t *testing.T) {
	stack := New[int]()

	if _, found := stack.Pop(); found {
		t.Fatalf("Expected stack.Pop() to not find an element")
	}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	if stack.Size() != 3 {
		t.Fatalf("Expected stack to have 3 elements, but got %v", stack.Size())
	}

	val, found := stack.Pop()
	if !found {
		t.Fatalf("Expected stack.Pop() to find an element")
	}
	if *val != 3 {
		t.Fatalf("Expected stack.Pop() to return 3")
	}
	if stack.Size() != 2 {
		t.Fatalf("Expected stack to have 2 elements after calling stack.Pop(), but got %v", stack.Size())
	}
}

func TestStack_Size(t *testing.T) {
	stack := New[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	if stack.Size() != 3 {
		t.Fatalf("Expected stack to have 3 elements, but got %v", stack.Size())
	}
}

func TestStack_String(t *testing.T) {
	stack := New[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	if stack.String() != "Stack([1, 2, 3])" {
		t.Fatalf("Expected stack.String() to return 'Stack([1, 2, 3])', but got %v", stack.String())
	}
}

func TestStack_Values(t *testing.T) {
	stack := New[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	expected := []int{1, 2, 3}
	slice := stack.Values()
	if !slices.Equal(slice, expected) {
		t.Fatalf("Expected stack.Slice() to return %v, but got %v", expected, slice)
	}
}
