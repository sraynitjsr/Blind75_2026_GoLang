package blind75

import "fmt"

type BST1 struct {
	Val   int
	Left  *BST1
	Right *BST1
}

func Insert(root *BST1, val int) *BST1 {
	if root == nil {
		return &BST1{Val: val}
	}

	if val < root.Val {
		root.Left = Insert(root.Left, val)
	} else {
		root.Right = Insert(root.Right, val)
	}

	return root
}

func Invert(root *BST1) *BST1 {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left

	Invert(root.Left)
	Invert(root.Right)

	return root
}

func InOrder(root *BST1) {
	if root == nil {
		return
	}
	InOrder(root.Left)
	fmt.Print(root.Val, " ")
	InOrder(root.Right)
}

func EighteenInvertBST() {
	var root *BST1

	values := []int{4, 2, 7, 1, 3, 6, 9}
	for _, v := range values {
		root = Insert(root, v)
	}

	fmt.Print("Original Tree (InOrder): ")
	InOrder(root)
	fmt.Println()

	root = Invert(root)

	fmt.Print("Inverted Tree (InOrder): ")
	InOrder(root)
	fmt.Println()
}
