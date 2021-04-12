package algos


// Given an array arr, this algorithm finds the total number of subarrays that when summed up equal to k
func NumOfSubarraysEqualToK(arr []int, k int) int {

	counter := 0
	currentSum := 0
	leftPointer := 0

	for rightPointer, value := range arr  {

		currentSum += value

		for currentSum >= k {
			if currentSum == k {
				counter++
			}
			currentSum -= arr[leftPointer]
			leftPointer++

		}

		rightPointer++

	}

	return counter
}