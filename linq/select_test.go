package linq

import (
	"fmt"
	"testing"

	"gogs.buffalo-robot.com/zouhy/linq/tools"
)

func TestSelect(t *testing.T) {
	type test struct {
		id   int
		name string
	}
	type args[T any, R any] struct {
		o Linq[T]
		f func(T) R
	}
	type testStruct[T any, R any] struct {
		name string
		args args[T, R]
		want []R
	}
	tests := []testStruct[test, int]{
		// TODO: Add test cases.
		{
			name: "nil",
			args: args[test, int]{
				o: FromSlice([]test{}),
				f: func(t test) int {
					return t.id
				},
			},
			want: []int{},
		},
		{
			name: "four",
			args: args[test, int]{
				o: FromSlice([]test{
					{
						id:   1,
						name: "杨婧祎",
					},
					{
						id:   2,
						name: "邹航宇",
					},
				}),
				f: func(t test) int {
					return t.id
				},
			},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Select(tt.args.o, tt.args.f)
			result := GetSlice(got)
			if !tools.SliceEqual(result, tt.want) {
				t.Errorf("Select() = %v,result=():%v,want %v", GetSlice(tt.args.o), result, tt.want)
			} else {
				fmt.Printf("Select() = %v,get %v", GetSlice(tt.args.o), result)
			}
		})
	}
}
