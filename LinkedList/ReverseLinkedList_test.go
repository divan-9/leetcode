package linkedlist

import (
	"reflect"
	"testing"
)

func TestReverseList(t *testing.T) {
	test := func(name string, inputSlice []int, expected []int) {
		t.Run(name, func(t *testing.T) {
			list := createList(inputSlice)
			result := reverseList(list)
			if !reflect.DeepEqual(result.Slice(100), expected) {
				t.Fatalf("Actual:%v Expected:%v", result.Slice(100), expected)
			}
		})
	}
	test("test1", []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1})
	test("test3", []int{1}, []int{1})
	test("test4", []int{}, []int{})
}

func TestReverseListR(t *testing.T) {
	test := func(name string, inputSlice []int, expected []int) {
		t.Run(name, func(t *testing.T) {
			list := createList(inputSlice)
			result := reverseListR(list)
			if !reflect.DeepEqual(result.Slice(100), expected) {
				t.Fatalf("Actual:%v Expected:%v", result.Slice(100), expected)
			}
		})
	}
	test("test1", []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1})
	test("test3", []int{1}, []int{1})
	test("test4", []int{}, []int{})
}
