package leetcode

import (
	"strconv"
)

//二叉树
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewBinaryTree(in []string) *TreeNode {
	if len(in) == 0 {
		return nil
	}

	n, is := parseNode(in[0])
	if !is {
		return nil
	}
	root := &TreeNode{
		Val: n,
	}
	i := 0
	queue := NewLinkedQueue()
	queue.Add(root)
	for queue.Size() != 0 {
		ele, _ := queue.Peek()
		queue.Remove()
		node := ele.(*TreeNode)
		//fmt.Println("out", node.Val)
		if len(in) > 2*i+1 {
			n, is = parseNode(in[2*i+1])
			if is {
				left := &TreeNode{
					Val: n,
				}
				node.Left = left
				queue.Add(left)
				//fmt.Println("in", left.Val)
			}
		}
		if len(in) > 2*i+2 {
			n, is = parseNode(in[2*i+2])
			if is {
				right := &TreeNode{
					Val: n,
				}
				node.Right = right
				queue.Add(right)
				//fmt.Println("in", right.Val)
			}
		}
		i++
	}

	return root
}

//返回value，是否有值
func parseNode(s string) (int, bool) {
	if s == "null" {
		return 0, false
	}

	n, _ := strconv.ParseInt(s, 10, 64)
	return int(n), true
}

//完美二叉树
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func NewPerfectTree(in []string) *Node {
	if len(in) == 0 {
		return nil
	}

	n, is := parseNode(in[0])
	if !is {
		return nil
	}
	root := &Node{
		Val: n,
	}
	i := 0
	queue := NewLinkedQueue()
	queue.Add(root)
	for queue.Size() != 0 {
		ele, _ := queue.Peek()
		queue.Remove()
		node := ele.(*Node)
		//fmt.Println("out", node.Val)
		if len(in) > 2*i+1 {
			n, is = parseNode(in[2*i+1])
			if is {
				left := &Node{
					Val: n,
				}
				node.Left = left
				queue.Add(left)
				//fmt.Println("in", left.Val)
			}
		}
		if len(in) > 2*i+2 {
			n, is = parseNode(in[2*i+2])
			if is {
				right := &Node{
					Val: n,
				}
				node.Right = right
				queue.Add(right)
				//fmt.Println("in", right.Val)
			}
		}
		i++
	}

	return root
}
