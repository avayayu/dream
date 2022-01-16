package linq

type IEnumerable[T any] interface {
	MoveNext() bool
	Current() T
	Reset()
}

type Collection[K comparable, T any] interface {
	[]T | map[K]T
}

// type enumerable [T []T |] struct {

// }
