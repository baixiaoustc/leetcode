/*
94. 二叉树的中序遍历
给定一个二叉树，返回它的中序 遍历。

示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [1,3,2]
进阶: 递归算法很简单，你可以通过迭代算法完成吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-inorder-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	ret := make([]int, 0)

	if root.Left != nil {
		ret = append(ret, inorderTraversal(root.Left)...)
	}
	ret = append(ret, root.Val)
	if root.Right != nil {
		ret = append(ret, inorderTraversal(root.Right)...)
	}

	return ret
}

func inorderTraversalDFS(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	ret := make([]int, 0)
	tmp := root

	stack := NewStackInterface()
	for tmp != nil || stack.Size() != 0 {
		//把左树全部入栈
		for tmp != nil {
			stack.Push(tmp)
			tmp = tmp.Left
		}

		ele := stack.Pop().(*TreeNode)
		ret = append(ret, ele.Val)
		if ele.Right != nil {
			tmp = ele.Right
		}
	}

	return ret
}
