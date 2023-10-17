package list

import (
	"reflect"
	"testing"
)

type arrayListScenario[T any] struct {
	name     string
	value    []T
	expected []T
}

func (s arrayListScenario[T]) test(list List[T], t *testing.T) {
	al := list.(*arrayList[T])
	s.testSize(al, t)
	s.testListOrder(al, t)
}

func (s arrayListScenario[T]) testSize(list *arrayList[T], t *testing.T) {
	if list.Size() != len(s.expected) {
		t.Fatalf("Expected size %d, got %d", len(s.expected), list.Size())
	}
}

func (s arrayListScenario[T]) testListOrder(list *arrayList[T], t *testing.T) {
	for i := 0; i < len(s.expected); i++ {
		if !reflect.DeepEqual(list.elements[i], s.expected[i]) {
			t.Fatalf("Expected element at index %d to be %v, but got %v", i, s.expected[i], list.elements[i])
		}
	}
}

func TestNewArrayList(t *testing.T) {
	list := NewArrayList[int]()
	if list == nil {
		t.Fatalf("Expected NewArrayList() to return a List, got nil")
	}
	if list.Size() != 0 {
		t.Fatalf("NewArrayList() should return an empty List, got size %d", list.Size())
	}
}

func TestArrayListAdd(t *testing.T) {
	scenarios := []arrayListScenario[int]{
		{name: "Add 1 element", value: []int{1}, expected: []int{1}},
		{name: "Add 2 elements", value: []int{1, 2}, expected: []int{1, 2}},
		{name: "Add 3 elements", value: []int{1, 2, 3}, expected: []int{1, 2, 3}},
	}
	for _, s := range scenarios {
		t.Run(s.name, func(t *testing.T) {
			list := NewArrayList[int]()
			for _, e := range s.value {
				list.Add(e)
			}
			s.test(list, t)
		})
	}
}

func TestArrayListClear(t *testing.T) {
	scenarios := []arrayListScenario[int]{
		{name: "Clear empty list", value: []int{}},
		{name: "Clear 1 element", value: []int{1}},
		{name: "Clear 2 elements", value: []int{1, 2}},
		{name: "Clear 3 elements", value: []int{1, 2, 3}},
	}
	for _, s := range scenarios {
		t.Run(s.name, func(t *testing.T) {
			list := NewArrayList[int]()
			for _, e := range s.value {
				list.Add(e)
			}
			list.Clear()
			s.test(list, t)
		})
	}
}

func TestArrayListContains(t *testing.T) {
	scenarios := []struct {
		arrayListScenario[int]
		element    int
		shouldFind bool
	}{
		{arrayListScenario: arrayListScenario[int]{
			name:     "Empty list does not contain element",
			value:    []int{},
			expected: []int{1}},
			element:    1,
			shouldFind: false,
		},
		{arrayListScenario: arrayListScenario[int]{
			name:     "Contains 1 element",
			value:    []int{1},
			expected: []int{1}},
			element:    1,
			shouldFind: true,
		},
		{arrayListScenario: arrayListScenario[int]{
			name:     "Contains 2 elements",
			value:    []int{1, 2},
			expected: []int{1}},
			element:    2,
			shouldFind: true,
		},
		{arrayListScenario: arrayListScenario[int]{
			name:     "Contains 3 elements",
			value:    []int{1, 2, 3},
			expected: []int{1}},
			element:    2,
			shouldFind: true,
		},
	}
	for _, s := range scenarios {
		t.Run(s.name, func(t *testing.T) {
			list := NewArrayList[int]()
			for _, e := range s.value {
				list.Add(e)
			}

			found := list.Contains(s.element)
			if s.shouldFind && !found {
				t.Fatalf("Expected list to contain %d", s.element)
			}
			if !s.shouldFind && found {
				t.Fatalf("Expected list to not contain %d", s.element)
			}
		})
	}
}

