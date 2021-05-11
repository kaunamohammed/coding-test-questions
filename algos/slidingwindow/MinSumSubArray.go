package slidingwindow

import (
	"math"

	"github.com/kaunamohammed/godsa/utilities"
)

//  MinSumSubArray finds the minimum sum subarray of size k
func MinSumSubArray(nums []int, k int) int {

	leftPointer := 0
	currentSum := 0
	minSum := math.MaxInt64
	
	for rightPointer, num := range nums {

		currentSum += num
		
		if rightPointer >= (k - 1) {
			minSum = utilities.Min(minSum, currentSum)
			currentSum -= nums[leftPointer]
			leftPointer++
		}

	}

	return minSum
}