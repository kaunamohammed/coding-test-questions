package slidingwindow

import (
	"math"

	"github.com/kaunamohammed/godsa/utilities"
)

// LengthOfMinSubArray finds the minimal length of a contiguous subarray of which the sum is greater than or equal to k
func LengthOfMinSubArray(nums []int, k int) int {

	leftPointer := 0
	currentSum := 0
	smallestLength := math.MaxInt64

	for rightPointer := 0; rightPointer < len(nums); rightPointer++ {

		currentSum += nums[rightPointer]

		for currentSum >= k {
			smallestLength = utilities.Min(smallestLength, (rightPointer - leftPointer) + 1)
			currentSum -= nums[leftPointer]
			leftPointer++
		}

	}

	if smallestLength == math.MaxInt64 {
		return 0
	}
	
	return smallestLength
}