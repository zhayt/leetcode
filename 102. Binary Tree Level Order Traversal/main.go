package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	levels := make([][]int, 0)
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		level := make([]int, 0, len(queue))
		length := len(queue)
		for _, node := range queue {
			if node != nil {
				level = append(level, node.Val)
				queue = append(queue, node.Left, node.Right)
			}
		}

		if len(level) != 0 {
			levels = append(levels, level)
		}

		queue = queue[length:]
	}

	return levels
}

func main() {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}
	fmt.Println(levelOrder(root))
}
