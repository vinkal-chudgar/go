//Program to count leaf nodes in a binary tree

package main

import (
	"fmt"
)

type BinaryTree struct {
	root *Node
}

func (bt *BinaryTree) CountLeafNodes() int {
	return bt.countLeafNodes(bt.root)
}
func (bt *BinaryTree) countLeafNodes(node *Node) int {
	if node == nil {
		return 0
	}
	if node.isLeafNode() {
		return 1
	}
	return bt.countLeafNodes(node.left) + bt.countLeafNodes(node.right)
}

type Node struct {
	left  *Node
	right *Node
	data  int
}

func (node *Node) isLeafNode() bool {
	if node.left == nil && node.right == nil {
		return true
	}
	return false
}

func main() {

	bt := BinaryTree{}
	bt.root = &Node{data: 1}
	bt.root.left = &Node{data: 2}
	bt.root.right = &Node{data: 3}
	bt.root.left.left = &Node{data: 4}
	bt.root.left.right = &Node{data: 5}

	fmt.Printf("count: %v\n", bt.CountLeafNodes())
}
