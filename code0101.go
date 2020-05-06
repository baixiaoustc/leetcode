/*
101. 对称二叉树
给定一个二叉树，检查它是否是镜像对称的。



例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

    1
   / \
  2   2
 / \ / \
3  4 4  3


但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:

    1
   / \
  2   2
   \   \
   3    3


进阶：

你可以运用递归和迭代两种方法解决这个问题吗？
*/

package leetcode

import (
	"fmt"
	"strings"
)

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	left := symmeticRecursive(root.Left, 1)
	right := symmeticRecursive(root.Right, 1)
	if len(left) != len(right) {
		return false
	}
	fmt.Println(left)
	fmt.Println(right)

	leftlist := strings.Split(left, "@")
	rightlist := strings.Split(right, "@")
	for i := 0; i < len(leftlist); i++ {
		if leftlist[i] != rightlist[len(leftlist)-1-i] {
			return false
		}
	}
	return true
}

func symmeticRecursive(root *TreeNode, depth int) string {
	if root == nil {
		return ""
	}

	left := symmeticRecursive(root.Left, depth+1)
	right := symmeticRecursive(root.Right, depth+1)
	return left + fmt.Sprintf("@%d.%d@", depth, root.Val) + right
}

func isSymmetricBFS(root *TreeNode) bool {
	if root == nil {
		return true
	}

	queue := NewLinkedQueue()
	queue.Add(root)
	queue.Add(root)
	for queue.Size() != 0 {
		ele, _ := queue.Peek()
		queue.Remove()
		left := ele.(*TreeNode)
		ele, _ = queue.Peek()
		queue.Remove()
		right := ele.(*TreeNode)
		if left == nil && right == nil {
			continue
		}
		if left == nil || right == nil {
			return false
		}
		if left.Val != right.Val {
			return false
		}
		queue.Add(left.Right)
		queue.Add(right.Left)
		queue.Add(left.Left)
		queue.Add(right.Right)
	}

	return true
}
