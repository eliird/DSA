package main

import "fmt"

//Balanced Tree is when all the leaves are on the same level
type NodeT struct {
	data   int
	parent *NodeT
	left   *NodeT
	right  *NodeT
}

type Tree struct {
	head *NodeT
	size int
}

func (t *Tree) add(value int) {

}

func (t *Tree) remove(value int) {

}

func recurse(node *NodeT) {
	fmt.Println(node.data)
	if node.left == nil {
		return
	}
	recurse(node.left)

	if node.right == nil {
		return
	}
	recurse(node.right)

}

func (t *Tree) DFS() {
	//DFS preserves the shape of the traversal
	recurse(t.head)
}

func bfs_traverse(node *NodeT) {
	//put the children of the node on the queue
	//pop the first item in the queue
	//Put the children of next element on the q
}
func (t *Tree) BFS() {
	bfs_traverse(t.head)

}

func compareTrees(a *NodeT, b *NodeT) bool {

	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if a.data != b.data {
		return false
	}

	return compareTrees(a.right, b.right) && compareTrees(a.left, b.left)
}

//Also implement a code to see if two trees are equal
func (t1 *Tree) equals(t2 *Tree) bool {
	return compareTrees(t1.head, t2.head)
}

//Binary Seach Trees
//Try and Implement binary tree by your self
//Implement find in a tree, insert in a BST, remove from BST

//Remove is very triky. There are three cases
//Case 1: If the node to be removed is leaf
//Case II:If there is only 1 child
//Case III: If there are 2 children
// For case 3 we replace the enode to be deleted with the smallesst value on the rght tree or largest from the left tree and replace the node with that node and delete the bottom
