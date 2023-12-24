package main

import "fmt"

type StackNode struct {
	data rune
	prev *StackNode
}

type Stack struct {
	head *StackNode
	size int
}

func (s *Stack) push(value rune) {
	s.size++
	new_node := &StackNode{data: value}

	if s.size == 0 {
		s.head = new_node
		return
	}
	new_node.prev = s.head
	s.head = new_node
}

func (s *Stack) pop() rune {
	if s.size == 0 {
		return 'x'
	}
	if s.size == 1 {
		s.size = 0
		value := s.head.data
		s.head = nil
		return value
	}

	value := s.head.data
	s.head = s.head.prev
	s.size--
	return value
}

func (s *Stack) print() {
	curr := s.head
	fmt.Println("Stack Buffer")
	for curr.prev != nil {
		fmt.Println("**********")
		fmt.Println(curr.data)
		curr = curr.prev
	}
	fmt.Println(curr.data)
	fmt.Println("********************")

}
func isValid(s string) bool {
	stk := new(Stack)
	for _, char := range s {
		stk_val := stk.pop()
		switch char {
		case '(':
			if stk_val != ')' {
				return false
			}
			break
		case '[':
			if stk_val != ']' {
				return false
			}
			break
		case '{':
			if stk_val != '}' {
				return false
			}
			break
		}
		fmt.Println(char)
		stk.push(char)
		//fmt.Println(stk.size)
		stk.print()

	}
	//fmt.Println(stk.size)
	if stk.size > 0 {
		return false
	}
	return true
}

func main() {
	fmt.Println(isValid("()"))
}
