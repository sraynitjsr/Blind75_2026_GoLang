package blind75

import "fmt"

type BST struct {
	Data  int
	Left  *BST
	Right *BST
}

func insert(root *BST, val int) *BST {
	if root == nil {
		return &BST{Data: val}
	}

	if val < root.Data {
		root.Left = insert(root.Left, val)
	} else {
		root.Right = insert(root.Right, val)
	}

	return root
}

func inorder(root *BST, inorderedData *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, inorderedData)
	*inorderedData = append(*inorderedData, root.Data)
	inorder(root.Right, inorderedData)
}

func SeventeenKthSmallestInBST() {
	var mySlice = []int{5, 3, 6, 2, 4, 1}
	var root *BST
	k := 4
	inorderedData := make([]int, 0)

	for _, val := range mySlice {
		root = insert(root, val)
	}

	inorder(root, &inorderedData)

	fmt.Printf("%d-th Smallest Element in The BST is %d\n", k, inorderedData[k-1])
}
