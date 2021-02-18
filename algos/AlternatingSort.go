package algos


// AlternatingSort sorts an array of ints by placing the nums element at index 1 in index 1 of a new array 
// then the element at index lens(nums) - 1, then the element at index 2, then index lens(nums) - 2 and so on
// until the array has been fully sorted.
// if the resultong array of the sort is not in ascending order,
// the function returns false along with an array of the indices of elements that first break the constraint
// otherwise it returns true with the elements sorted in ascending order
func AlternatingSort(nums []int) (bool, []int) {

	if len(nums) == 0 || len(nums) == 1 {
		return true, nums
	}

	result := make([]int, len(nums))

	leftPointer := 0
	rightPointer := len(nums) - 1

	result[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		if i % 2 == 0 {
			result[i] = nums[leftPointer]
		} else {
			result[i] = nums[rightPointer]
			rightPointer--
			leftPointer++
		}

		if result[i] < result[i - 1] {
			return false, []int { i, i - 1}
		}
	}

	return true, result
}