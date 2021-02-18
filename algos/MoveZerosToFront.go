package algos

import "github.com/kaunamohammed/godsa/utilities"

// MoveZerosToFront moves all occurences of the number 0 to 
// the front of the array
func MoveZerosToFront(nums []int) []int {

	leftPointer := len(nums) - 2
	rightPointer := len(nums) - 1

	for leftPointer >= 0 {
		if nums[rightPointer] != 0 {
			leftPointer--
			rightPointer--
		} else if nums[rightPointer] == 0 && nums[leftPointer] != 0 {
			utilities.Swap(nums, leftPointer, rightPointer)
			leftPointer--
			rightPointer--
		} else {
			leftPointer--
		}
	}

	return nums
}