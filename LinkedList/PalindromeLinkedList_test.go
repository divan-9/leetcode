package linkedlist

import (
	"reflect"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	test := func(name string, slice []int, expected bool) {
		t.Run(name, func(t *testing.T) {
			list := createList(slice)
			result := isPalindrome(list)
			if !reflect.DeepEqual(result, expected) {
				t.Fatalf("Input:%v Actual:%v Expected:%v", slice, result, expected)
			}
		})
	}

	test("test1", []int{}, true)
	test("test2", []int{1}, true)
	test("test3", []int{1, 2, 3}, false)
	test("test4", []int{1, 2, 2, 1}, true)
	test("test5", []int{1, 2, 3, 2, 1}, true)
}
