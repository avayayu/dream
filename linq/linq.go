package linq

import (
	"constraints"
)

type Enumerator[T any] interface {
	MoveNext() bool
	Current() *T
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

func Zero[T any]() *T {
	var d T
	return &d
}

func FromSlice[T any](data []T) Linq[T] {
	slice := enumerableSlice[T]{
		data:  data,
		index: -1,
	}
	return &sliceLINQ[T]{slice: &slice}
}

func (slice *sliceLINQ[T]) GetEnumerator() Enumerator[T] {
	return &enumerableSlice[T]{
		data:  slice.slice.data,
		index: -1,
	}
}

func (slice *enumerableSlice[T]) MoveNext() bool {
	slice.index++
	if slice.index < len(slice.data) {
		return true
	}
	return false
}

func (slice *enumerableSlice[T]) Current() *T {
	return &slice.data[slice.index]
}

//GetSlice 从slice中还原出slice
func GetSlice[T any](linq Linq[T]) []T {
	data := []T{}
	if IsEmpty(linq) {
		return data
	}
	iterator := linq.GetEnumerator()
	for {
		if !iterator.MoveNext() {
			break
		}
		current := iterator.Current()
		if current != nil {
			data = append(data, *current)
		}

	}
	return data
}

//IsEmpty 判断linq[T]是否为空
func IsEmpty[T any](e Linq[T]) bool {
	ito := e.GetEnumerator()
	if !ito.MoveNext() {
		return true
	}
	return false
}

func Max[T constraints.Ordered](e Linq[T]) T {
	var max T
	if IsEmpty(e) {
		return max
	}
	iterator := e.GetEnumerator()
	for {
		if !iterator.MoveNext() {
			break
		}
		current := iterator.Current()
		if current != nil {
			data := *current
			if max < data {
				max = data
			}
		}
	}
	return max
}

func Min[T constraints.Ordered](e Linq[T]) T {
	var min T
	if IsEmpty(e) {
		return min
	}
	iterator := e.GetEnumerator()
	for {
		if !iterator.MoveNext() {
			break
		}
		current := iterator.Current()
		if current != nil {
			data := *current
			if min > data {
				min = data
			}
		}
	}
	return min
}

//GroupBy 根据groupByFunc中的R为键 将序列划分为 key->array的形式
func GroupBy[T any, R comparable](e Linq[T], groupByFunc func(T) R) map[R][]T {
	groupData := map[R][]T{}
	if IsEmpty(e) {
		return groupData
	}
	iterator := e.GetEnumerator()
	for {
		if !iterator.MoveNext() {
			break
		}
		current := iterator.Current()
		key := groupByFunc(*current)
		_, ok := groupData[key]
		if !ok {
			groupData[key] = make([]T, 10)
		}
		groupData[key] = append(groupData[key], *current)
	}
	return groupData
}

//GroupBy 根据groupByFunc中的R为键 将序列划分为 key->T
func GroupByDistinct[T any, R comparable](e Linq[T], groupByFunc func(T) R) map[R]T {
	groupData := map[R]T{}
	if IsEmpty(e) {
		return groupData
	}
	iterator := e.GetEnumerator()
	for {
		if !iterator.MoveNext() {
			break
		}
		current := iterator.Current()
		if current != nil {
			key := groupByFunc(*current)
			groupData[key] = *current
		}
	}
	return groupData
}

func Count[T any](e Linq[T]) int {
	count := 0
	if IsEmpty(e) {
		return 0
	}
	iterator := e.GetEnumerator()
	for {
		if !iterator.MoveNext() {
			break
		}
		current := iterator.Current()
		if current != nil {
			count++
		}
	}
	return count

}
