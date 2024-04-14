package main

import (
	"algorithm-datastructures/containers/binary_tree"
	"fmt"
)

func main() {
	//NOTHING TODO HERE
	tree := binary_tree.New[string]()
	tree.Add("3")
	tree.Add("2")
	tree.Add("1")
	tree.Add("4")
	fmt.Printf("%v ", tree.Depth())
}
