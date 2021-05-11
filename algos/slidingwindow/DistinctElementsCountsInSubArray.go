package slidingwindow

import (
	"fmt"
	"strings"
)

// DistinctElementsCountsInSubArray finds the count of distinct elements in every subarray of size k
func DistinctElementsCountsInSubArray(nums []int, k int) string {

	str := ""
	leftPointer := 0
	window := make(map[int]int)

	for rightPointer := 0; rightPointer < len(nums); rightPointer++ {

		window[nums[rightPointer]] += 1

		if rightPointer >= (k - 1) {
			
			window[nums[leftPointer]] = window[nums[leftPointer]] - 1

			str += fmt.Sprintf("%d%s", len(window), " ")

			if window[nums[leftPointer]] == 0 {
				delete(window, nums[leftPointer])
			}

			leftPointer++
			
		}

	}

	return strings.TrimSpace(str)
}