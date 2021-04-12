package slidingwindow


// NumSubArrayProductLessThanK does
func NumSubArrayProductLessThanK(nums []int, k int) int  {

	counter, left, right := 0, 0, 0
	currentProduct := 1

	for right < len(nums) {
				
		currentProduct *= nums[right]

		for currentProduct >= k  {
 			currentProduct /= nums[left]
			left++
		}
		counter += (right - left) + 1
		right++

	}

	return counter
}

// NumSubArrayProductLessThanKSlow slow
func NumSubArrayProductLessThanKSlow(nums []int, k int) int {
    
    counter, left := 0, 0

	for right := 0; right < len(nums); right++ {
		
		currentProduct := 1
		left = right + 1
		currentProduct *= nums[right]

		if currentProduct < k {
			counter++
		}

		for currentProduct < k && left < len(nums) {

			currentProduct *= nums[left]
			left++

			if currentProduct < k {
				counter++
			} 

		}

	}

	return counter
    
}