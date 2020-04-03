package main

type Tree struct {
	Root *Node
}

type Node struct {
	Data interface{}
	Left *Node
	Right *Node
}
