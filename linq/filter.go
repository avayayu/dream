package linq

type filterEnumerable[T any] struct {
	Enumerator Linq[T]
	filterFunc func(T) bool
}

type filterEnumerator[T any] struct {
	Enumerator Enumerator[T]
	filterFunc func(T) bool
	current    *T
}

func (slice *filterEnumerable[T]) GetEnumerator() Enumerator[T] {
	return &filterEnumerator[T]{slice.Enumerator.GetEnumerator(), slice.filterFunc, nil}
}

func Filter[T any](o Linq[T], f func(T) bool) Linq[T] {
	return &filterEnumerable[T]{
		Enumerator: o,
		filterFunc: f,
	}
}

func (slice *filterEnumerator[T]) MoveNext() bool {
	if !slice.Enumerator.MoveNext() {
		return false
	}
	currest := slice.Enumerator.Current()
	if slice.filterFunc(*currest) {
		slice.current = currest
	} else {
		slice.current = nil
	}
	return true
}

func (slice *filterEnumerator[T]) Current() *T {
	return slice.current
}
