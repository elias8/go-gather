package list

import (
	"fmt"
	"reflect"
	"testing"
)

func intPtr(val int) *int {
	return &val
}

type linkedListScenario[T any] struct {
	name     string
	values   []T
	expected []T
}

func (s linkedListScenario[T]) setup(list LinkedList[T]) {
	for _, value := range s.values {
		list.Add(value)
	}
}

func (s linkedListScenario[T]) test(t *testing.T, list LinkedList[T]) {
	s.testSize(t, list)
	s.testValues(t, list)
	s.testFirst(t, list)
	s.testLast(t, list)
	s.testNodeConstruction(t, list)
}

func (s linkedListScenario[T]) testSize(t *testing.T, list LinkedList[T]) {
	if size := list.Size(); size != len(s.expected) {
		t.Fatalf("Expected list size to be %d, but found %d", len(s.expected), size)
	}
}

func (s linkedListScenario[T]) testValues(t *testing.T, list LinkedList[T]) {
	if values := list.Values(); !reflect.DeepEqual(values, s.expected) {
		t.Fatalf("Expected list to be %v, but found %v", s.expected, values)
	}
}

func (s linkedListScenario[T]) testFirst(t *testing.T, list LinkedList[T]) {
	if len(s.expected) > 0 {
		if first, ok := list.GetFirst(); !ok || !reflect.DeepEqual(*first, s.expected[0]) {
			t.Fatalf("Expected first element to be %v, but found %v", s.expected[0], *first)
		}
	}
}

func (s linkedListScenario[T]) testLast(t *testing.T, list LinkedList[T]) {
	if len(s.expected) > 0 {
		if last, ok := list.GetLast(); !ok || !reflect.DeepEqual(*last, s.expected[len(s.expected)-1]) {
			t.Fatalf("Expected last element to be %v, but found %v", s.expected[len(s.expected)-1], *last)
		}
	}
}

func (s linkedListScenario[T]) testNodeConstruction(t *testing.T, list LinkedList[T]) {
	ll := list.(*linkedList[T])
	current := ll.head
	for i := 0; i < len(s.expected); i++ {
		if current == nil {
			t.Fatalf("Expected current to be not nil")
		}
		if !reflect.DeepEqual(current.value, s.expected[i]) {
			t.Fatalf("Expected current.value to be %v, but found %v", s.expected[i], current.value)
		}
		if i == 0 {
			if current.prev != nil {
				t.Fatalf("Expected head.prev to be nil, but found %v", current.prev)
			}
			if ll.head != current {
				t.Fatalf("Expected head to be the first element (%v), but found %v", s.expected[i], ll.head.value)
			}
		} else {
			if current.prev == nil {
				t.Fatalf("Expected (middle) current.prev to be not nil")
			}
			if current.prev.next != current {
				t.Fatalf("Expected (middle) current.prev.next to be current")
			}
		}
		if i == len(s.expected)-1 {
			if current.next != nil {
				t.Fatalf("Expected tail.next to be nil, but found %v", current.next)
			}
			if ll.tail != current {
				t.Fatalf("Expected tail to be the last element (%v), but found %v", s.expected[i], ll.tail.value)
			}
		} else {
			if current.next == nil {
				t.Fatalf("Expected (middle) current.next to be not nil")
			}
			if current.next.prev != current {
				t.Fatalf("Expected (middle) current.next.prev to be current")
			}
		}
		current = current.next
	}
}

func TestNewLinkedList(t *testing.T) {
	ll := NewLinkedList[int]()
	if ll.Size() != 0 {
		t.Fatalf("expected size to be 0, got %d", ll.Size())
	}
	if !ll.IsEmpty() {
		t.Fatalf("expected linked list to be empty")
	}
}

func TestLinkedList_Add(t *testing.T) {
	scenarios := []struct {
		linkedListScenario[int]
		value int
	}{
		{
			linkedListScenario: linkedListScenario[int]{
				name:     "add to empty list",
				values:   []int{},
				expected: []int{1},
			},
			value: 1,
		},
		{
			linkedListScenario: linkedListScenario[int]{
				name:     "add to list with one element",
				values:   []int{1},
				expected: []int{1, 2},
			},
			value: 2,
		},
		{
			linkedListScenario: linkedListScenario[int]{
				name:     "add to list with multiple element",
				values:   []int{1, 2, 3},
				expected: []int{1, 2, 3, 4},
			},
			value: 4,
		},
	}

	for _, s := range scenarios {
		t.Run(s.name, func(t *testing.T) {
			ll := NewLinkedList[int]()
			s.setup(ll)
			ll.Add(s.value)
			s.test(t, ll)
		})
	}
}

