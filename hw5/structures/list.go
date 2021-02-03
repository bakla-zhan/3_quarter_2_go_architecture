package structures

type Node struct {
	Next *Node
	Data int
}

type List struct {
	Len  int
	Head *Node
	Tail *Node
}

func (l *List) Find(elem int) *Node {
	if l.Head != nil {
		for tmp := l.Head; tmp != nil; tmp = tmp.Next {
			if tmp.Data == elem {
				return tmp
			}
		}
	}
	return nil
}

func (l *List) Add(prev *Node, node *Node) {
	defer func() {
		l.Len++
	}()

	if l.Head == nil {
		l.Head = node
		l.Tail = node
		return
	}
	if l.Len == 1 {
		prev.Next = node
		l.Tail = node
		return
	}
	if l.Tail == prev {
		prev.Next = node
		l.Tail = node
		return
	}
	if l.Head == prev {
		node.Next = l.Head.Next
		l.Head.Next = node
		return
	}
	node.Next = prev.Next
	prev.Next = node
}

func (l *List) Append(node *Node) {
	if l.Len == 1 {
		l.Add(l.Head, node)
		return
	}
	l.Add(l.Tail, node)
}

func (l *List) Delete(node *Node) {
	if l.Head != nil {
		if l.Find(node.Data) != nil {
			if l.Len == 1 {
				l.Head = nil
				l.Tail = nil
				l.Len--
				return
			}
			if l.Head == node {
				l.Head = l.Head.Next
				l.Len--
				return
			}
			for tmp := l.Head; tmp != nil; tmp = tmp.Next {
				if tmp.Next == node {
					if tmp.Next == l.Tail {
						l.Tail = tmp
						l.Tail.Next = nil
						l.Len--
						return
					}
					tmp.Next = node.Next
					node.Next = nil
					l.Len--
					return
				}
			}
		}
	}
}
