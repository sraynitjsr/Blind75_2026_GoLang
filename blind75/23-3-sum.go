package blind75

import (
	"fmt"
	"strconv"
)

func TwentyThree3Sum() {
	nums, tripletSeen, res := []int{-1, 0, 1, 2, -1, -4}, make(map[string]bool), [][]int{}
	for i := range nums {
		thirdValuePresenceMap := make(map[int]bool)
		for j := i + 1; j < len(nums); j++ {
			thirdValue := -(nums[i] + nums[j])
			if thirdValuePresenceMap[thirdValue] {
				a, b, c := nums[i], nums[j], thirdValue
				if a > b {
					a, b = b, a
				}
				if b > c {
					b, c = c, b
				}
				if a > b {
					a, b = b, a
				}
				key := strconv.Itoa(a) + "," + strconv.Itoa(b) + "," + strconv.Itoa(c)
				if !tripletSeen[key] {
					tripletSeen[key] = true
					res = append(res, []int{a, b, c})
				}
			}
			thirdValuePresenceMap[nums[j]] = true
		}
	}
	fmt.Println("Trplets Are", res)
}
