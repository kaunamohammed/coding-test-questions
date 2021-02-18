package algos

// MutateTheArray mutates the array
func MutateTheArray(nums []int) []int {

	var b []int
	n := len(nums)

	for i := 0; i < n; i++ {

		// if i - 1 is lesser than 0 and i + 1 is lesser than n
		if i-1 < 0 && i+1 < n {
			b[i] = 0 + nums[i] + nums[i+1]

			// if i + 1 is greater or equal n and i - 1 is greater than -1
		} else if i+1 >= n && i-1 >= 0 {
			b[i] = nums[i-1] + nums[i] + 0

			// if i - 1 is lesser than 0 and i + 1 is greater than n
		} else if i-1 < 0 && i+1 >= n {
			b[i] = 0 + nums[i] + 0

		} else {
			b[i] = nums[i-1] + nums[i] + nums[i+1]
		}
	}

	return b
}