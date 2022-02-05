package linq

type selectEnumerable[T any, R any] struct {
	Enumerator Linq[T]
	selectFunc func(T) R
}

type selectEnumerator[T any, R any] struct {
	Enumerator Linq[T]
	selectFunc func(T) R
	current    R
}

func (slice *selectEnumerable[T, R]) GetEnumerator() Enumerator[R] {
	return &selectEnumerator[T, R]{slice.Enumerator, slice.selectFunc, Zero[R]()}
}

func Select[T any, R any](o Linq[T], f func(T) R) Linq[R] {
	return &selectEnumerable[T, R]{
		Enumerator: o,
		selectFunc: f,
	}
}

func (slice *selectEnumerator[T, R]) MoveNext() bool {
	return true
}

func (slice *selectEnumerator[T, R]) Current() R {
	return slice.current
}
