package blind75

import "fmt"

func maxConsecutiveOnes(nums []int) int {
	count, maxCount := 0, 0
	for _, num := range nums {
		if num == 1 {
			count++
			maxCount = max(maxCount, count)
		} else {
			count = 0
		}
	}
	return maxCount
}

func ElevenMaxConsecutiveOnes() {
	fmt.Println("Maximum Consecutive 1s Sub Array Length is", maxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1}))
}
