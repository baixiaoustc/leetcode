package leetcode

//删除数组下标为n的元素
func delArray(arr []int, n int) []int {
	length := len(arr)
	if length == 0 {
		return arr
	}

	for i := n + 1; i < length; i++ {
		arr[i-1] = arr[i]
	}

	return arr[0 : length-1]
}
