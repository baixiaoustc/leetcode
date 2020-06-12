package leetcode

import (
	"fmt"
	"strconv"
)

//二叉树
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewBinaryTree(in []string) *TreeNode {
	fmt.Println("in", in, len(in))
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
	fmt.Printf("in %d\n", root.Val)
	for queue.Size() != 0 {
		ele, _ := queue.Peek()
		queue.Remove()
		node := ele.(*TreeNode)
		fmt.Println("---- i", i, "out", node.Val)

		leftIndex := 2*i + 1
		fmt.Println("leftIndex", leftIndex)
		if len(in) > leftIndex {
			n, is = parseNode(in[leftIndex])
			if is {
				left := &TreeNode{
					Val: n,
				}
				node.Left = left
				queue.Add(left)
				fmt.Printf("in %d; %d left is %d\n", left.Val, node.Val, left.Val)
			} else {
				fmt.Println("is nil")
			}
		}

		rightIndex := 2*i + 2
		fmt.Println("rightIndex", rightIndex)
		if len(in) > rightIndex {
			n, is = parseNode(in[rightIndex])
			if is {
				right := &TreeNode{
					Val: n,
				}
				node.Right = right
				queue.Add(right)
				fmt.Printf("in %d; %d right is %d\n", right.Val, node.Val, right.Val)
			} else {
				fmt.Println("is nil")
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

func ParseBinaryTree(root *TreeNode) []string {
	fmt.Println("======================")
	if root == nil {
		return []string{}
	}
	nodes := make([]string, 0)
	queue := NewLinkedQueue()
	queue.Add(root)
	nodes = append(nodes, fmt.Sprint(root.Val))
	for queue.Size() != 0 {
		size := queue.Size()
		for i := 0; i < size; i++ {
			ele, _ := queue.Peek()
			queue.Remove()
			node := ele.(*TreeNode)
			if node.Left != nil {
				queue.Add(node.Left)
				nodes = append(nodes, fmt.Sprint(node.Left.Val))
			} else {
				nodes = append(nodes, "null")
			}
			if node.Right != nil {
				queue.Add(node.Right)
				nodes = append(nodes, fmt.Sprint(node.Right.Val))
			} else {
				nodes = append(nodes, "null")
			}
		}
	}

	return nodes
}

//==============================================================================

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
