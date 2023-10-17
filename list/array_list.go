package list

import (
	"fmt"
	"reflect"
)

type arrayList[T any] struct {
	elements []T
}

func NewArrayList[T any]() List[T] {
	return &arrayList[T]{}
}

func (a *arrayList[T]) Size() int {
	return len(a.elements)
}

func (a *arrayList[T]) IsEmpty() bool {
	return len(a.elements) == 0
}

func (a *arrayList[T]) Contains(element T) bool {
	for _, e := range a.elements {
		if reflect.DeepEqual(e, element) {
			return true
		}
	}
	return false
}

func (a *arrayList[T]) String() string {
	s := "ArrayList(["
	for i, e := range a.elements {
		if i > 0 {
			s += ", "
		}
		s += fmt.Sprintf("%v", e)
	}
	s += "])"
	return s
}

func (a *arrayList[T]) Values() []T {
	return append([]T(nil), a.elements...)
}

func (a *arrayList[T]) Add(element T) {
	a.elements = append(a.elements, element)
}

func (a *arrayList[T]) Clear() {
	a.elements = nil
}

func (a *arrayList[T]) Remove(element T) bool {
	for i, e := range a.elements {
		if reflect.DeepEqual(e, element) {
			a.elements = append(a.elements[:i], a.elements[i+1:]...)
			return true
		}
	}
	return false
}

func (a *arrayList[T]) Set(index int, element T) (*T, bool) {
	if index < 0 || index >= len(a.elements) {
		return nil, false
	}
	previous := a.elements[index]
	a.elements[index] = element
	return &previous, true
}

func (a *arrayList[T]) Get(index int) (*T, bool) {
	if index < 0 || index >= len(a.elements) {
		return nil, false
	}
	return &a.elements[index], true
}

func (a *arrayList[T]) IndexOf(element T) (int, bool) {
	for i, e := range a.elements {
		if reflect.DeepEqual(e, element) {
			return i, true
		}
	}
	return -1, false
}

func (a *arrayList[T]) LastIndexOf(element T) (int, bool) {
	for i := len(a.elements) - 1; i >= 0; i-- {
		if reflect.DeepEqual(a.elements[i], element) {
			return i, true
		}
	}
	return -1, false
}
