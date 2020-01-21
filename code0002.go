/*
add two nonnegative number into one
*/

package leetcode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	head := new(ListNode)
	p := head
	add := false
	for l1 != nil && l2 != nil {
		tmp := new(ListNode)
		tmp.Val = l1.Val + l2.Val

		if add {
			tmp.Val++
			add = false
		}
		if tmp.Val > 9 {
			add = true
			tmp.Val -= 10
		}

		p.Next = tmp
		p = p.Next

		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		tmp := new(ListNode)
		tmp.Val = l1.Val

		if add {
			tmp.Val++
			add = false
		}
		if tmp.Val > 9 {
			add = true
			tmp.Val -= 10
		}

		p.Next = tmp
		p = p.Next

		l1 = l1.Next
	}

	for l2 != nil {
		tmp := new(ListNode)
		tmp.Val = l2.Val

		if add {
			tmp.Val++
			add = false
		}
		if tmp.Val > 9 {
			add = true
			tmp.Val -= 10
		}

		p.Next = tmp
		p = p.Next

		l2 = l2.Next
	}

	if add {
		tmp := new(ListNode)
		tmp.Val = 1
		p.Next = tmp
	}

	return head.Next
}
