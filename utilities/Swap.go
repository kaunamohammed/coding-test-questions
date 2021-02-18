package utilities

func Swap(arr []int, a, b int) {
	arr[a], arr[b] = arr[b], arr[a]
}