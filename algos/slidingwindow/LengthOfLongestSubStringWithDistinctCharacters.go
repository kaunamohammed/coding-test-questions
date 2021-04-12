package slidingwindow

import "github.com/kaunamohammed/godsa/utilities"

// LengthOfLongestSubStringWithDistinctCharacters returns the longest substring which contains only
// distinct characters
func LengthOfLongestSubStringWithDistinctCharacters(s string, k int) int {

	leftPointer := 0
	longestDistinctSum := 0
	currentSum := 0

	window := make(map[byte]int)

	for rightPointer := 0; rightPointer < len(s); rightPointer++ {

		currentSum++

		// need to first check if we have the character
		// and if not, put it inside the hashmap with a value of 1
		// else increment the number of occurences of the value
		if val, ok := window[s[rightPointer]]; ok {
			window[s[rightPointer]] = val + 1
		} else {
			window[s[rightPointer]] = 1
		}

		// while the hashmap contains more than k characters
		// keep removing the left most character
		for len(window) > k {

			// decrement the current longest sum
			currentSum--
			// calculate the new max longest sum
			longestDistinctSum = utilities.Max(longestDistinctSum, currentSum)

			window[s[leftPointer]] = window[s[leftPointer]] - 1

			if window[s[leftPointer]] == 0 {
				delete(window, s[leftPointer])
			}

			leftPointer++

		}

	}

	return longestDistinctSum
}