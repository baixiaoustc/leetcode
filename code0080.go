/*
给定一个排序数组，你需要在原地删除重复出现的元素，使得每个元素最多出现两次，返回移除后数组的新长度。

不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。

示例 1:

给定 nums = [1,1,1,2,2,3],

函数应返回新长度 length = 5, 并且原数组的前五个元素被修改为 1, 1, 2, 2, 3 。

你不需要考虑数组中超出新长度后面的元素。
示例 2:

给定 nums = [0,0,1,1,1,1,2,3,3],

函数应返回新长度 length = 7, 并且原数组的前五个元素被修改为 0, 0, 1, 1, 2, 3, 3 。

你不需要考虑数组中超出新长度后面的元素。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package leetcode

func removeDuplicatesMedium(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}

	count := 1
	for i := 1; i < length; i++ {
		if nums[i] == nums[i-1] {
			count++
			if count > 2 {
				nums = delArray(nums, i)
				i--
				count--
				length--
			}
		} else {
			count = 1
		}
	}

	return length
}

func removeDuplicatesMedium1(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}

	count := 1
	j := 1
	for i := 1; i < length; i++ {
		if nums[i] == nums[i-1] {
			count++
		} else {
			count = 1
		}

		if count <= 2 {
			nums[j] = nums[i]
			j++
		}
	}

	return j
}
