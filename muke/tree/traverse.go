package tree

// 中序遍历
func (node *Node) Traverse() {
	node.TravereFunc(func(node *Node) {
		node.Print()
	})
}

func (node *Node) TravereFunc(f func (n *Node)) {
	if node == nil {
		return
	}
	node.Left.TravereFunc(f)
	f(node)
	node.Right.TravereFunc(f)
}

func (node *Node) TravereWithChannel() chan *Node {
	out := make(chan *Node)
	go func() {
		node.TravereFunc(func(n *Node) {
			out <- n
		})
		close(out)
	}()

	return out
}
