package heap

import (
	"testing"
)

func TestLength(t *testing.T) {
	tables := []struct {
		data     []int
		expected int
	}{
		{
			[]int{},
			0,
		},
		{
			[]int{1, 2, 3, 4, 5},
			5,
		},
		{
			[]int{1, 1, 1, 1, 1, 1},
			6,
		},
	}

	for _, table := range tables {
		h := NewBinaryHeap(table.data)
		if h.Length() != table.expected {
			t.Logf("Running TestLength for %v", table.data)
			t.Errorf("Heap's length is incorrect, got %d, expected %d", h.Length(), table.expected)
		}
	}
}

func TestInsert(t *testing.T) {
	data := []int{}
	h := NewBinaryHeap(data)

	t.Run("single insert", func(t *testing.T) {
		h.Insert(10)
		if h.Length() != 1 {
			t.Errorf("Heap's length is incorrect, got %d, expected %d", h.Length(), 1)
		}
		top, _ := h.Top()
		if top != 10 {
			t.Errorf("Incorrect min value, got %d, expected %d", top, 10)
		}
	})

	t.Run("multiple insert", func(t *testing.T) {
		h.Insert(20)
		h.Insert(30)
		h.Insert(8)
		h.Insert(1)

		if h.Length() != 5 {
			t.Errorf("Heap's length is incorrect, got %d, expected %d", h.Length(), 5)
		}
		top, _ := h.Top()
		if top != 1 {
			t.Errorf("Incorrect min value, got %d, expected %d", top, 1)
		}
	})
}

func TestExtract(t *testing.T) {
	h := NewBinaryHeap([]int{})
	t.Run("extract from empty heap", func(t *testing.T) {
		_, err := h.Extract()
		if err != ErrorEmptyHeap {
			t.Errorf("expected '%v', got '%v'", ErrorEmptyHeap, err)
		}
	})

	t.Run("extract from non-empty heap", func(t *testing.T) {
		h.Insert(3)
		h.Insert(4)
		h.Insert(1)
		h.Insert(100)
		h.Insert(20)

		min, _ := h.Extract()
		if min != 1 {
			t.Errorf("Incorrect min value, got %d, expected %d", min, 1)
		}

		if h.Length() != 4 {
			t.Errorf("Incorrect heap's length, got %d, expected %d", h.Length(), 4)
		}

		top, _ := h.Top()
		if top != 3 {
			t.Errorf("Incorrect top value, got %d, exptected %d", top, 3)
		}
	})
}

func TestDelete(t *testing.T) {
	//	data := []int{3, 4, 1, 100, 20}
	h := NewBinaryHeap([]int{})

	t.Run("delete from an empty heap", func(t *testing.T) {
		err := h.Delete(1)
		if err != ErrorEmptyHeap {
			t.Errorf("expected '%v', got '%v'", ErrorValueNotFound, err)
		}
	})

	t.Run("delete from a non-empty heap", func(t *testing.T) {
		h.Insert(3)
		h.Insert(4)
		h.Insert(1)
		h.Insert(100)
		h.Insert(20)

		err := h.Delete(1)

		if err != nil {
			t.Errorf("exptected '%v', got '%v'", nil, err)
		}

		if h.Length() != 4 {
			t.Errorf("Incorrect heap's length, got %d, expected %d", h.Length(), 4)
		}

		top, _ := h.Top()
		if top != 3 {
			t.Errorf("Incorrect top value, got %d, expected %d", top, 3)
		}

		h.Delete(4)
		top, _ = h.Top()
		if top != 3 {
			t.Errorf("Incorrect top value, got %d, exptect %d", top, 3)
		}

		err = h.Delete(200)
		if err != ErrorValueNotFound {
			t.Errorf("got '%v', exptected '%v'", err, ErrorValueNotFound)
		}
	})
}