func TestLinkedList_AddFirst(t *testing.T) {
	scenarios := []struct {
		linkedListScenario[int]
		value int
	}{
		{
			linkedListScenario: linkedListScenario[int]{
				name:     "add to empty list",
				values:   []int{},
				expected: []int{1},
			},
			value: 1,
		},
		{
			linkedListScenario: linkedListScenario[int]{
				name:     "add to list with one element",
				values:   []int{1},
				expected: []int{2, 1},
			},
			value: 2,
		},
		{
			linkedListScenario: linkedListScenario[int]{
				name:     "add to list with multiple element",
				values:   []int{1, 2, 3},
				expected: []int{4, 1, 2, 3},
			},
			value: 4,
		},
	}

	for _, s := range scenarios {
		t.Run(s.name, func(t *testing.T) {
			ll := NewLinkedList[int]()

			s.setup(ll)
			ll.AddFirst(s.value)
			s.test(t, ll)
		})
	}
}

func TestLinkedList_AddLast(t *testing.T) {
	scenarios := []struct {
		linkedListScenario[int]
		value int
	}{
		{
			linkedListScenario: linkedListScenario[int]{
				name:     "add to empty list",
				values:   []int{},
				expected: []int{1},
			},
			value: 1,
		},
		{
			linkedListScenario: linkedListScenario[int]{
				name:     "add to list with one element",
				values:   []int{1},
				expected: []int{1, 2},
			},
			value: 2,
		},
		{
			linkedListScenario: linkedListScenario[int]{
				name:     "add to list with multiple element",
				values:   []int{1, 2, 3},
				expected: []int{1, 2, 3, 4},
			},
			value: 4,
		},
	}

	for _, s := range scenarios {
		t.Run(s.name, func(t *testing.T) {
			ll := NewLinkedList[int]()

			s.setup(ll)
			ll.AddLast(s.value)
			s.test(t, ll)
		})
	}
}

func TestLinkedList_RemoveFirst(t *testing.T) {
	scenarios := []struct {
		linkedListScenario[int]
		value int
	}{
		{
			linkedListScenario: linkedListScenario[int]{
				name:     "remove from empty list",
				values:   []int{},
				expected: nil,
			},
		},
		{
			linkedListScenario: linkedListScenario[int]{
				name:     "remove from list with one element",
				values:   []int{1},
				expected: nil,
			},
			value: 1,
		},
		{
			linkedListScenario: linkedListScenario[int]{
				name:     "remove from list with multiple element",
				values:   []int{1, 2, 3},
				expected: []int{2, 3},
			},
			value: 1,
		},
	}

	for _, s := range scenarios {
		t.Run(s.name, func(t *testing.T) {
			ll := NewLinkedList[int]()

			s.setup(ll)
			if removed, ok := ll.RemoveFirst(); ok && !reflect.DeepEqual(s.value, *removed) {
				t.Fatalf("Expected removed element to be %v, but found %v", s.value, *removed)
			}
			s.test(t, ll)
		})
	}
}

func TestLinkedList_RemoveLast(t *testing.T) {
	scenarios := []struct {
		linkedListScenario[int]
		value int
	}{
		{
			linkedListScenario: linkedListScenario[int]{
				name:     "remove from empty list",
				values:   []int{},
				expected: nil,
			},
		},
		{
			linkedListScenario: linkedListScenario[int]{
				name:     "remove from list with one element",
				values:   []int{1},
				expected: nil,
			},
			value: 1,
		},
		{
			linkedListScenario: linkedListScenario[int]{
				name:     "remove from list with multiple element",
				values:   []int{1, 2, 3},
				expected: []int{1, 2},
			},
			value: 3,
		},
	}

	for _, s := range scenarios {
		t.Run(s.name, func(t *testing.T) {
			ll := NewLinkedList[int]()

			s.setup(ll)
			if removed, ok := ll.RemoveLast(); ok && !reflect.DeepEqual(s.value, *removed) {
				t.Fatalf("Expected removed element to be %v, but found %v", s.value, *removed)
			}
			s.test(t, ll)
		})
	}
}

