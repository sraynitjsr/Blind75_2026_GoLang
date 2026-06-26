package blind75

import "fmt"

func productExceptSelf(nums []int) []int {
	length := len(nums)

	left, right, answer := make([]int, length), make([]int, length), make([]int, length)

	left[0], right[length-1] = 1, 1

	for i := 1; i < length; i++ {
		left[i] = left[i-1] * nums[i-1]
	}

	for i := length - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i+1]
	}

	for i := range length {
		answer[i] = left[i] * right[i]
	}

	return answer
}

func FourArrayProductExceptSelf() {
	mySlice := []int{1, 2, 3, 4}
	fmt.Printf("Product of %v is %v", mySlice, productExceptSelf(mySlice))
}
