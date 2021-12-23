package linkedlist

import (
	"reflect"
	"testing"
)

func TestHasCycle(t *testing.T) {
	test := func(name string, inputSlice []int, n int, expected bool) {
		t.Run(name, func(t *testing.T) {
			list := createCycledList(inputSlice, n)
			result := hasCycle(list)
			if !reflect.DeepEqual(result, expected) {
				t.Fatalf("Input:%v:%v Actual:%v Expected:%v", inputSlice, n, result, expected)
			}
		})
	}
	test("test1", []int{}, -1, false)
	test("test2", []int{1}, -1, false)
	test("test3", []int{1}, 0, true)
	test("test4", []int{1, 2}, 1, true)
	test("test5", []int{1, 2, 3}, 1, true)
}
