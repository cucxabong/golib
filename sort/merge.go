package sort

func merge(data []int, l, m, r int) {
	left := append([]int{}, data[l:m+1]...)
	right := append([]int{}, data[m+1:r+1]...)
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			data[l+i+j] = left[i]
			i++
		} else if left[i] > right[j] {
			data[l+i+j] = right[j]
			j++
		} else {
			data[l+i+j] = left[i]
			i++
			data[l+i+j] = right[j]
			j++
		}
	}

	for i < len(left) {
		data[l+i+j] = left[i]
		i++
	}

	for j < len(right) {
		data[l+i+j] = right[j]
		j++
	}
}

func mergeSort(data []int, left, right int) {
	if left == right {
		return
	}
	mid := left + (right-left)/2
	mergeSort(data, left, mid)
	mergeSort(data, mid+1, right)
	merge(data, left, mid, right)
}

func MergeSort(data []int) {
	if len(data) == 0 {
		return
	}
	mergeSort(data, 0, len(data)-1)
}
