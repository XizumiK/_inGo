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
	head.Print()

	head.InsertAtBeginning(1) // You can use <head> to replace <(&head)>
	fmt.Println("InsertAtBeginning(1)")
	head.Print()

	head.InsertAtBeginning(2)
	fmt.Println("InsertAtBeginning(2)")
	head.Print()

	head.Insert(9, 1)
	fmt.Println("Insert(9, 1)")
	head.Print()

	head.InsertAtEnd(8)
	fmt.Println("InsertAtEnd(8)")
	head.Print()

	head.Update(0, 1)
	fmt.Println("Update(1)")
	head.Print()

	head.DeleteAtBeginning()
	fmt.Println("DeleteAtBeginning()")
	head.Print()

	head.Delete(1)
	fmt.Println("Delete(1)")
	head.Print()

	fmt.Println("Search(9)")
	resP, res := head.Search(9)
	fmt.Println(resP, res)

	head.Resverse()
	head.Print()

	head.ReverseRecursive()
	head.Print()

	fmt.Println("<----------Linked-list---------->")

	// Reverse linked-list
	fmt.Println("<------Reverse Linked-list------>")
	fmt.Println("<------Reverse Linked-list------>")
}
