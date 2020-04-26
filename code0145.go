/*
145. 二叉树的后序遍历
给定一个二叉树，返回它的 后序 遍历。

示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [3,2,1]
进阶: 递归算法很简单，你可以通过迭代算法完成吗？
*/

package leetcode

//取巧解法
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	ret := make([]int, 0)
	stack := NewStackInterface()
	stack.Push(root)
	for stack.Size() != 0 {
		ele := stack.Pop().(*TreeNode)
		ret = append(ret, ele.Val)
		if ele.Left != nil {
			stack.Push(ele.Left)
		}
		if ele.Right != nil {
			stack.Push(ele.Right)
		}
	}
	realRet := make([]int, len(ret))
	for i := range ret {
		realRet[i] = ret[len(ret)-1-i]
	}
	return realRet
}
