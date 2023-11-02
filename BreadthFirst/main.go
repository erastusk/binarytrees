package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

func AddNode(c int) *Node {
	return &Node{
		data: c,
	}
}

func BreadthFirstSearch_Iterative(a *Node) []int {
	// Uses a Queue to track values
	// Preorder Traversal
	// Expected output 0, 1, 2, 3, 4, 5, 6, 7
	var q []*Node
	var node *Node
	var res []int
	q = Prepend(a, q)
	for len(q) > 0 {
		q, node = Pop(q)
		res = append(res, node.data)
		if node.left != nil {
			q = Prepend(node.left, q)
		}
		if node.right != nil {
			q = Prepend(node.right, q)
		}
	}
	return res
}

var (
	res_rec   []int
	zero_node *Node
)

func BreadthFirstSearch_Recursion(a *Node) []int {
	// There isn't a way to implement a Breadth First Recursion see below
	if a == zero_node {
		return res_rec
	}
	res_rec = append(res_rec, a.data)
	BreadthFirstSearch_Recursion(a.left)
	BreadthFirstSearch_Recursion(a.right)
	return res_rec
}

func Prepend(a *Node, stack []*Node) []*Node {
	return append([]*Node{a}, stack...)
}

func Pop(stack []*Node) ([]*Node, *Node) {
	l := len(stack)
	return stack[:l-1], stack[l-1]
}

func CreateNodeTree() *Node {
	/*Depth First Definition
	  Go as deep as you can go then move laterally
		Expected output 0, 1, 2, 3, 4, 5, 6, 7
		  Tree Node

			           0
			      /        \
			      1         2
			   /    \     /   \
			  3      4   5     6
			/
		 7
	*/
	a := AddNode(0)
	b := AddNode(1)
	c := AddNode(2)
	d := AddNode(3)
	e := AddNode(4)
	f := AddNode(5)
	g := AddNode(6)
	h := AddNode(7)

	a.left = b
	a.right = c
	b.left = d
	b.right = e
	c.left = f
	c.right = g
	d.left = h
	return a
}

func main() {
	a := CreateNodeTree()
	result := BreadthFirstSearch_Iterative(a)
	fmt.Println(result)
	fmt.Println(BreadthFirstSearch_Recursion(a))
}
