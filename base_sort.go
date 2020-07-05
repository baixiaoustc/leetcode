package leetcode

func MergeSort(list []int) []int {
	if len(list) < 2 {
		return list
	}
	return mergeSortRecursive(list)
}

func mergeSortRecursive(list []int) []int {
	length := len(list)
	if length == 1 {
		return []int{list[0]}
	}
	mid := length / 2
	left := mergeSortRecursive(list[0:mid])
	right := mergeSortRecursive(list[mid:length])
	ret := make([]int, length)
	i, j := 0, 0
	for index := 0; index < length; {
		if left[i] <= right[j] {
			ret[index] = left[i]
			index++
			i++
			if i >= len(left) {
				for ; j < len(right); j++ {
					ret[index] = right[j]
					index++
				}
			}
		} else {
			ret[index] = right[j]
			index++
			j++
			if j >= len(right) {
				for ; i < len(left); i++ {
					ret[index] = left[i]
					index++
				}
			}
		}
	}
	return ret
}

func QuickSort(list []int) []int {
	if len(list) < 2 {
		return list
	}
	quickSortRecursive(list, 0, len(list)-1)
	return list
}

func quickSortRecursive(list []int, left, right int) {
	if left >= right {
		return
	}

	pivotIndex := partition1(list, left, right)
	quickSortRecursive(list, left, pivotIndex-1)
	quickSortRecursive(list, pivotIndex+1, right)
}

func partition1(list []int, left, right int) int {
	pivot := list[left] //定左边为pivot，就必须先从右边开始找
	start, end := left, right
	for start < end {
		for start < end && list[end] > pivot {
			end--
		}
		for start < end && list[start] <= pivot {
			start++
		}
		if start >= end {
			break
		}
		list[start], list[end] = list[end], list[start]
	}

	list[left], list[start] = list[start], list[left]
	return end
}
