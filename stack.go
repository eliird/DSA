// Implement stack using go
package main

type NodeStack struct {
	data int
	next *NodeStack
}

type stack struct {
	head *NodeStack
	tail *NodeStack
	size int
}

func (s *stack) peek() int {
	if s.size == 0 {
		return -1
	}
	return s.head.data
}

func (s *stack) insert(value int) {

	new_node := &NodeStack{data: value}
	s.size++
	if s.size == 1 {
		s.head = new_node
		s.tail = new_node
		return
	}

	new_node.next = s.head
	s.head = new_node
}

func (s *stack) pop() int {
	if s.size == 0 {
		return -1
	}
	value := s.head.data
	s.head = s.head.next
	return value
}
