package main

func bubbleSort(array []int) {

}

func qs(arr *[]int, high int, low int) {

	//base Case
	if low >= high {
		return
	}
	pivot_idx := partition(&(*arr), high, low)
	qs(&(*arr), pivot_idx-1, low)
	qs(&(*arr), high, pivot_idx+1)

}

func partition(arr *[]int, low int, high int) int {

	pivot := (*arr)[low]
	idx := low - 1

	for i, _ := range *arr {
		if (*arr)[i] <= pivot {
			//Move the element to the lower side of the pivot
			idx++
			temp := (*arr)[i]
			(*arr)[i] = (*arr)[idx]
			(*arr)[idx] = temp
		}
	}
	idx++
	temp := (*arr)[idx]
	(*arr)[idx] = (*arr)[high]
	(*arr)[high] = temp

	return idx

}

func quickSort(arr []int) []int {
	//select a pivot
	qs(&arr, 0, len(arr)-1)
	//place the pivot in the its sorted position
	return arr
	//sort the items in the each side
}
