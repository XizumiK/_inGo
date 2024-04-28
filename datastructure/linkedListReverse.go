package datastructure

func (pointerToHead *NodeHead) Resverse() {
	var prev *Node = nil
	curr := pointerToHead.Head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	pointerToHead.Head = prev
}

func (pointerToHead *NodeHead) ReverseRecursive() {
	if pointerToHead.Head == nil || pointerToHead.Head.Next == nil {
		return
	}
	pointerToHead.Head = reverseList(pointerToHead.Head)
}

func reverseList(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
