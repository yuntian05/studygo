package main

import (
	"fmt"
	"golang.org/x/tools/container/intsets"
	_ "golang.org/x/tools/container/intsets"
	"studygo/muke/tree"
)

// 后序遍历
type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	leftNode := &myTreeNode{myNode.node.Left}
	leftNode.postOrder()

	rightNode := &myTreeNode{myNode.node.Right}
	rightNode.postOrder()
	myNode.node.Print()
}

func testSparse()  {
	s := intsets.Sparse{}
	s.Insert(1)
	s.Insert(100)
	fmt.Println(s.Has(1))
}

func main() {
	var root tree.Node
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{Value: 5}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)

	root.Traverse()
	fmt.Println()
	myRootNode := &myTreeNode{&root}
	myRootNode.postOrder()
	fmt.Println()

	testSparse()

	c := root.TravereWithChannel()
	maxNode := 0
	for node := range c {
		if node.Value > maxNode {
			maxNode = node.Value
		}
	}
	fmt.Print("max node = ", maxNode)
}
