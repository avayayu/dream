package linq

import (
	"math/rand"
	"reflect"
	"testing"

	"gogs.buffalo-robot.com/zouhy/linq/tools"
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




func randomString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(61 + rand.Intn(25))
	}
	return string(bytes)
}

var test = []string{
	"aaa","bbb","ccc",
}
var test2 = []string{
	"aaa","bbb","ccc","aaa","bbb","ccc","aaa","bbb","ccc","aaa","bbb","ccc","aaa","bbb","ccc","aaa","bbb","ccc","aaa","bbb","ccc","aaa","bbb","ccc","aaa","bbb","ccc","aaa","bbb","ccc","aaa","bbb","ccc",
}

func BenchmarkSlice(b *testing.B) {
	test := []int{}
	
	for i := 0; i < b.N; i++ {
		reflect.DeepEqual(test, test)
	}
}

func Benchmark2Slice(b *testing.B) {

	for i := 0; i < b.N; i++ {
		tools.SliceEqual(test,test)
	}
}

func myFunction(arr []string) string {
	var c string
	for _, x := range arr {
		c = x
	}
	return c
}

func myFunction1(arr []string) string {
	var c string
	for i := 0; i < len(arr); i++ {
		c = arr[i]
	}
	return c
}

func randomString5(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(61 + rand.Intn(25))
	}
	return string(bytes)
}

func myFunction3(arr []string){

}

func Benchmark4Slice(b *testing.B) {

	for i := 0; i < b.N; i++ {
		s := make([]string, 0)
		for i := 0; i < 1000000; i++ {
			s = append(s, randomString(10))
		}
		myFunction(s)
	}
}

func Benchmark14Slice(b *testing.B) {

	for i := 0; i < b.N; i++ {
		s := make([]string, 0)
		for i := 0; i < 1000000; i++ {
			s = append(s, randomString(10))
		}
		myFunction1(s)
	}
}
