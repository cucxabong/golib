package sort

import "testing"

var ints = []int{100, -20, 30, 25, 80, -192929, -200, 40}
var sortedInts = []int{-192929, -200, -20, 25, 30, 40, 80, 100}

func TestQuickSort(t *testing.T) {
	data := ints
	QuickSort(data)

	if !IsSorted(data) {
		t.Errorf("expected %v", sortedInts)
		t.Errorf("got: %v", data)
	}
}