func TestArrayList_Values(t *testing.T) {
	scenarios := []arrayListScenario[int]{
		{
			name:     "empty list",
			value:    []int{},
			expected: nil,
		},
		{
			name:     "1 element list",
			value:    []int{1},
			expected: []int{1},
		},
		{
			name:     "2 element list",
			value:    []int{1, 2},
			expected: []int{1, 2},
		},
		{
			name:     "multiple element list",
			value:    []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
	}

	for _, s := range scenarios {
		t.Run(s.name, func(t *testing.T) {
			list := NewArrayList[int]()
			for _, e := range s.value {
				list.Add(e)
			}

			values := list.Values()
			if !reflect.DeepEqual(values, s.expected) {
				t.Fatalf("Expected values %v, got %v", s.expected, values)
			}
			s.test(list, t)
		})
	}
}

func TestArrayListRemove(t *testing.T) {
	scenarios := []struct {
		arrayListScenario[int]
		shouldBeRemoved bool
		element         int
	}{
		{
			arrayListScenario: arrayListScenario[int]{
				name:     "Remove from empty list",
				value:    []int{},
				expected: []int{},
			},
			element:         1,
			shouldBeRemoved: false,
		},
		{
			arrayListScenario: arrayListScenario[int]{
				name:     "Remove from 1 element list",
				value:    []int{1},
				expected: []int{},
			},
			element:         1,
			shouldBeRemoved: true,
		},
		{
			arrayListScenario: arrayListScenario[int]{
				name:     "Remove from 2 element list",
				value:    []int{1, 2},
				expected: []int{1},
			},
			element:         2,
			shouldBeRemoved: true,
		},
		{
			arrayListScenario: arrayListScenario[int]{
				name:     "Remove from multiple element list",
				value:    []int{1, 2, 3},
				expected: []int{1, 3},
			},
			element:         2,
			shouldBeRemoved: true,
		},
	}

	for _, s := range scenarios {
		t.Run(s.name, func(t *testing.T) {
			list := NewArrayList[int]()
			for _, e := range s.value {
				list.Add(e)
			}
			removed := list.Remove(s.element)
			if !s.shouldBeRemoved && removed {
				t.Fatalf("Expected element %d to not be removed from list", s.element)
			}

			if s.shouldBeRemoved && !removed {
				t.Fatalf("Expected element %d to be removed from list", s.element)
			}
			s.test(list, t)
		})
	}
}

func TestArrayListSet(t *testing.T) {
	scenarios := []struct {
		arrayListScenario[int]
		index         int
		element       int
		shouldSucceed bool
	}{
		{
			arrayListScenario: arrayListScenario[int]{
				name:     "Set on empty list",
				value:    []int{},
				expected: []int{},
			},
			index:         0,
			element:       1,
			shouldSucceed: false,
		},
		{
			arrayListScenario: arrayListScenario[int]{
				name:     "Set on 1 element list",
				value:    []int{1},
				expected: []int{2},
			},
			index:         0,
			element:       2,
			shouldSucceed: true,
		},
		{
			arrayListScenario: arrayListScenario[int]{
				name:     "Set on 2 element list",
				value:    []int{1, 2},
				expected: []int{1, 3},
			},
			index:         1,
			element:       3,
			shouldSucceed: true,
		},
		{
			arrayListScenario: arrayListScenario[int]{
				name:     "Set on multiple element list",
				value:    []int{1, 2, 3},
				expected: []int{1, 4, 3},
			},
			index:         1,
			element:       4,
			shouldSucceed: true,
		},
		{
			arrayListScenario: arrayListScenario[int]{
				name:     "Set on multiple element list with invalid index",
				value:    []int{1, 2, 3},
				expected: []int{1, 2, 3},
			},
			index:         3,
			element:       4,
			shouldSucceed: false,
		},
	}

	for _, s := range scenarios {
		t.Run(s.name, func(t *testing.T) {
			list := NewArrayList[int]()
			for _, e := range s.value {
				list.Add(e)
			}
			previous, success := list.Set(s.index, s.element)
			if s.shouldSucceed {
				if !success {
					t.Fatalf("Expected to set element at index %d", s.index)
				}

				if *previous != s.value[s.index] {
					t.Fatalf("Expected previous value to be %v but found %v", s.value[s.index], *previous)
				}
			}

			if !s.shouldSucceed && success {
				t.Fatalf("Expected to not set element at index %d", s.index)
			}

			s.test(list, t)
		})
	}
}

func TestArrayListGet(t *testing.T) {
	scenarios := []struct {
		arrayListScenario[int]
		index      int
		shouldFind bool
	}{
		{
			arrayListScenario: arrayListScenario[int]{
				name:     "Get on empty list",
				value:    []int{},
				expected: []int{},
			},
			index:      0,
			shouldFind: false,
		},
		{
			arrayListScenario: arrayListScenario[int]{
				name:     "Get on 1 element list",
				value:    []int{1},
				expected: []int{1},
			},
			index:      0,
			shouldFind: true,
		},
		{
			arrayListScenario: arrayListScenario[int]{
				name:     "Get on 2 element list",
				value:    []int{1, 2},
				expected: []int{1, 2},
			},
			index:      1,
			shouldFind: true,
		},
		{
			arrayListScenario: arrayListScenario[int]{
				name:     "Get on multiple element list",
				value:    []int{1, 2, 3},
				expected: []int{1, 2, 3},
			},
			index:      1,
			shouldFind: true,
		},
		{
			arrayListScenario: arrayListScenario[int]{
				name:     "Get on multiple element list with invalid index",
				value:    []int{1, 2, 3},
				expected: []int{1, 2, 3},
			},
			index:      3,
			shouldFind: false,
		},
	}

	for _, s := range scenarios {
		t.Run(s.name, func(t *testing.T) {
			list := NewArrayList[int]()
			for _, e := range s.value {
				list.Add(e)
			}
			element, found := list.Get(s.index)
			if s.shouldFind && !found {
				t.Fatalf("Expected to find element at index %d", s.index)
			}

			if !s.shouldFind && found {
				t.Fatalf("Expected to not find element at index %d", s.index)
			}

			if found && *element != s.value[s.index] {
				t.Fatalf("Expected element at index %d to be %v, but got %v", s.index, s.value[s.index], *element)
			}
		})
	}
}

func TestArrayList_IndexOf(t *testing.T) {
	scenarios := []struct {
		arrayListScenario[int]
		element    int
		index      int
		shouldFind bool
	}{
		{
			arrayListScenario: arrayListScenario[int]{
				name:     "IndexOf on empty list",
				value:    []int{},
				expected: []int{},
			},
			element:    1,
			index:      -1,
			shouldFind: false,
		},
		{
			arrayListScenario: arrayListScenario[int]{
				name:     "IndexOf on 1 element list",
				value:    []int{1},
				expected: []int{1},
			},
			element:    1,
			index:      0,
			shouldFind: true,
		},
		{
			arrayListScenario: arrayListScenario[int]{
				name:     "IndexOf on 2 element list",
				value:    []int{1, 2},
				expected: []int{1, 2},
			},
			element:    2,
			index:      1,
			shouldFind: true,
		},
		{
			arrayListScenario: arrayListScenario[int]{
				name:     "IndexOf on multiple element list",
				value:    []int{1, 2, 3},
				expected: []int{1, 2, 3},
			},
			element:    2,
			index:      1,
			shouldFind: true,
		},
		{
			arrayListScenario: arrayListScenario[int]{
				name:     "IndexOf on multiple element list with duplicate values",
				value:    []int{1, 2, 2, 3},
				expected: []int{1, 2, 2, 3},
			},
			element:    2,
			index:      1,
			shouldFind: true,
		},
		{
			arrayListScenario: arrayListScenario[int]{
				name:     "IndexOf on multiple element list with invalid index",
				value:    []int{1, 2, 3},
				expected: []int{1, 2, 3},
			},
			element:    4,
			index:      -1,
			shouldFind: false,
		},
	}

	for _, s := range scenarios {
		t.Run(s.name, func(t *testing.T) {
			list := NewArrayList[int]()
			for _, e := range s.value {
				list.Add(e)
			}
			index, found := list.IndexOf(s.element)
			if s.shouldFind && !found {
				t.Fatalf("Expected to find element %d", s.element)
			}

			if !s.shouldFind && found {
				t.Fatalf("Expected to not find element %d", s.element)
			}

			if found && index != s.index {
				t.Fatalf("Expected element %d to be at index %d, but got %d", s.element, s.index, index)
			}
		})
	}
}

func TestArrayList_LastIndexOf(t *testing.T) {
	scenarios := []struct {
		arrayListScenario[int]
		element    int
		index      int
		shouldFind bool
	}{
		{
			arrayListScenario: arrayListScenario[int]{
				name:     "LastIndexOf on empty list",
				value:    []int{},
				expected: []int{},
			},
			element:    1,
			index:      -1,
			shouldFind: false,
		},
		{
			arrayListScenario: arrayListScenario[int]{
				name:     "LastIndexOf on 1 element list",
				value:    []int{1},
				expected: []int{1},
			},
			element:    1,
			index:      0,
			shouldFind: true,
		},
		{
			arrayListScenario: arrayListScenario[int]{
				name:     "LastIndexOf on 2 element list",
				value:    []int{1, 2},
				expected: []int{1, 2},
			},
			element:    2,
			index:      1,
			shouldFind: true,
		},
		{
			arrayListScenario: arrayListScenario[int]{
				name:     "LastIndexOf on multiple element list",
				value:    []int{1, 2, 3},
				expected: []int{1, 2, 3},
			},
			element:    2,
			index:      1,
			shouldFind: true,
		},
		{
			arrayListScenario: arrayListScenario[int]{
				name:     "LastIndexOf on multiple element list with duplicate values",
				value:    []int{1, 2, 2, 3},
				expected: []int{1, 2, 2, 3},
			},
			element:    2,
			index:      2,
			shouldFind: true,
		},
		{
			arrayListScenario: arrayListScenario[int]{
				name:     "LastIndexOf on multiple element list with invalid index",
				value:    []int{1, 2, 3},
				expected: []int{1, 2, 3},
			},
			element:    4,
			index:      -1,
			shouldFind: false,
		},
	}

	for _, s := range scenarios {
		t.Run(s.name, func(t *testing.T) {
			list := NewArrayList[int]()
			for _, e := range s.value {
				list.Add(e)
			}
			index, found := list.LastIndexOf(s.element)
			if s.shouldFind && !found {
				t.Fatalf("Expected to find element %d", s.element)
			}

			if !s.shouldFind && found {
				t.Fatalf("Expected to not find element %d", s.element)
			}

			if found && index != s.index {
				t.Fatalf("Expected element %d to be at index %d, but got %d", s.element, s.index, index)
			}
		})
	}
}

func TestArrayList_IsEmpty(t *testing.T) {
	scenarios := []arrayListScenario[int]{
		{
			name:     "Empty list",
			value:    []int{},
			expected: []int{},
		},
		{
			name:     "1 element list",
			value:    []int{1},
			expected: []int{1},
		},
		{
			name:     "2 element list",
			value:    []int{1, 2},
			expected: []int{1, 2},
		},
	}

	for _, s := range scenarios {
		t.Run(s.name, func(t *testing.T) {
			list := NewArrayList[int]()
			for _, e := range s.value {
				list.Add(e)
			}
			if list.IsEmpty() != (len(s.value) == 0) {
				t.Fatalf("Expected list to be empty: %v", len(s.value) == 0)
			}
		})
	}
}

func TestArrayList_String(t *testing.T) {
	scenarios := []struct {
		arrayListScenario[int]
		expected string
	}{
		{
			arrayListScenario: arrayListScenario[int]{
				name:  "Empty list",
				value: []int{},
			},
			expected: "ArrayList([])",
		},
		{
			arrayListScenario: arrayListScenario[int]{
				name:  "1 element list",
				value: []int{1},
			},
			expected: "ArrayList([1])",
		},
		{
			arrayListScenario: arrayListScenario[int]{
				name:  "2 element list",
				value: []int{1, 2},
			},
			expected: "ArrayList([1, 2])",
		},
		{
			arrayListScenario: arrayListScenario[int]{
				name:  "Multiple element list",
				value: []int{1, 2, 3},
			},
			expected: "ArrayList([1, 2, 3])",
		},
	}

	for _, s := range scenarios {
		t.Run(s.name, func(t *testing.T) {
			list := NewArrayList[int]()
			for _, e := range s.value {
				list.Add(e)
			}
			if list.String() != s.expected {
				t.Fatalf("Expected list to be %s, got %s", s.expected, list.String())
			}
		})
	}
}
