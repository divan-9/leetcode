package linkedlist

import (
	"reflect"
	"testing"
)

func TestCreateList(t *testing.T) {
	test := func(name string, inputSlice []int, n int, expected []int) {
		t.Run(name, func(t *testing.T) {
			list := createCycledList(inputSlice, n)
			result := list.Slice(5)
			if !reflect.DeepEqual(result, expected) {
				t.Fatalf("Input:%v:%v Actual:%v Expected:%v", inputSlice, n, result, expected)
			}
		})
	}
	test("test1", []int{}, -1, []int{})
	test("test2", []int{1}, -1, []int{1})
	test("test3", []int{1}, 0, []int{1, 1, 1, 1, 1})
	test("test4", []int{1, 2}, 1, []int{1, 2, 2, 2, 2})
	test("test5", []int{1, 2, 3}, 1, []int{1, 2, 3, 2, 3})
}
