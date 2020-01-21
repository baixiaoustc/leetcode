/*
sort list, time complexity O(nlogn), space complexity O(1)
*/

package leetcode

func sortList(head *ListNode) *ListNode {
	length := 0
	p := head
	for p != nil {
		p = p.Next
		length++
	}

	dummy := new(ListNode)
	dummy.Next = head

	for size := 1; size < length; size *= 2 {
		cur := dummy.Next
		tail := dummy

		for cur != nil {
			left := cur
			right := cutList(cur, size)
			cur = cutList(right, size)

			tail.Next = mergeTwoLists(left, right)
			for tail.Next != nil {
				tail = tail.Next
			}
		}
	}

	return dummy.Next
}
