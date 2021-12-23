package trees

import (
	"reflect"
	"testing"
)

var tests = []struct {
	name     string
	input    []int
	expected *TreeNode
}{
	{"test1", []int{}, nil},
	{"test2", []int{0, 1}, tr(1, leaf(0), nil)},
	{"test3", []int{-10, -3, 0, 5, 9}, tr(0,
		tr(-3,
			leaf(-10),
			nil),
		tr(9,
			leaf(5),
			nil))},
}

func leaf(val int) *TreeNode {
	return &TreeNode{Val: val}
}

func tr(val int, left, right *TreeNode) *TreeNode {
	return &TreeNode{Val: val, Left: left, Right: right}
}

func TestSortedArrayToBST(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := sortedArrayToBST(test.input)
			if !reflect.DeepEqual(res, test.expected) {
				t.Fatalf("Expected:%v Actual:%v Input:%v", test.expected, res, test.input)
			}
		})
	}
}
