package slidingwindow

import "github.com/kaunamohammed/godsa/utilities"

// MoveZerosToBack moves all occurences of the number 0 to 
// the back of the array
func MoveZerosToBack(nums []int) []int {

	leftPointer := 0
	rightPointer := leftPointer + 1

	for rightPointer <= len(nums) - 1 {
		if nums[leftPointer] == 0 && nums[rightPointer] == 0 {
		} else if nums[leftPointer] != 0 && nums[rightPointer] != 0 || nums[leftPointer] != 0 && nums[rightPointer] == 0 {
			leftPointer++
		} else {
			utilities.Swap(nums, leftPointer, rightPointer)
			leftPointer++
		}
		rightPointer++
	}

	return nums
}