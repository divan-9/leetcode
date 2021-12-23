package trees

import (
	"testing"
)

func TestMaxDepth1(t *testing.T) {
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
	res := maxDepth(&input)
	if res != 3 {
		t.Fatal("Expected 3, got", res)
	}
}

func TestMaxDepth2(t *testing.T) {
	input := TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val: 2,
		},
	}
	res := maxDepth(&input)
	if res != 2 {
		t.Fatal("Expected 2, got", res)
	}
}

func TestMaxDepth3(t *testing.T) {
	res := maxDepth(nil)
	if res != 0 {
		t.Fatal("Expected 0, got", res)
	}
}

func TestMaxDepth4(t *testing.T) {
	input := TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	res := maxDepth(&input)
	if res != 1 {
		t.Fatal("Expected 1, got", res)
	}
}
