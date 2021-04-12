package slidingwindow

// LengthOfSmallestSubArrayGivenSum returns the length of the smallest sub array that 
// adds up to a given sum
func LengthOfSmallestSubArrayGivenSum(sum int, nums []int) int {

	leftPointer := 0
	currentSum := 0
	smallestSubArray := 0

	for right := 0; right < len(nums); right++ {

		// accumulating on every iteration
		currentSum += nums[right]

		// accumulation reaches or goes past the target sum
		for currentSum >= sum {

			if smallestSubArray == 1 {
				return smallestSubArray
			}

			if smallestSubArray < currentSum { smallestSubArray = (right + 1) - leftPointer }

			// subtract the first element from the previous subarray
			currentSum -= nums[leftPointer]
			leftPointer++

		}

	}

	return smallestSubArray
}