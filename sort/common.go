package sort

// IsSorted return true if the input slice is sorted in non-decreasing order
func IsSorted(data []int) bool {
	for i := 0; i < len(data)-1; i++ {
		if data[i] > data[i+1] {
			return false
		}
	}

	return true
}