func TestLinkedList_Remove(t *testing.T) {
	scenarios := []struct {
		linkedListScenario[int]
		value int
	}{
		{
			linkedListScenario: linkedListScenario[int]{
				name:     "remove from empty list",
				values:   []int{},
				expected: nil,
			},
			value: 1,
		},
		{
			linkedListScenario: linkedListScenario[int]{
				name:     "remove element from list with one element",
				values:   []int{1},
				expected: nil,
			},
			value: 1,
		},
		{
			linkedListScenario: linkedListScenario[int]{
				name:     "remove last element from list with two element",
				values:   []int{1, 2},
				expected: []int{1},
			},
			value: 2,
		},
		{
			linkedListScenario: linkedListScenario[int]{
				name:     "remove first element from list with two element",
				values:   []int{1, 2},
				expected: []int{2},
			},
			value: 1,
		},
		{
			linkedListScenario: linkedListScenario[int]{
				name:     "remove middle element from list with multiple element",
				values:   []int{1, 2, 3},
				expected: []int{1, 3},
			},
			value: 2,
		},
	}

	for _, s := range scenarios {
		t.Run(s.name, func(t *testing.T) {
			ll := NewLinkedList[int]()

			s.setup(ll)
			if removed := ll.Remove(s.value); len(s.values) > 0 && !removed {
				t.Fatalf("Expected %v to be removed, but it wasn't", s.value)
			}
			s.test(t, ll)
		})
	}
}

func TestLinkedList_Reverse(t *testing.T) {
	scenarios := []linkedListScenario[int]{
		{
			name:     "reverse empty list",
			values:   []int{},
			expected: nil,
		},
		{
			name:     "reverse list with one element",
			values:   []int{1},
			expected: []int{1},
		},
		{
			name:     "reverse list with two element",
			values:   []int{1, 2},
			expected: []int{2, 1},
		},
		{
			name:     "reverse list with multiple element",
			values:   []int{1, 2, 3, 4},
			expected: []int{4, 3, 2, 1},
		},
	}

	for _, s := range scenarios {
		t.Run(s.name, func(t *testing.T) {
			ll := NewLinkedList[int]()

			s.setup(ll)
			ll.Reverse()
			s.test(t, ll)
			fmt.Println(ll)
		})
	}
}

func TestLinkedList_Get(t *testing.T) {
	scenarios := []struct {
		linkedListScenario[int]
		index int
	}{
		{
			linkedListScenario: linkedListScenario[int]{
				name:     "get from empty list",
				values:   []int{},
				expected: nil,
			},
			index: 0,
		},
		{
			linkedListScenario: linkedListScenario[int]{
				name:     "get from list with one element",
				values:   []int{1},
				expected: []int{1},
			},
			index: 0,
		},
		{
			linkedListScenario: linkedListScenario[int]{
				name:     "get firs element from list with multiple element",
				values:   []int{1, 2, 3, 4},
				expected: []int{1, 2, 3, 4},
			},
			index: 0,
		},
		{
			linkedListScenario: linkedListScenario[int]{
				name:     "get middle element from list with multiple element",
				values:   []int{1, 2, 3, 4},
				expected: []int{1, 2, 3, 4},
			},
			index: 2,
		},
		{
			linkedListScenario: linkedListScenario[int]{
				name:     "get last element from list with multiple element",
				values:   []int{1, 2, 3, 4},
				expected: []int{1, 2, 3, 4},
			},
			index: 3,
		},
	}

	for _, s := range scenarios {
		t.Run(s.name, func(t *testing.T) {
			ll := NewLinkedList[int]()

			s.setup(ll)
			if element, ok := ll.Get(s.index); ok && !reflect.DeepEqual(s.expected[s.index], *element) {
				t.Fatalf("Expected element to be %v, but found %v", s.expected[s.index], *element)
			}
			s.test(t, ll)
		})
	}
}

func TestLinkedList_GetFirst(t *testing.T) {
	scenarios := []linkedListScenario[int]{
		{
			name:     "get from empty list",
			values:   []int{},
			expected: nil,
		},
		{
			name:     "get from list with one element",
			values:   []int{1},
			expected: []int{1},
		},
		{
			name:     "get firs element from list with multiple element",
			values:   []int{1, 2, 3, 4},
			expected: []int{1, 2, 3, 4},
		},
	}

	for _, s := range scenarios {
		t.Run(s.name, func(t *testing.T) {
			ll := NewLinkedList[int]()

			s.setup(ll)
			if element, ok := ll.GetFirst(); ok && !reflect.DeepEqual(s.expected[0], *element) {
				t.Fatalf("Expected element to be %v, but found %v", s.expected[0], *element)
			}
			s.test(t, ll)
		})
	}
}

