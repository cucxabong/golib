package sort

import "testing"

func TestMergeSort(t *testing.T) {
	for _, tt := range tables {
		t.Run(tt.desc, func(t *testing.T) {
			MergeSort(tt.in)
			if !IsSorted(tt.in) {
				t.Errorf("exptected %v", tt.expect)
				t.Errorf("got %v", tt.in)
			}
		})
	}
}
