package main

import "fmt"

type BinaryTree struct {
	root *Node
}

type Node struct {
	left  *Node
	right *Node
	data  int64
}

func (bt *BinaryTree) Insert(data int64) {
	if bt.root == nil {
		bt.root = &Node{data: data}
	} else {
		bt.root.Insert(data)
	}
}

func (n *Node) Insert(data int64) {
	if data <= n.data {
		if n.left == nil {
			n.left = &Node{data: data}
		} else {
			n.left.Insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &Node{data: data}
		} else {
			n.right.Insert(data)
		}
	}
}

func (bt *BinaryTree) InOrder(n *Node, printf func(*Node)) {
	if n == nil {
		return
	}
	bt.InOrder(n.left, printf)
	printf(n)
	bt.InOrder(n.right, printf)

}

func (bt *BinaryTree) Mirror() {
	bt.root = bt.mirror(bt.root)
}

func (bt *BinaryTree) mirror(node *Node) *Node {
	if node == nil {
		return node
	}
	left := bt.mirror(node.left)
	right := bt.mirror(node.right)

	node.left = right
	node.right = left

	return node
}

func main() {
	var tree BinaryTree
	tree.Insert(50)
	tree.Insert(60)
	tree.Insert(40)

	fmt.Println("Original tree: ")
	tree.InOrder(tree.root, func(n *Node) {
		fmt.Println(n.data)
	})

	tree.Mirror()
	fmt.Println("Mirrored tree: ")
	tree.InOrder(tree.root, func(n *Node) {
		fmt.Println(n.data)
	})
}
