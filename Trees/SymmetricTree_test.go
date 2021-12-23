package trees

import (
	"testing"
)

func TestIsSymmetric(t *testing.T) {
	res := isSymmetric(nil)
	if !res {
		t.Fatal("Expected true")
	}
}

func TestIsSymmetric2(t *testing.T) {
	input := &TreeNode{
		Val: 1,
	}
	res := isSymmetric(input)
	if !res {
		t.Fatal("Expected true")
	}
}

func TestIsSymmetric3(t *testing.T) {
	input := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 3}},
		Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}},
	}
	res := isSymmetric(input)
	if !res {
		t.Fatal("Expected true")
	}
}

func TestIsSymmetric4(t *testing.T) {
	input := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 4},
	}
	res := isSymmetric(input)
	if res {
		t.Fatal("Expected false")
	}
}

func TestIsSymmetric5(t *testing.T) {
	input := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 3}},
		Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}},
	}
	res := isSymmetric(input)
	if res {
		t.Fatal("Expected false")
	}
}

func TestIsSymmetric6(t *testing.T) {
	input := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 3}},
		Right: &TreeNode{Val: 2},
	}
	res := isSymmetric(input)
	if res {
		t.Fatal("Expected false")
	}
}
