package linq

func Where[T any](data IEnumerable[T], predicate func(d T) bool) IEnumerable[T] {

	for data.MoveNext() {

		if predicate(data.Current()) {

		}

	}
}
