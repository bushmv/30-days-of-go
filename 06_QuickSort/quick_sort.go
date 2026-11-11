package main

func QuickSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	rquicksort(arr, 0, len(arr)-1)
}

func rquicksort(arr []int, l int, r int) {
	if l >= r {
		return
	}
	pivot := partition(arr, l, r)
	rquicksort(arr, l, pivot-1)
	rquicksort(arr, pivot, r)
}

func partition(arr []int, l int, r int) int {
	pivot := arr[r]
	i := l - 1
	for j := l; j < r; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[r] = arr[r], arr[i+1]
	return i + 1
}
