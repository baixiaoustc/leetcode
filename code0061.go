package leetcode

func rotateRight(head *ListNode, k int) *ListNode {
	length := head.Length()
	mod := k / length
	m := k - mod*length

	if m == 0 {
		return head
	}

	n := length - m

	newHead := cutList(head, n)
	p := newHead
	for p.Next != nil {
		p = p.Next
	}
	p.Next = head
	return newHead
}
