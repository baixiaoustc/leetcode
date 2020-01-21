/*
scan a list, delete the nodes which duplicates with others
*/

package leetcode

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummy := new(ListNode)
	dummy.Next = head

	pre := dummy
	cur := head
	next := cur.Next
	for cur != nil && next != nil {
		found := false
		for next != nil && next.Val == cur.Val {
			found = true
			next = next.Next
		}

		if found {
			pre.Next = next
		} else {
			pre = cur
		}
		cur = next
		if cur == nil {
			break
		}
		next = cur.Next
	}

	return dummy.Next
}

//another eazy problem
func deleteDuplicates83(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummy := new(ListNode)
	dummy.Next = head

	cur := head
	next := cur.Next
	for cur != nil && next != nil {
		found := false
		for next != nil && next.Val == cur.Val {
			found = true
			next = next.Next
		}

		if found {
			cur.Next = next
		}
		cur = next
		if cur == nil {
			break
		}
		next = cur.Next
	}

	return dummy.Next
}
