package trees

import (
	"testing"
)

func TestIsValidBST(t *testing.T) {
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
	res := isValidBST(&input)
	if res {
		t.Fatal("Expected false")
	}
}

func TestIsValidBST2(t *testing.T) {
	input := TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 6,
			},
		},
	}
	res := isValidBST(&input)
	if res {
		t.Fatal("Expected false")
	}
}

func TestIsValidBST3(t *testing.T) {
	res := isValidBST(nil)
	if !res {
		t.Fatal("Expected true")
	}
}

func TestIsValidBST4(t *testing.T) {
	input := TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
		},
		Right: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	res := isValidBST(&input)
	if res {
		t.Fatal("Expected false")
	}
}

func TestIsValidBST5(t *testing.T) {
	input := TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
		Right: &TreeNode{
			Val: 7,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 8,
			},
		},
	}
	res := isValidBST(&input)
	if res {
		t.Fatal("Expected false")
	}
}
