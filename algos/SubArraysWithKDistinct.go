package algos

// SubArraysWithKDistinct dpes
func SubArraysWithKDistinct(arr []int, k int) int {
	
	left, counter := 0, 0
	window := make(map[int]int)

	for right := 0; right < len(arr); right++ {

		if val, ok := window[arr[right]]; ok {
			window[arr[right]] = val + 1
		} else {
			window[arr[right]] = 1
		}

		for left < right {

			left++

			if window[left] == 0 {
				delete(window, window[left])
			}

		}


	}
	

	return counter
}