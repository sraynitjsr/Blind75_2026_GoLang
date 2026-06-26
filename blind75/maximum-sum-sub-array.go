package blind75

import "fmt"

func maxSubArray(nums []int) int {
	currentSum, maxSum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		currentSum = max(nums[i], nums[i]+currentSum)
		maxSum = max(maxSum, currentSum)
	}
	return maxSum
}

func FiveMaximumSumSubArray() {
	mySlice := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println("Maximum Sum Sub Array is", maxSubArray(mySlice))
}
