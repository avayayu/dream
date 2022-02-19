package stream

type ICollection[T any] interface {
	IIterable[T]

	// Index returns the value in the position indicated by the index.
	//
	//   - index:  The index of the element to be retrieved.
	//
	Index(index int) T

	// Add appends the element into the iterable. Returns error if the item is not the proper type
	//
	//   - item:  The item to be added to the collection.
	//
	Add(item T) error

	// AddAll appends another iterable into this ICollection instance.
	//
	//   - iterable:  The iterable of elements to be added to the collection.
	//
	AddAll(iterable IIterable[T]) error

	// Remove removes the element at the provided index. Returns the removed item or `nil` if no item was found in that position.
	//
	//   - index:      The index of the element to remove
	//   - keepOrder:  (false by default) Optional flag that indicates if the removal of the element should guarantee the order of the remaining elements. In some cases,
	//                 guaranteeing the order of elements after a removal can me a costly operation since the remaining elements have to be shifted in the collection.
	//
	Remove(index int, keepOrder ...bool) T
}

// IMapCollection represents a collection of `*KeyValuePairs` tied to a `map`
type IMapCollection[T any, R comparable] interface {
	ICollection[T]

	// ToMap returns a map representation of the IMapIterable
	ToMap() map[R]T

	// Get returns the value at index `key`, or the value mapped to the key `key` if the collection represents a `map`.
	Get(key R) T

	// Set is mapCollection specific function that allows a value to be added to the map without having to wrap it in a *KeyValuePair
	Set(key, value interface{}) error
}
