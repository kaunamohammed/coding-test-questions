package algos

import (

	"github.com/kaunamohammed/godsa/utilities"
)

// RemoveElement remove all instances of val in-place and return the new length containing all other elements
func RemoveElement(nums []int, val int) int {

	count := 0
	leftPointer := 0
	rightPointer := leftPointer + 1

	for rightPointer <= len(nums) - 1 {
		if nums[leftPointer] != val {
			leftPointer++
			rightPointer++
		} else if nums[leftPointer] == val && nums[rightPointer] == val {
			rightPointer++
		} else {
			utilities.Swap(nums, leftPointer, rightPointer)
			leftPointer++
			rightPointer++
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			count++
		}
	}

	return count
}