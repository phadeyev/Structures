package main

import (
	"errors"
	"fmt"
	"strings"
)

type Node struct {
	Data     string
	children []*Node
}

func NewNode(data string) *Node {
	return &Node{
		Data:     data,
		children: nil,
	}
}

func (n *Node) Add(node *Node) {
	n.children = append(n.children, node)
}

func (n *Node) Get(i int) (*Node, error) {
	if i < 0 || i >= len(n.children) {
		return nil, errors.New("out of range")
	}
	return n.children[i], nil
}

func (n *Node) Len() int {
	return len(n.children)
}

func (n *Node) BFS(needle string, step int) *Node {
	step++
	for _, child := range n.children {
		fmt.Printf("%s--> %s\n", strings.Repeat("--", step), child.Data)
		if child.Data == needle {
			return child
		}
	}

	for _, child := range n.children {
		if res := child.BFS(needle, step); res != nil {
			return res
		}
	}

	return nil
}

func (n *Node) DFS(needle string, step int) *Node {
	step++
	for _, child := range n.children {
		fmt.Printf("%s--> %s\n", strings.Repeat("--", step), child.Data)
		if child.Data == needle {
			return child
		}
		if res := child.DFS(needle, step); res != nil {
			return res
		}
	}

	return nil
}
