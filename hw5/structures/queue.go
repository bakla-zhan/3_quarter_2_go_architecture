package structures

type Queue struct {
	List *List
}

func (q *Queue) Len() int {
	return q.List.Len
}

func (q *Queue) Push(elem int) {
	node := &Node{
		Data: elem,
	}
	q.List.Append(node)
}

func (q *Queue) Pop() *Node {
	if q.Len() > 0 {
		head := q.List.Head
		q.List.Delete(head)
		return head
	}
	return nil
}
