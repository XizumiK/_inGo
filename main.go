package main

import (
	"fmt"

	"github.com/XizumiK/_inGo/datastructure"
)

func main() {
	// Basic linked-list
	fmt.Println("<----------Linked-list---------->")
	var head datastructure.NodeHead // Create head which contains a pointer that points to a Node
	fmt.Println("Linked-list init")

	(&head).InsertAtBeginning(0)
	fmt.Println("InsertAtBeginning(0)")

	head.InsertAtBeginning(1) // You can use <head> to replace <(&head)>
	fmt.Println("InsertAtBeginning(1)")

	head.InsertAtBeginning(2)
	fmt.Println("InsertAtBeginning(2)")

	head.Insert(9, 1)
	fmt.Println("Insert(9, 1)")

	head.InsertAtEnd(8)
	fmt.Println("InsertAtEnd(8)")

	head.Print()
	fmt.Println("")

	head.Delete(1)
	fmt.Println("Delete(1)")

	head.Print()
	fmt.Println("")

	fmt.Println("<----------Linked-list---------->")

	// Reverse linked-list
	fmt.Println("<------Reverse Linked-list------>")
	fmt.Println("<------Reverse Linked-list------>")
}
