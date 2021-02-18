package algos

import "github.com/kaunamohammed/godsa/utilities"

// LongestSubSetZeroSum returns the length of the longest subset of numbers in an array that
// sums up to 0
func LongestSubSetZeroSum(nums []int) int {

	currentSum := 0
	longest := 0
	hashmap := make(map[int]int)

	for i := 0; i < len(nums); i++ {

		currentSum += nums[i]

		if currentSum == 0 {
			longest = i + 1
		} else {
			if val, ok := hashmap[currentSum]; ok {
				longest = utilities.Max(longest, i - val) 
			} else {
				hashmap[currentSum] = i
			}
		}		
	}
	
	return longest
}