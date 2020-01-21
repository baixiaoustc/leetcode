/*
swap adjacent pairs nodes
*/

package leetcode

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummy := new(ListNode)
	dummy.Next = head

	pre := dummy
	cur := head
	next := cur.Next
	for cur != nil && next != nil {
		cur.Next = next.Next
		pre.Next = next
		next.Next = cur

		pre = cur
		cur = cur.Next
		if cur == nil {
			break
		}
		next = cur.Next
	}

	return dummy.Next
}
