package slidingwindow

// SubArraysWithKDistinct dpes
func SubArraysWithKDistinct(arr []int, k int) int {
	
	left, counter := 0, 0
	window := make(map[int]int)

	for right := 0; right < len(arr); right++ {

		window[arr[right]] += 1

		for left < right {

			left++

			if window[left] == 0 {
				delete(window, window[left])
			}

		}

	}

	return counter
}