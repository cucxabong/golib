package sort

import "testing"

func TestQuickSort(t *testing.T) {
	for _, tt := range tables {
		t.Run(tt.desc, func(t *testing.T) {
			QuickSort(tt.in)
			if !IsSorted(tt.in) {
				t.Errorf("exptected %v", tt.expect)
				t.Errorf("got %v", tt.in)
			}
		})
	}
}
