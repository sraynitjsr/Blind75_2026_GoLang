package blind75

import "fmt"

func tweleveMaxConsecutiveOnesKFlipsAllowed(nums []int, k int) int {
	left, maxLen, zeros := 0, 0, 0
	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			zeros++
		}
		for zeros > k {
			if nums[left] == 0 {
				zeros--
			}
			left++
		}
		maxLen = max(maxLen, maxLen, right-left+1)
	}
	return maxLen
}

func TweleveMaxConsecutiveOnes() {
	fmt.Println("Maximum Consecutive 1s With Atmost K Flips Allowed is", tweleveMaxConsecutiveOnesKFlipsAllowed([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2))
}

// Done
