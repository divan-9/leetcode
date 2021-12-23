package linkedlist

import (
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	test := func(name string, slice1 []int, slice2 []int, expected []int) {
		t.Run(name, func(t *testing.T) {
			list1 := createList(slice1)
			list2 := createList(slice2)
			result := mergeTwoLists(list1, list2)
			if !reflect.DeepEqual(result.Slice(100), expected) {
				t.Fatalf("Actual:%v Expected:%v", result.Slice(100), expected)
			}
		})
	}

	test("test1", []int{}, []int{}, []int{})
	test("test2", []int{}, []int{1}, []int{1})
	test("test3", []int{1}, []int{}, []int{1})
	test("test4", []int{1, 2, 3}, []int{1, 2, 4}, []int{1, 1, 2, 2, 3, 4})
}
