/*
请判断一个链表是否为回文链表。

示例 1:

输入: 1->2
输出: false
示例 2:

输入: 1->2->2->1
输出: true
进阶：
你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/palindrome-linked-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package leetcode

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}

	fast, slow := head, head
	pre, tmp := new(ListNode), new(ListNode)
	for fast != nil && fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
		pre.Next = tmp
		tmp = pre
	}

	if fast != nil {
		//说明是奇数个节点
		slow = slow.Next
	}

	for slow != nil {
		if slow.Val != pre.Val {
			return false
		}

		slow = slow.Next
		pre = pre.Next
	}

	return true
}
