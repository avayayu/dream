package linq

type IEnumerable[T any] interface {
	MoveNext() bool
	Current() T
	Reset()
}

type enumerableSlice[T any] struct {
	data  []T
	index int
}

func NewLinQS[T any](slice []T) IEnumerable[T] {
	e := enumerableSlice[T]{
		data:  []T{},
		index: 0,
	}
	return &e
}

func (e *enumerableSlice[T]) MoveNext() bool {
	return false
}

func (e *enumerableSlice[T]) Current() T {
	return e.data[e.index]
}

func (e *enumerableSlice[T]) Reset() {
	e.index = 0
	return
}
