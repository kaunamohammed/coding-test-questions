package slidingwindow

import "github.com/kaunamohammed/godsa/utilities"

// LengthOfLongestSubStringWithoutRepeatingCharacters returns the length of the longest substring
// without repeating characters
func LengthOfLongestSubStringWithoutRepeatingCharacters(s string) int {

	if len(s) < 1 {
		return 0
	}

	if len(s) == 1 {
		return 1
	}

	leftPointer := 0
	rightPointer := 0
	longest := 0

	window := make(map[byte]bool)

	for rightPointer < len(s) {

		if !window[s[rightPointer]] {
			window[s[rightPointer]] = true
			rightPointer++
			longest = utilities.Max(longest, len(window))
		} else {
			delete(window, s[leftPointer])
			leftPointer++
		}

	}

	return longest
}
