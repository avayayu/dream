package stream

type IIterator[T any] interface {
	Current() T
	MoveNext() bool
	HasNext() bool
	Next() T
	Skip(n int) IIterator[T]
}
