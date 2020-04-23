/*
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
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
