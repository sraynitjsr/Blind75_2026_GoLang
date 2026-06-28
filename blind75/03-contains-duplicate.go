package blind75

import "fmt"

func containsDuplicate(nums []int) bool {
	mySet := make(map[int]struct{})

	for _, data := range nums {
		if _, present := mySet[data]; present {
			return true
		} else {
			mySet[data] = struct{}{}
		}
	}

	return false
}

func ThreeContainsDuplicate() {
	nums := []int{1, 2, 3, 1}

	fmt.Println("Contains Dupicate - ", containsDuplicate(nums))
}

// Done
