package structures

type Stack struct {
	List *List
}

func (s *Stack) Len() int {
	return s.List.Len
}

func (s *Stack) Push(elem int) {
	node := &Node{
		Data: elem,
	}
	s.List.Append(node)
}

func (s *Stack) Pop() *Node {
	if s.Len() > 0 {
		tail := s.List.Tail
		s.List.Delete(tail)
		return tail
	}
	return nil
}
