package blind75

import (
	"fmt"
	"slices"
)

func groupAnagrams(str []string) [][]string {
	myMap := make(map[string][]string)

	for _, s := range str {
		sInBytes := []byte(s)
		slices.Sort(sInBytes)
		newKeyForMap := string(sInBytes)
		myMap[newKeyForMap] = append(myMap[newKeyForMap], s)
	}

	ans := make([][]string, 0)
	for _, value := range myMap {
		ans = append(ans, value)
	}

	return ans
}

func FourteenGroupAnagrams() {
	str := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println("Group of Anagrams is", groupAnagrams(str))
}

// Done
