package main

type Node struct {
	data  int
	left  *Node
	right *Node
}

var root *Node

func AddNode(c int) *Node {
	return &Node{
		data: c,
	}
}

func DepthFirstSearch(a *Node) []int {
	var dp_slice []*Node
	var res []int
	var p *Node
	dp_slice = Push(a, dp_slice)

	for len(dp_slice) > 0 {
		dp_slice, p = Pop(dp_slice)
		res = append(res, p.data)
		if p.right != nil {
			dp_slice = Push(p.right, dp_slice)
		}
		if p.left != nil {
			dp_slice = Push(p.left, dp_slice)
		}
	}
	return res
}

var (
	res_rec   []int
	zero_node *Node
)

func DepthFirstSearch_Recursion(root *Node) []int {
	if root == zero_node {
		return res_rec
	}
	res_rec = append(res_rec, root.data)
	DepthFirstSearch_Recursion(root.left)
	DepthFirstSearch_Recursion(root.right)
	return res_rec
}

func Push(a *Node, stack []*Node) []*Node {
	return append(stack, a)
}

func Pop(stack []*Node) ([]*Node, *Node) {
	l := len(stack)
	return stack[:l-1], stack[l-1]
}

func CreateNodeTree() *Node {
	/*  Tree Node

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
	// Expected output 0, 1, 3, 7, 4, 2, 5, 6
	_ = DepthFirstSearch(a)
	_ = DepthFirstSearch_Recursion(a)
}
