package main

import "fmt"

func main() {

	tree := NewNode("root")
	a := NewNode("a")
	b := NewNode("b")
	c := NewNode("c")
	d := NewNode("d")
	e := NewNode("e")
	f := NewNode("f")

	tree.Add(a)
	tree.Add(b)

	a.Add(f)

	b.Add(c)
	b.Add(d)

	d.Add(e)

	fmt.Println("BFS")
	fmt.Println(tree.BFS("", 0))

	fmt.Println("DFS")
	fmt.Println(tree.DFS("", 0))

}
