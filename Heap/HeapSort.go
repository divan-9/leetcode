package Heap

func Sort(array []int) {
	for index := len(array) - 1; index >= 0; index-- {
		heapify(array, index, len(array)-1)
	}

	for index := len(array) - 1; index > 0; index-- {
		swap(array, 0, index)
		heapify(array, 0, index-1)
	}
}

func heapify(array []int, root int, stop int) {
	var left = root*2 + 1
	var right = root*2 + 2
	var largest = root

	if left <= stop && array[left] > array[largest] {
		largest = left
	}

	if right <= stop && array[right] > array[largest] {
		largest = right
	}

	if largest != root {
		swap(array, root, largest)
		heapify(array, largest, stop)
	}
}

func swap(array []int, i int, j int) {
	tmp := array[i]
	array[i] = array[j]
	array[j] = tmp
}
