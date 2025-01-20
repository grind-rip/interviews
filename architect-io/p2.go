package main

import (
	"fmt"
	"math/rand"

	"github.com/google/uuid"
)

// Node is a node in the tree.
type Node struct {
	// id is a random UUID.
	id uuid.UUID
	// score is a random non-negative integer from 0 to 99 inclusive.
	score int
	// children is a list of pointers to the node's children.
	children []*Node
}

// String prints the node's score.
func (n *Node) String() string {
	return fmt.Sprintf("Node(score: %d)", n.score)
}

// NewNode creates a new node and returns a pointer to it.
func NewNode() *Node {
	node := new(Node)
	node.id = uuid.New()
	node.score = rand.Intn(100)
	node.children = []*Node{}
	return node
}

// createTreeHelper creates a tree of the specified depth by recursively adding
// three children to each node until the desired depth is reached.
//
// The depth (or height) of a tree is the length of the longest path from the
// root node to any leaf node in the tree. A tree with a depth of 0 is a single
// node, the root.
func createTreeHelper(parent *Node, depth int) *Node {
	if depth <= 1 {
		return parent
	}
	// Each descendant of parent has three children.
	for i := 0; i < 3; i++ {
		parent.children = append(parent.children, createTree(depth-1))
	}
	return parent
}

// createTree creates a new tree of depth `depth`.
//
// Example:
//
//	   ●
//	 / | \
//	●  ●  ●  Depth = 1
func createTree(depth int) *Node {
	root := NewNode()
	return createTreeHelper(root, depth)
}

// findBiggestScore finds the biggest score of all the nodes in a list of
// trees.
//
// NOTE: Since trees are by definition acyclic, we do not need to maintain a
// list of visited nodes.
func findBiggestScore(trees []*Node) int {
	biggestScore := 0
	for _, root := range trees {
		score := findBiggestScoreHelper(root, 0)
		if score > biggestScore {
			biggestScore = score
		}
	}
	return biggestScore
}

// findBiggestScoreHelper uses depth-first search to find the biggest score of
// all the nodes in a tree.
//
// When using depth-first search, the biggest score is carried throughout the
// tree traversal. Effectively, the biggest score is found for each subtree.
func findBiggestScoreHelper(node *Node, biggestScore int) int {
	if node == nil {
		return biggestScore
	}

	if node.score > biggestScore {
		biggestScore = node.score
	}

	for _, child := range node.children {
		score := findBiggestScoreHelper(child, biggestScore)
		if score > biggestScore {
			biggestScore = score
		}
	}

	return biggestScore
}
