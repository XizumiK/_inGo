package datastructure

func (pointerToHead *NodeHead) Resverse() {
	if pointerToHead.Head == nil || pointerToHead.Head.Next == nil {
		return
	}

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

	pointerToHead.Head = reverseNode(pointerToHead.Head)
}

func reverseNode(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := reverseNode(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
