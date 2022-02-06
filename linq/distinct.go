package linq

type distinctEnumerable[T any, R comparable] struct {
	Enumerator   Linq[T]
	distinctFunc func(T) R
}

type distinctEnumerator[T any, R comparable] struct {
	Enumerator   Enumerator[T]
	distinctFunc func(T) R
	current      *T
	distinctMap  map[R]struct{}
}

func (slice *distinctEnumerable[T, R]) GetEnumerator() Enumerator[T] {
	return &distinctEnumerator[T, R]{slice.Enumerator.GetEnumerator(), slice.distinctFunc, nil, map[R]struct{}{}}
}

func (slice *distinctEnumerator[T, R]) MoveNext() bool {
	if !slice.Enumerator.MoveNext() {
		return false
	}
	currest := slice.Enumerator.Current()
	key := slice.distinctFunc(*currest)
	_, ok := slice.distinctMap[key]
	if !ok {
		slice.distinctMap[key] = struct{}{}
	} else {
		slice.current = nil
	}
	return true
}

func (slice *distinctEnumerator[T, R]) Current() *T {
	return slice.current
}
