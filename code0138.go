/*
138. 复制带随机指针的链表
给定一个链表，每个节点包含一个额外增加的随机指针，该指针可以指向链表中的任何节点或空节点。

要求返回这个链表的 深拷贝。

我们用一个由 n 个节点组成的链表来表示输入/输出中的链表。每个节点用一个 [val, random_index] 表示：

val：一个表示 Node0138.val 的整数。
random_index：随机指针指向的节点索引（范围从 0 到 n-1）；如果不指向任何节点，则为  null 。


示例 1：



输入：head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
输出：[[7,null],[13,0],[11,4],[10,2],[1,0]]
示例 2：



输入：head = [[1,1],[2,1]]
输出：[[1,1],[2,1]]
示例 3：



输入：head = [[3,null],[3,0],[3,null]]
输出：[[3,null],[3,0],[3,null]]
示例 4：

输入：head = []
输出：[]
解释：给定的链表为空（空指针），因此返回 null。


提示：

-10000 <= Node0138.val <= 10000
Node0138.random 为空（null）或指向链表中的节点。
节点数目不超过 1000 。
*/

package leetcode

type Node0138 struct {
	Val    int
	Next   *Node0138
	Random *Node0138
}

//拷贝的意思是每当遇到一个新的未访问过的节点，你都需要创造一个新的节点。
func copyRandomList(head *Node0138) *Node0138 {
	if head == nil {
		return nil
	}

	visit := make(map[*Node0138]*Node0138)

	newHead := new(Node0138)
	newHead.Val = head.Val
	visit[head] = newHead
	dummy := new(Node0138)
	dummy.Next = newHead

	newNode := func(node *Node0138) *Node0138 {
		if node == nil {
			return nil
		}
		if v, ok := visit[node]; ok {
			return v
		}
		tmp := new(Node0138)
		tmp.Val = node.Val
		visit[node] = tmp
		return tmp
	}

	for head != nil {
		newHead.Random = newNode(head.Random)
		newHead.Next = newNode(head.Next)

		head = head.Next
		newHead = newHead.Next
	}

	return dummy.Next
}

//Old List: A --> B --> C --> D
//InterWeaved List: A --> A' --> B --> B' --> C --> C' --> D --> D'
func copyRandomList1(head *Node0138) *Node0138 {
	if head == nil {
		return nil
	}

	dummy := new(Node0138)
	dummy.Next = head

	//1. 将复制节点添加到原节点后面
	for head != nil {
		tmp := new(Node0138)
		tmp.Val = head.Val
		tmp.Next = head.Next
		head.Next = tmp
		head = head.Next.Next
	}

	//2. 复制random节点
	head = dummy.Next
	for head != nil {
		if head.Random != nil {
			head.Next.Random = head.Random.Next
		}
		head = head.Next.Next
	}

	//3. 分离链表
	head = dummy.Next
	newHead := head.Next
	tmp := newHead
	for head != nil {
		head.Next = tmp.Next
		if tmp.Next != nil {
			tmp.Next = tmp.Next.Next
		}
		head = head.Next
		tmp = tmp.Next
	}

	return newHead
}
