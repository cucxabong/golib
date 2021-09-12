package sort

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	k := low - 1

	for i := low; i < high; i++ {
		if arr[i] < pivot {
			k++
			arr[k], arr[i] = arr[i], arr[k]
		}
	}

	arr[k+1], arr[high] = arr[high], arr[k+1]
	return k + 1
}

func quickSort(arr []int, low, high int) {
	if low < high {
		pivotIndex := partition(arr, low, high)
		quickSort(arr, low, pivotIndex-1)
		quickSort(arr, pivotIndex+1, high)
	}
}

func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}
