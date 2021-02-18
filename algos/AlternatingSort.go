package algos

// AlternatingSort sorts an array of ints by placing the nums element at index 1 in 
// then the element at index lens(nums) - 1 
func AlternatingSort(nums []int) bool {

	if len(nums) == 0 || len(nums) == 1 {
		return true
	}

	temp := make([]int, len(nums))

	leftPointer := 0
	rightPointer := len(nums) - 1

	temp[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		if leftPointer%2 == 0 {
			temp[i] = nums[leftPointer]
		} else {
			temp[i] = nums[rightPointer]
			rightPointer--
			leftPointer++
		}
	}

	for i := 0; i < len(temp)-1; i++ {
		if temp[i] > temp[i+1] {
			return false
		}
	}

	return true
}