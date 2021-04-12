package algos

import "github.com/kaunamohammed/godsa/utilities"

// Given an array arr, this algorithm finds the maximum sequence of 1's that can be formed with at most k 0's
func MaxSequenceWithContiniousOne(arr []int, k int) int {

	zeroCounter := 0
	leftPointer := 0
	rightPointer := 0
	maxTotal := 0

	for rightPointer < len(arr) { 

		if arr[rightPointer] == 0 {
			zeroCounter++
		}

		for zeroCounter > k {
			if arr[leftPointer] == 0 {
				zeroCounter--
			}			
			leftPointer++	
		}
		
		rightPointer++
		maxTotal = utilities.Max(maxTotal, (rightPointer - leftPointer))
	}

	return maxTotal
}