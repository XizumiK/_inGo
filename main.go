package main

import "github.com/XizumiK/_inGo/datastructure"

func main() {
	var head datastructure.NodeHead // We create head which contains a pointer that points to a Node
	(&head).InsertAtBeginning(0)
	head.InsertAtBeginning(1) // You can use <head> to replace <(&head)>
	head.InsertAtBeginning(2)
	head.Insert(9, 1)
	head.InsertAtEnd(8)
	head.Print()
}
