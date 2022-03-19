package Heap

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	test := func(name string, input []int, expected []int) {
		t.Run(name, func(t *testing.T) {
			result := make([]int, len(input))
			copy(result, input)
			Sort(result)
			if !reflect.DeepEqual(result, expected) {
				t.Fatalf("Input:%v Actual:%v Expected:%v", input, result, expected)
			}
		})
	}
	test("test1", []int{0}, []int{0})
	test("test2", []int{0, 1}, []int{0, 1})
	test("test3", []int{1, 0}, []int{0, 1})
	test("test4", []int{1, 0, 2}, []int{0, 1, 2})
	test("test5", []int{0, 1, 2, 3, 4, 6, 7, 8, 9, 10, 11, 12, 5}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12})
}
