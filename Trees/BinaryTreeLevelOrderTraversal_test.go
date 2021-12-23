package trees

import (
	"reflect"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	input := TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	expected := [][]int{{3}, {9, 20}, {15, 7}}
	res := levelOrder(&input)
	if !reflect.DeepEqual(res, expected) {
		t.Fatalf("Expected:%v Actual:%v", expected, res)
	}
}

func TestLevelOrder2(t *testing.T) {
	var input *TreeNode
	expected := [][]int{}
	res := levelOrder(input)
	if !reflect.DeepEqual(res, expected) {
		t.Fatalf("Expected:%v Actual:%v", expected, res)
	}
}

func TestLevelOrder3(t *testing.T) {
	input := TreeNode{Val: 1}
	expected := [][]int{{1}}
	res := levelOrder(&input)
	if !reflect.DeepEqual(res, expected) {
		t.Fatalf("Expected:%v Actual:%v", expected, res)
	}
}
