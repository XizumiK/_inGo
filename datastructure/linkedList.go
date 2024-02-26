package datastructure

import (
	"errors"
	"fmt"
	"log"
)

// Node
// the basic type of linked list
type Node struct {
	Data int
	Next *Node // point to the next node
}

// NodeHead
// the head of linked list, it doesn't contain data
type NodeHead struct {
	Next *Node // point to the first node
	Nums int   // the length of this linked list
}

// InsertAtBeginning
// insert a node with provided data in the first place
func (pointerToHead *NodeHead) InsertAtBeginning(data int) {
	newNode := &Node{ // Creat a new node with data
		Data: data,
	}
	// insert newNode by exchanging the pointer
	(*newNode).Next = (*pointerToHead).Next
	(*pointerToHead).Next = newNode
	pointerToHead.Nums++ // length + 1
}

// InsertAtEnd
// insert a node with provided data in the end of the linked list
func (pointerToHead *NodeHead) InsertAtEnd(data int) {
	newNode := &Node{ // Creat a new node with data
		Data: data,
	}
	// find the end of the linked list
	endNode := pointerToHead.Next
	for i := 1; i < pointerToHead.Nums; i++ {
		endNode = endNode.Next
	}
	// insert newNode by exchanging the pointer
	newNode.Next = endNode.Next
	endNode.Next = newNode
	pointerToHead.Nums++ // length + 1
}

// Insert
// insert data behind position, 0 - First; 1 - 2nd etc
func (pointerToHead *NodeHead) Insert(data int, position int) {
	if position > pointerToHead.Nums || position < 0 { // if the position is out of range
		err := errors.New("position is not exist in the linked list")
		log.Println(err.Error())
		return
	} else if position == 0 { // InsertAtBeginning
		pointerToHead.InsertAtBeginning(data)
		return
	}

	newNode := &Node{ // Creat a new node with data
		Data: data,
	}
	// find the insert position of the linked list
	theNode := pointerToHead.Next
	for i := 1; i < position; i++ {
		theNode = theNode.Next
	}
	// insert newNode by exchanging the pointer
	newNode.Next = theNode.Next
	theNode.Next = newNode
	pointerToHead.Nums++ // length + 1
}

func (pointerToHead *NodeHead) DeleteAtBeginning() {
	firstNode := pointerToHead.Next
	pointerToHead.Next = firstNode.Next
	pointerToHead.Nums--
}

func (pointerToHead *NodeHead) Delete(position int) {
	if position > pointerToHead.Nums || position < 0 { // if the position is out of range
		err := errors.New("position is not exist in the linked list")
		log.Println(err.Error())
		return
	} else if position == 0 {
		pointerToHead.DeleteAtBeginning()
		return
	}
	// find the delete position of the linked list
	theNode := pointerToHead.Next
	for i := 1; i < position; i++ {
		theNode = theNode.Next
	}
	theNextNode := theNode.Next
	// delete the node by change the pointer
	theNode.Next = theNextNode.Next
	// don't worry about the memory used by the node, golang will release it automatically
	pointerToHead.Nums-- // length + 1
}

// Print
// output all data and the length in the linked list
func (pointerToHead *NodeHead) Print() {
	temp := pointerToHead.Next
	for temp != nil {
		fmt.Print(temp.Data)
		temp = temp.Next
	}
	fmt.Print(" ")
	fmt.Print(pointerToHead.Nums)
}
