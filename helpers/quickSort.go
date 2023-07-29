package helpers

func partition(array []int, start int, end int) ([]int, int) {
	pivot := array[end]
	newPivotIndex := start
	for j := start; j < end; j++ {
		if array[j] < pivot {
			//swap
			array[newPivotIndex], array[j] = array[j], array[newPivotIndex]
			newPivotIndex++
		}
	}
	array[newPivotIndex], array[end] = array[end], array[newPivotIndex]
	return array, newPivotIndex
}

func quickSort(arr []int, start int, end int) []int {
	if start < end {
		var p int
		arr, p = partition(arr, start, end)
		arr = quickSort(arr, start, p-1)
		arr = quickSort(arr, p+1, end)
	}
	return arr
}

func QuickSort(arr []int, start int, end int) []int {
	return quickSort(arr, start, end)
}
