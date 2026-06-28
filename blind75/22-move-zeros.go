package blind75

import "fmt"

func moveZerosToEnd(nums []int) []int {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			if i != j {
				nums[i], nums[j] = nums[j], nums[i]
			}
			j++
		}
	}
	return nums
}

func moveZerosToBegin(nums []int) []int {
	j := len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] != 0 {
			if i != j {
				nums[i], nums[j] = nums[j], nums[i]
			}
			j--
		}
	}
	return nums
}

func TwentyTwoMoveZeros() {
	nums := []int{0, 1, 0, 3, 12}

	fmt.Println("Original:", nums)

	fmt.Println("Move Zeroes to Start:", moveZerosToBegin(nums))

	fmt.Println("Move Zeroes to End:  ", moveZerosToEnd(nums))
}

// Done
