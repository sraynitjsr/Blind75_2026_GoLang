package blind75

import "fmt"

func maxProduct(nums []int) int {
	maxProduct, minProduct, ans := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			minProduct, maxProduct = maxProduct, minProduct
		}
		maxProduct = max(nums[i], maxProduct*nums[i])
		minProduct = min(nums[i], minProduct*nums[i])

		ans = max(ans, maxProduct)
	}

	return ans
}

func SixMaximumProductSubArray() {
	fmt.Println("Maximum Product is", maxProduct([]int{2, 3, -2, 4}))
}

// Done