func TestLinkedList_GetLast(t *testing.T) {
	scenarios := []linkedListScenario[int]{
		{
			name:     "get from empty list",
			values:   []int{},
			expected: nil,
		},
		{
			name:     "get from list with one element",
			values:   []int{1},
			expected: []int{1},
		},
		{
			name:     "get firs element from list with multiple element",
			values:   []int{1, 2, 3, 4},
			expected: []int{1, 2, 3, 4},
		},
	}

	for _, s := range scenarios {
		t.Run(s.name, func(t *testing.T) {
			ll := NewLinkedList[int]()

			s.setup(ll)
			if element, ok := ll.GetLast(); ok && !reflect.DeepEqual(s.expected[len(s.expected)-1], *element) {
				t.Fatalf("Expected element to be %v, but found %v", s.expected[len(s.expected)-1], *element)
			}
			s.test(t, ll)
		})
	}
}

func TestLinkedList_IndexOf(t *testing.T) {
	scenarios := []struct {
		linkedListScenario[int]
		value int
	}{
		{
			linkedListScenario: linkedListScenario[int]{
				name:   "index of element in empty list",
				values: []int{},
			},
			value: 1,
		},
		{
			linkedListScenario: linkedListScenario[int]{
				name:   "index of element in list with one element",
				values: []int{1},
			},
			value: 1,
		},
		{
			linkedListScenario: linkedListScenario[int]{
				name:   "index of element in list with multiple element",
				values: []int{1, 2, 3, 4},
			},
			value: 3,
		},
		{
			linkedListScenario: linkedListScenario[int]{
				name:   "index of element not in list with multiple element",
				values: []int{1, 2, 3, 4},
			},
			value: 5,
		},
	}

	for _, s := range scenarios {
		t.Run(s.name, func(t *testing.T) {
			ll := NewLinkedList[int]()
			s.setup(ll)

			// find first index of element in values
			expected := -1
			for i, v := range s.values {
				if v == s.value {
					expected = i
					break
				}
			}

			index, ok := ll.IndexOf(s.value)
			if expected != -1 && (!ok || index != expected) {
				t.Fatalf("Expected index to be %d, but found %d", expected, index)
			}
			if expected == -1 && ok {
				t.Fatalf("Expected element to not be found, but found %d", index)
			}
		})
	}
}

func TestLinkedList_LastIndexOf(t *testing.T) {
	scenarios := []struct {
		linkedListScenario[int]
		value int
	}{
		{
			linkedListScenario: linkedListScenario[int]{
				name:   "last index of element in empty list",
				values: []int{},
			},
			value: 1,
		},
		{
			linkedListScenario: linkedListScenario[int]{
				name:   "last index of element in list with one element",
				values: []int{1},
			},
			value: 1,
		},
		{
			linkedListScenario: linkedListScenario[int]{
				name:   "last index of element in list with multiple element",
				values: []int{1, 2, 3, 4},
			},
			value: 3,
		},
		{
			linkedListScenario: linkedListScenario[int]{
				name:   "last index of element not in list with multiple element",
				values: []int{1, 2, 3, 4},
			},
			value: 5,
		},
	}

	for _, s := range scenarios {
		t.Run(s.name, func(t *testing.T) {
			ll := NewLinkedList[int]()
			s.setup(ll)

			// find last index of element in values
			expected := -1
			for i := len(s.values) - 1; i >= 0; i-- {
				if s.values[i] == s.value {
					expected = i
					break
				}
			}

			index, ok := ll.LastIndexOf(s.value)
			if expected != -1 && (!ok || index != expected) {
				t.Fatalf("Expected index to be %d, but found %d", expected, index)
			}
			if expected == -1 && ok {
				t.Fatalf("Expected element to not be found, but found %d", index)
			}
		})
	}
}

