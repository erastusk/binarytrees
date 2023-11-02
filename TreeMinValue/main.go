package main

import (
	"fmt"
	"math"
)

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

var (
	stack     []*Node
	zero_node *Node
	pop_node  *Node
	status    bool
)

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

var min = math.MaxInt

func BR_Iterative(a *Node) int {
	// Queue
	stack = prependNode(stack, a)
	min = a.data
	for len(stack) > 0 {
		stack, pop_node = pop(stack)
		if pop_node.data < min {
			min = pop_node.data
		}
		if pop_node.left != nil {
			stack = prependNode(stack, pop_node.left)
		}
		if pop_node.right != nil {
			stack = prependNode(stack, pop_node.right)
		}
	}
	return min
}

var sum int

func BR_Recursive(a *Node) int {
	if a == zero_node {
		return 0
	}
	if a.data < min {
		min = a.data
	}
	BR_Recursive(a.left)
	BR_Recursive(a.right)
	return min
}

func prependNode(stack []*Node, a *Node) []*Node {
	return append([]*Node{a}, stack...)
}

func pop(stack []*Node) ([]*Node, *Node) {
	return stack[:len(stack)-1], stack[len(stack)-1]
}

func main() {
	a := CreateNodeTree()
	fmt.Println(BR_Iterative(a))
	// fmt.Println(DP_Iterative(a))
	fmt.Println(BR_Recursive(a))
}
