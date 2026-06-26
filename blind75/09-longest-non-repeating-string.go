package blind75

import "fmt"

func lengthOfLongestSubstring(s string) int {
	left, maxLen, myMap := 0, 0, make(map[byte]int)
	for right := 0; right < len(s); right++ {
		if idx, present := myMap[s[right]]; present && idx >= left {
			left = idx + 1
		}
		myMap[s[right]] = right
		maxLen = max(maxLen, right-left+1)
	}
	return maxLen
}

func NineLongestNonRepeatingString() {
	fmt.Println("Longest Non Repeating Sub String Length is", lengthOfLongestSubstring("pwwkew"))
}
