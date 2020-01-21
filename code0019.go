/*
delete the n th node from the end
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
