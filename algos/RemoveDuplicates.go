package algos

// RemoveDuplicates removes duplicates from an array
func RemoveDuplicates(nums []int) []int {

	newArr := make([]int, 0)
	unique := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		
		if val, ok := unique[nums[i]]; ok {
			unique[nums[i]] = val + 1
		} else {
			unique[nums[i]] = 1
		}

	}

	for k := range unique {
        newArr = append(newArr, k)
    }

	return newArr
}