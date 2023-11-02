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

func CreateNodeTree() *Node {
	/*
		  Tree Node

			           10
			      /        \
			      1         2
			   /    \     /   \
			  3      4   5     6
			/
		 7
	*/
	a := AddNode(10)
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

// Given a Tree and include value, return a boolean if include value exists in the tree.
// We'll use Both search aligorithms Iteratively and Recursively
var (
	stack     []*Node
	zero_node *Node
	pop_node  *Node
	status    bool
)

func DP_Iterative(a *Node, inc int) bool {
	// Stack
	// push root Node
	// pop first Node
	// push first node left and righ children nodes into stack and so on
	stack = append(stack, a)
	for len(stack) > 0 {
		stack, pop_node = pop(stack)
		if pop_node.data == inc {
			return true
		}
		if pop_node.left != zero_node {
			stack = append(stack, pop_node.left)
		}
		if pop_node.right != zero_node {
			stack = append(stack, pop_node.right)
		}
	}

	return false
}

func DP_Recursive(a *Node, inc int) bool {
	if a == zero_node {
		return status
	}
	if a.data == inc {
		status = true
		return status
	}
	DP_Recursive(a.left, inc)
	DP_Recursive(a.right, inc)

	return status
}

func pop(st []*Node) ([]*Node, *Node) {
	return st[:len(st)-1], st[len(st)-1]
}

func BR_Iterative(root *Node, inc int) bool {
	// Queue - FIFO
	// Add (Prepend) root Node from the back
	// Remove Node in the front
	// Prepend left and right nodes to the back of Queue
	// Do this until there are no nodes in Queue
	stack = PrependNode(root, []*Node{})
	for len(stack) > 0 {
		stack, pop_node = PopNode(stack)
		if pop_node.data == inc {
			return true
		}
		if pop_node.left != nil {
			stack = PrependNode(pop_node.left, stack)
		}

		if pop_node.right != nil {
			stack = PrependNode(pop_node.right, stack)
		}
	}
	return false
}

func PrependNode(a *Node, stack []*Node) []*Node {
	stack = append([]*Node{a}, stack...)
	return stack
}

func PopNode(stack []*Node) ([]*Node, *Node) {
	return stack[:len(stack)-1], stack[len(stack)-1]
}

func BR_Recursive(a *Node, inc int) bool {
	if a == zero_node {
		return status
	}
	if a.data == inc {
		status = true
		return status
	}
	BR_Recursive(a.left, inc)
	BR_Recursive(a.right, inc)
	return status
}

func main() {
	a := CreateNodeTree()
	fmt.Println(DP_Iterative(a, 19))
	fmt.Println(DP_Recursive(a, 19))
	fmt.Println(BR_Iterative(a, 1))
	fmt.Println(BR_Recursive(a, 19))
}
