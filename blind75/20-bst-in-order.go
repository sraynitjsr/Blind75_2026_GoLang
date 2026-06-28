package blind75

import "fmt"

type BST3 struct {
	Data  int
	Left  *BST3
	Right *BST3
}

func insert3(root *BST3, data int) *BST3 {
	if root == nil {
		return &BST3{Data: data}
	}

	if data < root.Data {
		root.Left = insert3(root.Left, data)
	} else {
		root.Right = insert3(root.Right, data)
	}

	return root
}

func inOrderTraversal(root *BST3) {
	if root == nil {
		return
	}

	inOrderTraversal(root.Left)
	fmt.Print(root.Data, " ")
	inOrderTraversal(root.Right)
}

func TwentyBSTInOrderTraversal() {
	var root *BST3

	values := []int{5, 3, 7, 2, 4, 6, 1, 8}

	for _, v := range values {
		root = insert3(root, v)
	}

	fmt.Print("In Order Traversal: ")
	inOrderTraversal(root)
	fmt.Println()
}