func TestLinkedList_Set(t *testing.T) {
	scenarios := []struct {
		linkedListScenario[int]
		insertIndex   int
		insertValue   int
		replacedValue *int
	}{
		{
			linkedListScenario: linkedListScenario[int]{
				name:     "set in empty list",
				values:   []int{},
				expected: nil,
			},
			insertIndex:   0,
			insertValue:   1,
			replacedValue: nil,
		},
		{
			linkedListScenario: linkedListScenario[int]{
				name:     "set in list with one element",
				values:   []int{1},
				expected: []int{2},
			},
			insertIndex:   0,
			insertValue:   2,
			replacedValue: intPtr(1),
		},
		{
			linkedListScenario: linkedListScenario[int]{
				name:     "set first element in list with multiple element",
				values:   []int{1, 2, 3, 4},
				expected: []int{5, 2, 3, 4},
			},
			insertIndex:   0,
			insertValue:   5,
			replacedValue: intPtr(1),
		},
		{
			linkedListScenario: linkedListScenario[int]{
				name:     "set middle element in list with multiple element",
				values:   []int{1, 2, 3, 4},
				expected: []int{1, 5, 3, 4},
			},
			insertIndex:   1,
			insertValue:   5,
			replacedValue: intPtr(2),
		},
		{
			linkedListScenario: linkedListScenario[int]{
				name:     "set last element in list with multiple element",
				values:   []int{1, 2, 3, 4},
				expected: []int{1, 2, 3, 5},
			},
			insertIndex:   3,
			insertValue:   5,
			replacedValue: intPtr(4),
		},
		{
			linkedListScenario: linkedListScenario[int]{
				name:     "set out of bounds in list with multiple element",
				values:   []int{1, 2, 3, 4},
				expected: []int{1, 2, 3, 4},
			},
			insertIndex:   4,
			insertValue:   5,
			replacedValue: nil,
		},
		{
			linkedListScenario: linkedListScenario[int]{
				name:     "set out of bounds in empty list",
				values:   []int{},
				expected: nil,
			},
			insertIndex:   1,
			insertValue:   5,
			replacedValue: nil,
		},
	}

	for _, s := range scenarios {
		t.Run(s.name, func(t *testing.T) {
			ll := NewLinkedList[int]()
			s.setup(ll)

			replaced, ok := ll.Set(s.insertIndex, s.insertValue)
			if replaced != nil && !ok {
				t.Fatalf("Expected replaced value to be %v, but found %v", s.replacedValue, replaced)
			}
			if ok && !reflect.DeepEqual(s.replacedValue, replaced) {
				t.Fatalf("Expected replaced value to be %v, but found %v", s.replacedValue, replaced)
			}
			s.test(t, ll)
		})
	}
}

func TestLinkedList_Clear(t *testing.T) {
	scenarios := []linkedListScenario[int]{
		{
			name:     "clear empty list",
			values:   []int{},
			expected: nil,
		},
		{
			name:     "clear list with one element",
			values:   []int{1},
			expected: nil,
		},
		{
			name:     "clear list with multiple element",
			values:   []int{1, 2, 3, 4},
			expected: nil,
		},
	}

	for _, s := range scenarios {
		t.Run(s.name, func(t *testing.T) {
			ll := NewLinkedList[int]()

			s.setup(ll)
			ll.Clear()
			s.test(t, ll)
		})
	}
}

func TestLinkedList_Contains(t *testing.T) {
	scenarios := []struct {
		linkedListScenario[int]
		value int
	}{
		{
			linkedListScenario: linkedListScenario[int]{
				name:   "contains element in empty list",
				values: []int{},
			},
			value: 1,
		},
		{
			linkedListScenario: linkedListScenario[int]{
				name:   "contains element in list with one element",
				values: []int{1},
			},
			value: 1,
		},
		{
			linkedListScenario: linkedListScenario[int]{
				name:   "contains element in list with multiple element",
				values: []int{1, 2, 3, 4},
			},
			value: 3,
		},
		{
			linkedListScenario: linkedListScenario[int]{
				name:   "contains element not in list with multiple element",
				values: []int{1, 2, 3, 4},
			},
			value: 5,
		},
	}

	for _, s := range scenarios {
		t.Run(s.name, func(t *testing.T) {
			ll := NewLinkedList[int]()

			s.setup(ll)

			contains := false
			for _, v := range s.values {
				if v == s.value {
					contains = true
					break
				}
			}

			if contains != ll.Contains(s.value) {
				t.Fatalf("Expected contains to be %v, but found %v", contains, ll.Contains(s.value))
			}
		})
	}
}

func TestLinkedList_String(t *testing.T) {
	scenarios := []struct {
		linkedListScenario[int]
		expected string
	}{
		{
			linkedListScenario: linkedListScenario[int]{
				name:   "string of empty list",
				values: []int{},
			},
			expected: "LinkedList([])",
		},
		{
			linkedListScenario: linkedListScenario[int]{
				name:   "string of list with one element",
				values: []int{1},
			},
			expected: "LinkedList([1])",
		},
		{
			linkedListScenario: linkedListScenario[int]{
				name:   "string of list with multiple element",
				values: []int{1, 2, 3, 4},
			},
			expected: "LinkedList([1 <-> 2 <-> 3 <-> 4])",
		},
	}

	for _, s := range scenarios {
		t.Run(s.name, func(t *testing.T) {
			ll := NewLinkedList[int]()

			s.setup(ll)
			if s.expected != ll.String() {
				t.Fatalf("Expected string to be %v, but found %v", s.expected, ll.String())
			}
		})
	}
}
