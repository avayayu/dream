package stream

type IIterable[T any] interface {
	// Returns a new collectionIterator for the iterable
	Iterator() IIterator[T]

	// Len returns the size of the iterable if the size is finite and known, otherwise returns -1.
	Len() int

	// ToArray returns an array representation of the iterable
	ToArray(defaultArray ...T) []T
}

type IMapIterable[T any, R comparable] interface {
	IIterable[T]

	// ToMap returns a map representation of the IMapIterable
	ToMap() map[R]T
}
