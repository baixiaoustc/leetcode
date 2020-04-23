/*
删除链表中等于给定值 val 的所有节点。

示例:

输入: 1->2->6->3->4->5->6, val = 6
输出: 1->2->3->4->5
*/

package leetcode

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	dummy := new(ListNode)
	dummy.Next = head

	pre := dummy
	for head != nil {
		if head.Val == val {
			pre.Next = head.Next
		} else {
			pre = pre.Next
		}
		head = head.Next
	}

	return dummy.Next
}
