package blind75

import "fmt"

type BST2 struct {
	Val   int
	Left  *BST2
	Right *BST2
}

func insert2(root *BST2, val int) *BST2 {
	if root == nil {
		return &BST2{Val: val}
	}

	if val < root.Val {
		root.Left = insert2(root.Left, val)
	} else {
		root.Right = insert2(root.Right, val)
	}

	return root
}

func searchInBST(root *BST2, val int) *BST2 {
	if root == nil || root.Val == val {
		return root
	}

	if val < root.Val {
		return searchInBST(root.Left, val)
	}

	return searchInBST(root.Right, val)
}

func NinteenInvertBST() {
	var root *BST2

	values := []int{4, 2, 7, 1, 3, 6, 9}
	for _, v := range values {
		root = insert2(root, v)
	}

	key := 4
	node := searchInBST(root, key)

	if node != nil {
		fmt.Printf("Key %d Found.\n", key)
	} else {
		fmt.Printf("Key %d Not Found.\n", key)
	}
}

// Done
