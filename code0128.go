/*
给定一个未排序的整数数组，找出最长连续序列的长度。

要求算法的时间复杂度为 O(n)。

示例:

输入: [100, 4, 200, 1, 3, 2]
输出: 4
解释: 最长连续序列是 [1, 2, 3, 4]。它的长度为 4。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-consecutive-sequence
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package leetcode

//并查集
//https://mp.weixin.qq.com/s/tjpBEe840_NLSqX-abY7-A
func longestConsecutive(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	m := make(map[int]int)
	for _, n := range nums {
		if _, ok := m[n]; !ok {
			m[n] = n
			if _, ok := m[n+1]; ok {
				m[n] = n + 1
			}
			if _, ok := m[n-1]; ok {
				m[n-1] = n
			}
		}
	}

	ret := 1
	for _, n := range nums {
		if n != m[n] {
			p := parent(m, n)
			if ret < p-n+1 {
				ret = p - n + 1
			}
		}
	}

	return ret
}

func parent(m map[int]int, n int) int {
	for {
		if v, ok := m[n]; ok && v == n {
			return n
		} else {
			n = v
		}
	}

	return -1
}
