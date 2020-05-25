/*
105. 从前序与中序遍历序列构造二叉树
根据一棵树的前序遍历与中序遍历构造二叉树。

注意:
你可以假设树中没有重复的元素。

例如，给出

前序遍历 preorder = [3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]
返回如下的二叉树：

    3
   / \
  9  20
    /  \
   15   7
*/

package leetcode

import "fmt"

func buildTree0105(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	return buildTree0105Recursive(preorder, inorder)
}

func buildTree0105Recursive(preorder []int, inorder []int) *TreeNode {
	fmt.Println(preorder, inorder)
	if len(preorder) == 0 {
		return nil
	}

	val := preorder[0]
	root := &TreeNode{Val: val}
	for i := range inorder {
		if val == inorder[i] {
			fmt.Println(i)
			root.Left = buildTree0105Recursive(preorder[1:i+1], inorder[0:i])
			root.Right = buildTree0105Recursive(preorder[i+1:len(preorder)], inorder[i+1:len(preorder)])
		}
	}
	return root
}
