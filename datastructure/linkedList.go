package datastructure

import (
	"errors"
	"fmt"
	"log"
)

type Node struct {
	Data int
	Next *Node
}

type NodeHead struct {
	Next *Node
	Nums int
}

func (pointerToHead *NodeHead) InsertAtBeginning(data int) {
	newNode := &Node{
		Data: data,
	}
	(*newNode).Next = (*pointerToHead).Next
	(*pointerToHead).Next = newNode
	pointerToHead.Nums++
}

func (pointerToHead *NodeHead) InsertAtEnd(data int) {
	newNode := &Node{
		Data: data,
	}
	endNode := pointerToHead.Next
	for i := 1; i < pointerToHead.Nums; i++ {
		endNode = endNode.Next
	}
	newNode.Next = endNode.Next
	endNode.Next = newNode
	pointerToHead.Nums++
}

func (pointerToHead *NodeHead) Insert(data int, position int) {
	if position > pointerToHead.Nums {
		err := errors.New("position is not exist in the linked list")
		log.Println(err.Error())
		return
	} else if position == 0 {
		pointerToHead.InsertAtBeginning(data)
		return
	}

	newNode := &Node{
		Data: data,
	}
	theNode := pointerToHead.Next
	for i := 1; i < position; i++ {
		theNode = theNode.Next
	}
	newNode.Next = theNode.Next
	theNode.Next = newNode
	pointerToHead.Nums++
}

func (pointerToHead *NodeHead) Print() {
	temp := pointerToHead.Next
	for temp != nil {
		fmt.Print(temp.Data)
		temp = temp.Next
	}
	fmt.Print(" ")
	fmt.Print(pointerToHead.Nums)
}
