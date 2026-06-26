package blind75

import "fmt"

func maxArea(height []int) int {
	left, right, maxArea := 0, len(height)-1, 0

	for left < right {
		h := min(height[left], height[right])

		area := h * (right - left)

		maxArea = max(maxArea, area)

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}

func SevenContainerWithMaximumWater() {
	fmt.Println("Container With Most Water Has Area", maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
