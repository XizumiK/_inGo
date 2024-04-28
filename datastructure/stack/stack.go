package stack

import "github.com/XizumiK/_inGo/datastructure/linkedList"

type Stack struct {
	Array linkedList.NodeHead
}

func (s *Stack) Push(i int) {
	s.Array.InsertAtBeginning(i)
}

func (s *Stack) Pop() int {
	peek := s.Array.Head.Data
	s.Array.DeleteAtBeginning()
	return peek
}

func (s *Stack) Peek() int {
	return s.Array.Head.Data
}

func (s *Stack) IsEmpty() bool {
	return s.Array.Nums == 0
}
