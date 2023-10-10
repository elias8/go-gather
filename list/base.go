package list

import (
	"github.com/elias8/go-gather/base"
)

// List represents a list of elements.
type List[T any] interface {
	base.Collection[T]

	// Add adds the specified element to the list.
	Add(element T)

	// Remove removes the first occurrence of the specified element from the
	// list. Returns true if the element is removed, false otherwise.
	Remove(element T) bool

	// Set replaces the element at the specified position in the list with the
	// specified element.
	Set(index int, element T) (*T, bool)

	// Get returns the element at the specified position in the list. If the
	// index is out of range (index < 0 || index >= Size()), returns nil and
	// false.
	Get(index int) (*T, bool)

	// IndexOf returns the index of the first occurrence of the specified
	// element in the list. If the list does not contain the element, returns
	// -1 and false.
	IndexOf(element T) (int, bool)

	// LastIndexOf returns the index of the last occurrence of the specified
	// element in the list. If the list does not contain the element, returns
	// -1 and false.
	LastIndexOf(element T) (int, bool)
}

// LinkedList is a doubly-linked
type LinkedList[T any] interface {
	List[T]

	// Add appends the specified element to the end of the list.
	//
	// The operation is performed in O(1) time.
	Add(element T)

	// AddFirst inserts the specified element at the beginning of the list.
	//
	// The operation is performed in O(1) time.
	AddFirst(element T)

	// AddLast appends the specified element to the end of the list (equivalent
	// to Add).
	//
	// The operation is performed in O(1) time.
	AddLast(element T)

	// Reverse reverses the order of the elements in the list. The head becomes
	// the tail and vice versa.
	//
	// The operation is performed in O(n) time.
	Reverse()

	// Remove removes the first occurrence of the specified element from the
	// list. Returns true if the element is removed, false otherwise.
	//
	// The operation is performed in O(n) time in the worst case. Removing head
	// or tail is performed in O(1) time.
	Remove(element T) bool

	// RemoveFirst removes and returns the first element from the list. Returns
	// the removed element and true if the list is not empty, nil and false
	// otherwise.
	//
	// The operation is performed in O(1) time.
	RemoveFirst() (*T, bool)

	// RemoveLast removes and returns the last element from the list. Returns
	// the removed element and true if the list is not empty, nil and false
	// otherwise.
	//
	// The operation is performed in O(1) time.
	RemoveLast() (*T, bool)

	// GetFirst returns the first element in the list. Returns the first element
	// and true if the list is not empty, nil and false otherwise.
	//
	// The operation is performed in O(1) time.
	GetFirst() (*T, bool)

	// GetLast returns the last element in the list. Returns the last element
	// and true if the list is not empty, nil and false otherwise.
	//
	// The operation is performed in O(1) time.
	GetLast() (*T, bool)

	// Get returns the element at the specified position in the list. If the
	// index is out of range (index < 0 || index >= Size()), returns nil and
	// false.
	//
	// The operation is performed in O(n) time in the worst case.
	Get(index int) (*T, bool)

	// IndexOf returns the index of the first occurrence of the specified
	// element in the list. If the list does not contain the element, returns
	// -1 and false.
	//
	// The operation is performed in O(n) time in the worst case.
	IndexOf(element T) (int, bool)

	// LastIndexOf returns the index of the last occurrence of the specified
	// element in the list. If the list does not contain the element, returns
	// -1 and false.
	//
	// The operation is performed in O(n) time in the worst case.
	LastIndexOf(element T) (int, bool)
}
