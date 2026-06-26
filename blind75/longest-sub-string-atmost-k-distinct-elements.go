package blind75

import "fmt"

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	left, maxLen, myMap := 0, 0, make(map[byte]int)
	for right := 0; right < len(s); right++ {
		myMap[s[right]]++
		for len(myMap) > k {
			myMap[s[left]]--
			if myMap[s[left]] == 0 {
				delete(myMap, s[left])
			}
			left++
		}
		maxLen = max(maxLen, right-left+1)
	}
	return maxLen
}

func TenLengthOfLongestSubstringKDistinct() {
	fmt.Println("Longest Sub String With Atmost K Distinct Elements is", lengthOfLongestSubstringKDistinct("abcadcacacaca", 3))
}
