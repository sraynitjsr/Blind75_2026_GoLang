package main

import (
	"fmt"
	"sray/blind75"
)

func main() {
	fmt.Println("Blind75 Beast - GoLang DSA")

	fmt.Println("\nInside 01 - Two Sum Problem")
	blind75.OneTwoSum()

	fmt.Println("\nInside 02 - Stock Buy Sell Problem")
	blind75.TwoStockBuySell()

	fmt.Println("\nInside 03 - Contains Duplicate Problem")
	blind75.ThreeContainsDuplicate()

	fmt.Println("\nInside 04 - Product of Array Except Self Problem")
	blind75.FourArrayProductExceptSelf()

	fmt.Println("\nInside 05 - Maximum Sum Sub Array Problem")
	blind75.FiveMaximumSumSubArray()

	fmt.Println("\nInside 06 - Maximum Product Sub Array Problem")
	blind75.SixMaximumProductSubArray()

	fmt.Println("\nInside 07 - Container With Maximum Water Problem")
	blind75.SevenContainerWithMaximumWater()

	fmt.Println("\nInside 08 - Rain Water Trapping Problem")
	blind75.EightTotalWaterByAllConatiners()

	fmt.Println("\nInside 09 - Longest Non Repeating String Length Problem")
	blind75.NineLongestNonRepeatingString()

	fmt.Println("\nInside 10 - Longest Non Repeating String With Atmost K Distinct Elements")
	blind75.TenLengthOfLongestSubstringKDistinct()
}
