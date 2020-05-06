/*
106. 从中序与后序遍历序列构造二叉树
根据一棵树的中序遍历与后序遍历构造二叉树。

注意:
你可以假设树中没有重复的元素。

例如，给出

中序遍历 inorder = [9,3,15,20,7]
后序遍历 postorder = [9,15,7,20,3]
返回如下的二叉树：

    3
   / \
  9  20
    /  \
   15   7
*/

package leetcode

import "fmt"

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	return buildTreeRecursive(inorder, postorder)
}

//ir表示root的index
//is，ie表示root开始的树在中序数组中的左右index
func buildTreeRecursive(inorder []int, postorder []int) *TreeNode {
	fmt.Println(inorder, postorder)
	if len(inorder) == 0 {
		return nil
	}

	val := postorder[len(postorder)-1]
	root := &TreeNode{Val: val}
	var index int
	for i := range inorder {
		if val == inorder[i] {
			index = i
			break
		}
	}
	fmt.Println(index)
	root.Left = buildTreeRecursive(inorder[0:index], postorder[0:index])
	root.Right = buildTreeRecursive(inorder[index+1:], postorder[index:len(postorder)-1])
	return root
}
