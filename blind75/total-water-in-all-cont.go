package blind75

import "fmt"

func trap(height []int) int {
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	water := 0

	for left < right {
		if height[left] < height[right] {
			leftMax = max(leftMax, height[left])
			water += leftMax - height[left]
			left++
		} else {
			rightMax = max(rightMax, height[right])
			water += rightMax - height[right]
			right--
		}
	}

	return water
}

func EightTotalWaterByAllConatiners() {
	fmt.Println("Total Water By Rain Water Trapping is", trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}
