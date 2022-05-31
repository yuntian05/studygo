package main

import "studygo/muke/tree"

func main() {
	var root tree.Node
	root = tree.Node{Value: 3}
	root.Left = tree.CreateNode(9)
	root.Right = tree.CreateNode(20)
	root.Right.Left = tree.CreateNode(15)
	root.Right.Right = tree.CreateNode(7)
	minDepth(&root)
}

// 二叉树的最小深度
func minDepth(root *tree.Node) int {
	if root == nil {
		return 0
	}

	queue := []*tree.Node{root}
	depth := 1
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left == nil && node.Right == nil {
				// 叶子节点
				return depth
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		depth++
	}
	return depth
}
