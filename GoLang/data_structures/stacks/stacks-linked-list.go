/*In this script we will implement a stack as a linked list*/

package main 

import "fmt"

type Node struct {
	value int
	next *Node
}

type Stack struct {
	head *Node
	num_elements int
}

// pushes an item at the top of  the stack - O(1)
func (s *Stack) push (value int) {
	new_node := &Node{value: value}

	if s.head == nil {
		s.head = new_node
	} else {
		new_node.next = s.head // place the new node at the head of the linked list (top)
		s.head = new_node
	}
	s.num_elements += 1
}

// return the item that is at the top of the stack and removes it - O(1)
func (s *Stack) pop () (int, bool) {
	if s.is_empty() {
		return 0, false // the boolean value indicates if the stack has something to pop
	} else {
		value := s.head.value
		s.head = s.head.next
		s.num_elements -= 1
		return value, true
	}
}

// previews the item which is at the top of the stack - return nothing - O(1)
func (s *Stack) top() (int, bool) {
	if s.head == nil {
		return 0, false
	} else {
		s.head.value, true
	}
}

// returns whether the stack is empty or not - O(1)
func (s *Stack) is_empty() bool {
	return s.num_elements == 0
}

// for testing purposes
func (s *Stack) display() {
	node := s.head
	for node != nil {
		fmt.Printf("%v <-", node.value)
		node = node.next
	}
	fmt.Printf("\n===================================\n")
}

// for testing purposes
func (s *Stack) to_array() []int {

	arr := []int{}
	node := s.head
	for node != nil {
		arr = append(arr, node.value)
		node = node.next
	}	
	return arr
}

func main () {
	
	stack := []int{1, 2, 3, 4, 5}
	s := &Stack{}

	// test push 
	for _, v := range stack {
		s.push(v)
	}

	stack_to_array := s.to_array() 
	stack_test := []int{5, 4, 3, 2, 1}// we expect to be reversed
	status := "Pass"
	for i, v := range stack_to_array {
		if v != stack_test[i] {
			status = "Fail"
			break
		}
	}

	fmt.Println(status)

	// Test num_elements
	if s.num_elements == len(stack_test) {
		fmt.Println("Pass")
	} else {
		fmt.Println("Fail")
	}

	// Test pop
	poped_item, is_valid := s.pop()
	if poped_item == 5 && is_valid{
		fmt.Println("Pass")
	} else {
		fmt.Println("Fail")
	}
}