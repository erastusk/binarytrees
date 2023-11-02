package main

import "testing"

func Benchmark_Iterative(b_ *testing.B) {
	a := CreateNodeTree()
	for i := 0; i < b_.N; i++ {
		DepthFirstSearch_Iterative(a)
	}
}

func Benchmark_Recursion(b_ *testing.B) {
	a := CreateNodeTree()
	for i := 0; i < b_.N; i++ {
		DepthFirstSearch_Recursion(a)
	}
}
