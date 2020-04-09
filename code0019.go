/*
给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

示例：

给定一个链表: 1->2->3->4->5, 和 n = 2.

当删除了倒数第二个节点后，链表变为 1->2->3->5.
说明：

给定的 n 保证是有效的。

进阶：

你能尝试使用一趟扫描实现吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package leetcode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || n < 1 {
		return nil
	}

	length := 0
	p := head
	for p != nil {
		p = p.Next
		length++
	}

	m := length - n
	if m == 0 {
		return cutList(head, 1)
	} else if m < 0 {
		return nil
	} else {
		cur := head
		next := cur.Next
		for ; m > 1; m-- {
			cur = next
			next = cur.Next
		}

		cur.Next = next.Next
		next.Next = nil
		return head
	}
}

func removeNthFromEndOneScan(head *ListNode, n int) *ListNode {
	if head == nil || n < 1 {
		return nil
	}
	dummy := new(ListNode)
	dummy.Next = head

	left, right := head, head

	for m := n + 1; m > 0; {
		right = right.Next
		m--
		if right == nil {
			if m == 0 {
				if left.Next.Next != nil {
					left.Next = left.Next.Next
				} else {
					left.Next = nil
				}
			} else if m == 1 {
				//kill the left node
				if left.Next != nil {
					dummy.Next = left.Next
				} else {
					dummy.Next = nil
				}
			}
			return dummy.Next
		}
	}

	for right != nil {
		right = right.Next
		left = left.Next
	}

	left.Next = left.Next.Next
	return dummy.Next
}

func removeNthFromEndOneScan1(head *ListNode, n int) *ListNode {
	if head == nil || n < 1 {
		return nil
	}

	dummy := new(ListNode)
	dummy.Next = head

	fast := head
	for skip := n; skip > 0; skip-- {
		if fast.Next == nil {
			if skip == 1 {
				//special
				dummy.Next = dummy.Next.Next
				return dummy.Next
			} else {
				return nil
			}
		}
		fast = fast.Next
	}

	slow := head
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return dummy.Next
}
