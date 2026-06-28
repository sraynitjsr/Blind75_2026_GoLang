package blind75

import "fmt"

type BST1 struct {
	Val   int
	Left  *BST1
	Right *BST1
}

func insert1(root *BST1, val int) *BST1 {
	if root == nil {
		return &BST1{Val: val}
	}

	if val < root.Val {
		root.Left = insert1(root.Left, val)
	} else {
		root.Right = insert1(root.Right, val)
	}

	return root
}

func invert(root *BST1) *BST1 {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left

	invert(root.Left)
	invert(root.Right)

	return root
}

func inOrder1(root *BST1) {
	if root == nil {
		return
	}
	inOrder1(root.Left)
	fmt.Print(root.Val, " ")
	inOrder1(root.Right)
}

func EighteenInvertBST() {
	var root *BST1

	values := []int{4, 2, 7, 1, 3, 6, 9}
	for _, v := range values {
		root = insert1(root, v)
	}

	fmt.Print("Original Tree (InOrder): ")
	inOrder1(root)
	fmt.Println()

	root = invert(root)

	fmt.Print("inverted Tree (InOrder): ")
	inOrder1(root)
	fmt.Println()
}

// Done
