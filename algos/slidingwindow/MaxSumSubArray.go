package slidingwindow

import (
	"math"

	"github.com/kaunamohammed/godsa/utilities"
)

// MaxSumSubArray returns the maximum sum of elements in an array given a constraint on the total amount
// of numbers in a subarray
func MaxSumSubArray(nums []int, k int) int {

	maxSum := math.MinInt64
	currentSum := 0
	leftPointer := 0

	for rightPointer, num := range nums {
		currentSum += num
		if rightPointer >= (k - 1) {
			maxSum =  utilities.Max(maxSum, currentSum)
			currentSum -= nums[leftPointer]
			leftPointer++
		}
	}

	return maxSum
}