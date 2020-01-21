/*
insertion sort of list
*/

package leetcode

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	length := 0
	p := head
	for p != nil {
		p = p.Next
		length++
	}

	dummy := new(ListNode)
	dummy.Next = head
	//dummy.Print()

	for size := 0; size < length; size++ {
		tail := findList(dummy, size)
		cur := tail.Next
		//cur.P()
		//dummy.Print()

		for h := dummy; h.Next != nil && h.Next != cur; h = h.Next {
			next := h.Next
			if next.Val <= cur.Val {
				continue
			} else {
				h.Next = cur
				tail.Next = cur.Next
				cur.Next = next
				break
			}
		}
		//dummy.Print()
	}

	return dummy.Next
}

func insertionSortListRecursive(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	p1 := head
	p2 := insertionSortListRecursive(head.Next)

	if p2 == nil {
		return p1
	}

	if p1.Val <= p2.Val {
		p1.Next = p2

		return p1
	} else {
		pre := p2
		cur := p2.Next

		for cur != nil && cur.Val < p1.Val {
			pre = cur
			cur = cur.Next
		}

		if cur == nil {
			pre.Next = p1
			p1.Next = nil
		} else {
			pre.Next = p1
			p1.Next = cur
		}

		return p2
	}
}
