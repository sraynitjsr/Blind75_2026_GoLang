package blind75

import "fmt"

func twoSum(mySlice []int, target int) []int {
	myMap := make(map[int]int)
	for i, v := range mySlice {
		if j, present := myMap[target-v]; present {
			return []int{i, j}
		} else {
			myMap[v] = i
		}
	}
	return []int{}
}

func OneTwoSum() {
	mySlice, target := []int{0, -1, 2, -3, 1}, 1

	fmt.Println("Requires Indeces Are", twoSum(mySlice, target))
}

// Done
