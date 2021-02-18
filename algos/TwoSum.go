package algos

// TwoSum returns the sum of two numbers in an array that sums up to a target
func TwoSum(nums []int, target int) [2]int {

	start := 0
	end := len(nums) - 1
	var result [2]int

	for start < end {
		sum := nums[start] + nums[end]

		if sum == target {
			result[0] = start
			result[1] = end
			break
		} else if sum > target {
			end--
		} else {
			start++
		}
	}

	return result
}