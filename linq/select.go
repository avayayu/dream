package linq

type selectEnumerable[T any, R any] struct {
	Enumerator Linq[T]
	selectFunc func(T) R
}

type selectEnumerator[T any, R any] struct {
	Enumerator Enumerator[T]
	selectFunc func(T) R
	current    *R
}

func (slice *selectEnumerable[T, R]) GetEnumerator() Enumerator[R] {
	return &selectEnumerator[T, R]{slice.Enumerator.GetEnumerator(), slice.selectFunc, nil}
}

func Select[T any, R any](o Linq[T], f func(T) R) Linq[R] {
	return &selectEnumerable[T, R]{
		Enumerator: o,
		selectFunc: f,
	}
}

func (slice *selectEnumerator[T, R]) MoveNext() bool {
	if !slice.Enumerator.MoveNext() {
		return false
	}
	currest := slice.Enumerator.Current()
	if currest != nil {
		cur := slice.selectFunc(*currest)
		slice.current = &cur
	}
	return true
}

func (slice *selectEnumerator[T, R]) Current() *R {
	return slice.current
}
