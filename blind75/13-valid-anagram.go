package blind75

import "fmt"

func isAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	myMap := make(map[byte]int)

	for i := 0; i < len(s1); i++ {
		myMap[s1[i]]++
		if myMap[s1[i]] == 0 {
			delete(myMap, s1[i])
		}

		myMap[s2[i]]--
		if myMap[s2[i]] == 0 {
			delete(myMap, s2[i])
		}
	}

	return len(myMap) == 0
}

func FifteenVallidAnagram() {
	fmt.Printf("The Strings %v and %v Are Anagrams or Not => %v", "anagram", "nagaram", isAnagram("nagaram", "anagram"))
	fmt.Println()
}

// Done
