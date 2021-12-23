// https://leetcode.com/explore/interview/card/top-interview-questions-easy/93/linked-list/603/

package linkedlist

import (
	"reflect"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	test := func(name string, inputSlice []int, n int, expected []int) {
		t.Run(name, func(t *testing.T) {
			list := createList(inputSlice)
			result := removeNthFromEnd(list, n)
			if !reflect.DeepEqual(result.Slice(100), expected) {
				t.Fatalf("Actual:%v Expected:%v", result.Slice(100), expected)
			}
		})
	}
	test("test1", []int{1, 2, 3, 4, 5}, 2, []int{1, 2, 3, 5})
	test("test2", []int{1, 2, 3, 4, 5}, 3, []int{1, 2, 4, 5})
	test("test3", []int{1}, 1, []int{})
	test("test4", []int{1, 2}, 2, []int{2})
}
