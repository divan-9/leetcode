package SortingAndSearching

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	test := func(name string, nums1 []int, m int, nums2 []int, n int, expected []int) {
		t.Run(name, func(t *testing.T) {
			merge(nums1, m, nums2, n)
			if !reflect.DeepEqual(nums1, expected) {
				t.Fatalf("Actual:%v Expected:%v", nums1, expected)
			}
		})
	}

	test("test1", []int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6})
	test("test2", []int{1}, 1, []int{}, 0, []int{1})
	test("test3", []int{0}, 0, []int{1}, 1, []int{1})
}
