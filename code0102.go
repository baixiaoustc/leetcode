/*
102. 二叉树的层序遍历
给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。



示例：
二叉树：[3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层次遍历结果：

[
  [3],
  [9,20],
  [15,7]
]
*/

package leetcode

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	ret := make([][]int, 0)
	queue := NewLinkedQueue()
	queue.Add(root)
	for queue.Size() != 0 {
		n := queue.Size()
		tmp := make([]int, n)
		for i := 0; i < n; i++ {
			ele, _ := queue.Peek()
			node := ele.(*TreeNode)
			tmp[i] = node.Val
			if node.Left != nil {
				queue.Add(node.Left)
			}
			if node.Right != nil {
				queue.Add(node.Right)
			}
			queue.Remove()
		}
		ret = append(ret, tmp)
	}
	return ret
}
