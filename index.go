package heap

func IndexParent(childIndex int) int {
	if childIndex == 0 {
		return -1
	}
	return (childIndex - 1) / 2
}

func IndexLeftChild(parentIndex int) int {
	return parentIndex*2 + 1
}

func IndexRightChild(parentIndex int) int {
	return parentIndex*2 + 2
}
