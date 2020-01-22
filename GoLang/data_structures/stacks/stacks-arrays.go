package main 

import "fmt"

type Stack struct {
	arr []int
	next_index int
	num_elements int
}

func (s *Stack) initialize(initial_size int) {
	s.arr = make([]int, initial_size)
	s.next_index = 0
	s.num_elements = 0
}


func (s *Stack) push(item int) {
	if s.next_index == len(s.arr) {
		s.handle_stack_overflow()
	}

	s.arr[s.next_index] = item
	s.next_index += 1
	s.num_elements += 1
}

func (s *Stack) pop() (int, bool){
	if s.is_empty() {
		s.next_index = 0
		return 0, false
	}
	s.next_index -= 1
	s.num_elements -= 1
	return s.arr[s.next_index], true
}

func (s *Stack) is_empty() bool {
	return s.num_elements == 0
}

func (s *Stack) size() int {
	return s.num_elements
}

func (s *Stack) handle_stack_overflow() {
	old_arr := s.arr
	s.arr =  make([]int, 2 * len(old_arr))

	for i, v:= range old_arr{
		s.arr[i] = v
	}
}


func main () {
	
	stack := []int{1, 2, 3, 4, 5}
	s := &Stack{}
	s.initialize(10)

	// test push 
	for _, v := range stack {
		s.push(v)
	}

	status := "Pass"
	for i := 0; i < len(stack); i++ {
		if s.arr[i] != stack[i]{
			status = "Fail"
		}
	}
	fmt.Println(status)

	// Test num_elements
	if s.num_elements == len(stack) {
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