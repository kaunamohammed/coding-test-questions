package slidingwindow

import (
	"math"

	"github.com/kaunamohammed/godsa/utilities"
)

// LengthOfSmallestSubArrayGreaterThanOrEqualToGivenSum finds the smallest subarray's length whose sum is greater than or equal to k 
func LengthOfSmallestSubArrayGreaterThanOrEqualToGivenSum(k int, nums []int) int {

	leftPointer := 0
	currentSum := 0
	smallestLength := math.MaxInt64

	for rightPointer := 0; rightPointer < len(nums); rightPointer++ {

		// accumulating on every iteration
		currentSum += nums[rightPointer]

		// accumulation reaches or goes past the target sum
		for currentSum >= k {

			smallestLength = utilities.Min(smallestLength, (rightPointer - leftPointer) + 1)
			
			// subtract the first element from the previous subarray
			currentSum -= nums[leftPointer]
			leftPointer++

		}

	}

	if smallestLength == math.MaxInt64 {
		return 0
	}

	return smallestLength
}