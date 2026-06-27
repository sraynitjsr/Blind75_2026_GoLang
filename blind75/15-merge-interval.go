package blind75

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		last := result[len(result)-1]

		if intervals[i][0] <= last[1] {
			if intervals[i][1] > last[1] {
				last[1] = intervals[i][1]
			}
		} else {
			result = append(result, intervals[i])
		}
	}

	return result
}

func FifteenMergeInterval() {
	intervals := make([][]int, 0)
	intervals = append(intervals, []int{1, 3}, []int{2, 6}, []int{8, 10}, []int{15, 18})
	fmt.Println("Merged Intervals Are", merge(intervals))
}
