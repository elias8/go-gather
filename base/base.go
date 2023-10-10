package base

// Collection represents a generic collection of elements.
type Collection[T any] interface {
	// Contains returns true if the collection contains the specified element.
	Contains(element T) bool

	// Clear removes all elements from the collection.
	Clear()

	// IsEmpty returns true if the collection contains no elements.
	IsEmpty() bool

	// Size returns the number of elements in the collection.
	Size() int

	// Values returns a slice representation of the collection.
	Values() []T

	// String returns string representation of the collection.
	String() string
}
