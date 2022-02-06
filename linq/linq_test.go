package linq

import (
	"reflect"
	"testing"
)

type TT interface {
	int | string
}

func TestFromSlice(t *testing.T) {
	type args[T any] struct {
		data []T
	}
	type testStruct[T any] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testStruct[string]{
		{
			name: "nil",
			args: args[string]{
				data: []string{},
			},
			want: []string{},
		},
		{
			name: "two",
			args: args[string]{
				data: []string{"1", "2", "3", "4"},
			},
			want: []string{"1", "2", "3", "4"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FromSlice(tt.args.data)
			result := GetSlice(got)
			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf("FromSlice() = %v, GetSlice() = %v  want %v", got, result, tt.want)
			} else {
				t.Logf("FromSlice() = %v, GetSlice() = %v", got, result)
			}
		})
	}
}
