package main

import "fmt"

func main() {

	list := &List{}

	list.Add(5)
	list.Add(8)
	list.Add(3)

	//fmt.Println(list.Get(3))
	fmt.Println(list.ToArray())
	fmt.Println(list.Remove(3))
	fmt.Println(list.ToArray())

}
