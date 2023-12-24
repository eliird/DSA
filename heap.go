package main

//Max Heap or Min Heap

type heap struct {
	data []int
	len  int
}

func (h heap) insert(value int) {
	h.data = append(h.data, value)
	h.len++
	h.heapifyUp(h.len - 1)
}

func (h heap) delete() int {
	if h.len == 0 {
		return -1
	}
	value := h.data[0]
	h.data[0] = h.data[h.len-1]
	h.len--
	h.data = h.data[:h.len-1]
	h.heapifyDown(0)
	return value
}

func getLeftChild(index int) int {
	childLeft := 2*index + 1
	//childRight := 2*index + 2
	return childLeft
}

func getRightChild(index int) int {
	// childLeft := 2*index + 1
	childRight := 2*index + 2
	return childRight
}
func getParent(index int) int {
	return (index - 1) / 2
}

func (h heap) heapifyDown(idx int) {
	//Needs to be fixed
	if idx > h.len {
		return
	}

	leftChild := getLeftChild(idx)
	rightChild := getRightChild(idx)

	if leftChild >= h.len || rightChild >= h.len {
		return
	}
	min_child := min(leftChild, rightChild)
	if h.data[idx] > min_child {
		temp := h.data[idx]
		h.data[idx] = min_child
		if h.data[leftChild] == min_child {
			h.data[leftChild] = temp
			h.heapifyDown(leftChild)
		} else {
			h.data[rightChild] = temp
			h.heapifyDown((rightChild))
		}
	}

}

func (h heap) heapifyUp(idx int) {
	if idx == 0 {
		return
	}
	parent := getParent(idx)
	if h.data[idx] > h.data[parent] {
		return
	}

	temp := h.data[parent]
	h.data[parent] = h.data[idx]
	h.data[idx] = temp
	h.heapifyUp(parent)
	return
}
