package algos

// SubArraysWithKDistinct dpes
func SubArraysWithKDistinct(A []int, K int) int {
	
	left, counter := 0, 0
	window := make(map[int]int)

	for right := 0; right < len(A); right++ {

		if val, ok := window[A[right]]; ok {
			window[A[right]] = val + 1
		} else {
			window[A[right]] = 1
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