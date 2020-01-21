package leetcode

/*
给定一个链表和一个特定值 x，对链表进行分隔，使得所有小于 x 的节点都在大于或等于 x 的节点之前。

你应当保留两个分区中每个节点的初始相对位置。

示例:

输入: head = 1->4->3->2->5->2, x = 3
输出: 1->2->2->4->3->5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/partition-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}
	dummy := new(ListNode)
	dummy.Next = head

	pre := dummy
	cur := head
	target := cur

	for cur != nil {
		if cur.Val >= x {
			break
		}
		cur = cur.Next
		target = cur
		pre = pre.Next
	}
	targetPre := pre

	for cur != nil {
		tmp := cur
		cur = cur.Next
		if tmp.Val < x {
			pre.Next = cur
			targetPre.Next = tmp
			targetPre = targetPre.Next
			tmp.Next = target
		} else {
			pre = pre.Next
		}
	}

	return dummy.Next
}
