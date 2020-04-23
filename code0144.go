/*
144. 二叉树的前序遍历
给定一个二叉树，返回它的 前序 遍历。

 示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [1,2,3]
进阶: 递归算法很简单，你可以通过迭代算法完成吗？
*/

package leetcode

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	ret := make([]int, 0)
	stack := NewStackInterface()
	stack.Push(root)
	for stack.Size() != 0 {
		ele := stack.Pop().(*TreeNode)
		ret = append(ret, ele.Val)
		if ele.Right != nil {
			stack.Push(ele.Right)
		}
		if ele.Left != nil {
			stack.Push(ele.Left)
		}
	}
	return ret
}

func preorderTraversalRecursive(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	ret := make([]int, 0)
	ret = append(ret, root.Val)
	left := preorderTraversalRecursive(root.Left)
	ret = append(ret, left...)
	right := preorderTraversalRecursive(root.Right)
	ret = append(ret, right...)
	return ret
}
