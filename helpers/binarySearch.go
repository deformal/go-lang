package helpers

func BinarySearch(arr []int, t int) int {
	low := 0
	high := len(arr) - 1
	for high >= low {
		mid := low + (high-low)/2

		if arr[mid] == t {
			return mid
		}
		if arr[mid] > t {
			high = mid - 1

		} else {
			low = mid + 1
		}
	}
	return -1
}
