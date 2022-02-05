package linq

type Enumerator[T any] interface {
	MoveNext() bool
	Current() T
}

type Linq[T any] interface {
	GetEnumerator() Enumerator[T]
}

type sliceLINQ[T any] struct {
	slice *enumerableSlice[T]
}

type enumerableSlice[T any] struct {
	data  []T
	index int
}

func Zero[T any]() T {
	var d T
	return d
}

func FromSlice[T any](data []T) Linq[T] {
	slice := enumerableSlice[T]{
		data:  data,
		index: 0,
	}
	return &sliceLINQ[T]{slice: &slice}
}

func (slice *sliceLINQ[T]) GetEnumerator() Enumerator[T] {
	return &enumerableSlice[T]{
		data:  slice.slice.data,
		index: 0,
	}
}

func (slice *enumerableSlice[T]) MoveNext() bool {
	if slice.index < len(slice.data) {
		slice.index++
		return true
	}
	return false
}

func (slice *enumerableSlice[T]) Current() T {
	return slice.data[slice.index]
}

func GetSlice[T any](linq Linq[T]) []T {
	data := []T{}
	iterator := linq.GetEnumerator()
	for {
		current := iterator.Current()
		data = append(data, current)
		if !iterator.MoveNext() {
			break
		}
	}
	return data
}
