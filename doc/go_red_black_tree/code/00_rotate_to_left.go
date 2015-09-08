package main

import "fmt"

// Tree contains a Root node of a binary search tree.
type Tree struct {
	Root *Node
}

// New returns a new Tree with its root Node.
func New(root *Node) *Tree {
	tr := &Tree{}
	root.Black = true
	tr.Root = root
	return tr
}

// Interface represents a single object in the tree.
type Interface interface {
	// Less returns true when the receiver item(key)
	// is less than the given(than) argument.
	Less(than Interface) bool
}

// Node is a Node and a Tree itself.
type Node struct {
	// Left is a left child Node.
	Left *Node

	Key   Interface
	Black bool // True when the color of parent link is black.
	// In Left-Leaning Red-Black tree, new nodes are always red
	// because the zero boolean value is false.
	// Null links are black.

	// Right is a right child Node.
	Right *Node
}

// NewNode returns a new Node.
func NewNode(key Interface) *Node {
	nd := &Node{}
	nd.Key = key
	nd.Black = false
	return nd
}

func (tr *Tree) String() string {
	return tr.Root.String()
}

func (nd *Node) String() string {
	if nd == nil {
		return "[]"
	}
	s := ""
	if nd.Left != nil {
		s += nd.Left.String() + " "
	}
	s += fmt.Sprintf("%v(%v)", nd.Key, nd.Black)
	if nd.Right != nil {
		s += " " + nd.Right.String()
	}
	return "[" + s + "]"
}

func isRed(nd *Node) bool {
	if nd == nil {
		return false
	}
	return !nd.Black
}

// insert inserts nd2 with nd1 as a root.
func (nd1 *Node) insert(nd2 *Node) *Node {
	if nd1 == nil {
		return nd2
	}
	if nd1.Key.Less(nd2.Key) {
		// nd1 is smaller than nd2
		// nd1 < nd2
		nd1.Right = nd1.Right.insert(nd2)
	} else {
		// nd1 is greater than nd2
		// nd1 >= nd2
		nd1.Left = nd1.Left.insert(nd2)
	}
	// Balance from nd1
	return Balance(nd1)
}

// Insert inserts a Node to a Tree without replacement.
// It does standard BST insert and colors the new link red.
// If the new red link is a right link, rotate left.
// If two left red links in a row, rotate to right and flip color.
// (https://youtu.be/lKmLBOJXZHI?t=20m43s)
//
// Note that it recursively balances from its parent nodes
// to the root node at the top.
//
// And make sure paint the Root black(not-red).
func (tr *Tree) Insert(nd *Node) {
	if tr.Root == nd {
		return
	}
	tr.Root = tr.Root.insert(nd)

	// Root node must be always black.
	tr.Root.Black = true
}

// RotateToLeft runs when there is a right-leaning link.
// tr.Root = RotateToLeft(tr.Root) overwrite the Root
// with the new top Node.
func RotateToLeft(nd *Node) *Node {
	if nd.Right.Black {
		panic("Can't rotate a black link")
	}

	// exchange x and nd
	// nd is parent node, x is Right child
	x := nd.Right
	nd.Right = x.Left
	x.Left = nd

	x.Black = nd.Black
	nd.Black = false

	return x
}

// RotateToRight runs when there are two left red links in a row.
// tr.Root = RotateToRight(tr.Root) overwrite the Root
// with the new top Node.
func RotateToRight(nd *Node) *Node {
	if nd.Left.Black {
		panic("Can't rotate a black link")
	}

	// exchange x and nd
	// nd is parent node, x is Left child
	x := nd.Left
	nd.Left = x.Right
	x.Right = nd

	x.Black = nd.Black
	nd.Black = false

	return x
}

// FlipColor flips the color.
// Left and Right children must be present
func FlipColor(nd *Node) {
	// nd is parent node
	nd.Black = !nd.Black
	nd.Left.Black = !nd.Left.Black
	nd.Right.Black = !nd.Right.Black
}

// Balance balances the Node.
func Balance(nd *Node) *Node {
	// nd is parent node
	if isRed(nd.Right) && !isRed(nd.Left) {
		nd = RotateToLeft(nd)
	}
	if isRed(nd.Left) && isRed(nd.Left.Left) {
		nd = RotateToRight(nd)
	}
	if isRed(nd.Left) && isRed(nd.Right) {
		FlipColor(nd)
	}
	return nd
}

type Float64 float64

// Less returns true if float64(a) < float64(b).
func (a Float64) Less(b Interface) bool {
	return a < b.(Float64)
}

func main() {
	node3 := NewNode(Float64(3))
	node3.Black = true

	node1 := NewNode(Float64(1))
	node1.Black = true

	node13 := NewNode(Float64(13))
	node13.Black = false

	node9 := NewNode(Float64(9))
	node9.Black = true

	node17 := NewNode(Float64(17))
	node17.Black = true

	tr := New(node3)
	tr.Root.Right = node13
	tr.Root.Right.Left = node9
	tr.Root.Right.Right = node17
	tr.Root.Left = node1
	/*
	        3(B)
	      /      \
	   1(B)      13(R)
	            /   \
	         9(B)  17(B)
	*/
	fmt.Println("Before tr.Root = RotateToLeft(tr.Root)")
	fmt.Println(tr.Root.Left)
	fmt.Println(tr.Root)
	fmt.Println(tr.Root.Right)

	tr.Root = RotateToLeft(tr.Root)
	/*
			   	   13(B)
			   	  /     \
			   3(R)     17(B)
			  /   \
		   1(B)   9(B)
	*/

	fmt.Println("After tr.Root = RotateToLeft(tr.Root)")
	fmt.Println(tr.Root.Left)
	fmt.Println(tr.Root)
	fmt.Println(tr.Root.Right)
	// Output:
	// Before tr.Root = RotateToLeft(tr.Root)
	// [1(true)]
	// [[1(true)] 3(true) [[9(true)] 13(false) [17(true)]]]
	// [[9(true)] 13(false) [17(true)]]
	// After tr.Root = RotateToLeft(tr.Root)
	// [[1(true)] 3(false) [9(true)]]
	// [[[1(true)] 3(false) [9(true)]] 13(true) [17(true)]]
	// [17(true)]
}
