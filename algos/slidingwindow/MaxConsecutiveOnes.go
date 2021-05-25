package slidingwindow

import (
	"github.com/kaunamohammed/godsa/utilities"
)

func MaxConsecutiveOnes(nums []int) int {

	leftPointer := 0
	rightPointer := 0
	maxConsecutiveOnes := 0

	for rightPointer < len(nums) {

		if nums[rightPointer] == 0 {
			leftPointer = rightPointer + 1
		}
		rightPointer++

		maxConsecutiveOnes = utilities.Max(maxConsecutiveOnes, (rightPointer - leftPointer))

	}
    
	return maxConsecutiveOnes
}