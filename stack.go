package main

// Stack is a basic LIFO stack that resizes as needed.
type Stack struct {
	Nodes []*Node
	Count int
}
type Node struct {
	Value string
}

func NewStack() *Stack {
	return &Stack{}
}

// Push adds a node to the stack.
func (s *Stack) Push(n *Node) {
	s.Nodes = append(s.Nodes[:s.Count], n)
	s.Count++
}

// Pop removes and returns a node from the stack in last to first order.
func (s *Stack) Pop(count int) (*Node, int) {
	if count != 0 {
		s.Count = count
	}
	if s.Count == 0 {
		return nil, 0
	}
	s.Count--
	return s.Nodes[s.Count], s.Count
}
