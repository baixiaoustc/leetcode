/*
给定一个链表，旋转链表，将链表每个节点向右移动 k 个位置，其中 k 是非负数。

示例 1:

输入: 1->2->3->4->5->NULL, k = 2
输出: 4->5->1->2->3->NULL
解释:
向右旋转 1 步: 5->1->2->3->4->NULL
向右旋转 2 步: 4->5->1->2->3->NULL
示例 2:

输入: 0->1->2->NULL, k = 4
输出: 2->0->1->NULL
解释:
向右旋转 1 步: 2->0->1->NULL
向右旋转 2 步: 1->2->0->NULL
向右旋转 3 步: 0->1->2->NULL
向右旋转 4 步: 2->0->1->NULL

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/rotate-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package leetcode

func rotateRight(head *ListNode, k int) *ListNode {
	length := head.Length()
	mod := k / length
	m := k - mod*length

	if m == 0 {
		return head
	}

	n := length - m

	newHead := cutList(head, n)
	p := newHead
	for p.Next != nil {
		p = p.Next
	}
	p.Next = head
	return newHead
}

/*
给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。

示例 1:

输入: [1,2,3,4,5,6,7] 和 k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右旋转 1 步: [7,1,2,3,4,5,6]
向右旋转 2 步: [6,7,1,2,3,4,5]
向右旋转 3 步: [5,6,7,1,2,3,4]
示例 2:

输入: [-1,-100,3,99] 和 k = 2
输出: [3,99,-1,-100]
解释:
向右旋转 1 步: [99,-1,-100,3]
向右旋转 2 步: [3,99,-1,-100]
说明:

尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。
要求使用空间复杂度为 O(1) 的 原地 算法。
*/
func rotate(nums []int, k int) {
	length := len(nums)
	mod := k / length
	m := k - mod*length
	l := length - m

	if m == 0 {
		return
	}

	tmps := make([]int, m)
	for i := 0; i < m; i++ {
		tmps[i] = nums[i+l]
	}

	for i := l - 1; i >= 0; i-- {
		nums[m+i] = nums[i]
	}

	for i := 0; i < m; i++ {
		nums[i] = tmps[i]
	}

	return
}
