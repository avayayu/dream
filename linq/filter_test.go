package linq

import (
	"fmt"
	"testing"

	"gogs.buffalo-robot.com/zouhy/linq/tools"
)

func TestFilter(t *testing.T) {
	type args[T any] struct {
		o Linq[T]
		f func(T) bool
	}
	type testStruct[T any] struct {
		name string
		args args[T]
		want []T
	}
	// predictFuc := func(d string) bool {
	// 	return strings.HasPrefix(d, "a")
	// }
	tests3 := []testStruct[int]{
		{
			name: "nil",
			args: args[int]{
				o: FromSlice([]int{}),
				f: func(d int) bool {
					return d > 1
				},
			},
			want: []int{},
		},
		{
			name: "serveral",
			args: args[int]{
				o: FromSlice([]int{1, 2, 3, 4, 5}),
				f: func(d int) bool {
					return d > 3
				},
			},
			want: []int{4, 5},
		},
	}

	for _, tt := range tests3 {
		t.Run(tt.name, func(t *testing.T) {
			got := Filter(tt.args.o, tt.args.f)
			result := GetSlice(got)
			if !tools.SliceEqual(result, tt.want) {
				t.Errorf("Select() = %v,result=():%v,want %v", GetSlice(tt.args.o), result, tt.want)
			} else {
				fmt.Printf("Select() = %v,get %v", GetSlice(tt.args.o), result)
			}

		})
	}
}

// type A struct {
// 	name string
// }
// type complexStruct struct {
// 	name string
// 	id   int
// 	A    *A
// }

// type testStruct2[T any] struct {
// 	name string
// 	args args[T]
// 	want []string
// }

// tests2 := []testStruct2[A]{
// 	{
// 		name: "complex test",
// 		args: args[A]{
// 			o: FromSlice([]A{
// 				{
// 					name: "邹航宇",
// 				},
// 				{
// 					name: "杨婧祎",
// 				},
// 			}),
// 			f: func(t A) bool {
// 				if strings.HasPrefix(t.name, "邹") {
// 					return true
// 				}
// 				return false
// 			},
// 		},
// 		want: []string{"邹航宇"},
// 	},
// }

// for _, tt := range tests2 {
// 	t.Run(tt.name, func(t *testing.T) {
// 		got := Filter(tt.args.o, tt.args.f)
// 		got2 := Select(got, func(t A) string {
// 			return t.name
// 		})
// 		result := GetSlice(got2)
// 		if !tools.SliceEqual(result, tt.want) {
// 			t.Errorf("Select() = %v,result=():%v,want %v", GetSlice(tt.args.o), result, tt.want)
// 		} else {
// 			fmt.Printf("Select() = %v,get %v", GetSlice(tt.args.o), result)
// 		}

// 	})
// }
//}
