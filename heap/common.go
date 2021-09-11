package heap

func parent(idx int) int {
	return (idx - 1) / 2
}

func leftChild(idx int) int {
	return 2*idx + 1
}

func rightChild(idx int) int {
	return 2*idx + 2
}
