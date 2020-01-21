/*
反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。

说明:
1 ≤ m ≤ n ≤ 链表长度。

示例:

输入: 1->2->3->4->5->NULL, m = 2, n = 4
输出: 1->4->3->2->5->NULL

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-linked-list-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package leetcode

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil {
		return nil
	}
	if m > n || m < 1 {
		return nil
	}

	dummy := new(ListNode)
	dummy.Next = head

	tail := cutList(head, n)

	pre := dummy
	cur := head
	for i := 0; i < m-1; i++ {
		cur = cur.Next
		pre = pre.Next
		if cur == nil {
			break
		}
	}
	if cur == nil {
		return nil
	}

	next := cutList(pre, 1)
	next = reverseList(next)
	pre.Next = next

	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = tail

	return dummy.Next
}
