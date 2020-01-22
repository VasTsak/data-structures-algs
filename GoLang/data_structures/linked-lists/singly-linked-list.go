/*Singly Linked Lists

In this linked list, each node in the list is connected only to the next node in 
the list.*/
package  main

import "fmt"

type Node struct {
	value int
	next *Node
}

type LinkedList struct {
	head *Node
}

// implement an append method - O(N) only with head
func (l *LinkedList) append(value int) error {
	n := &Node{value: value}
	if l.head == nil {
		l.head = n
		return nil
	} else {
		node := l.head
		for node.next != nil {
			node = node.next
		}
		node.next = n
		return nil
	}
}

func (l *LinkedList) prepend(value int) error {
	//  Prepend a node to the beginning of the list
	n := &Node{value: value}
	if l.head == nil {
		l.head = n
		return nil
	}
	new_head := n
	new_head.next = l.head
	l.head = new_head
	return nil
}

// implement a remove method - O(N) 

func (l *LinkedList) remove(value int) error{
	// Delete the first node with the desired data
	if l.head == nil {
		return nil
	}

	if l.head.value == value {
		l.head = l.head.next
		return nil
	}

	node := l.head
	for node.next != nil {
		if node.next.value == value {
			node.next = node.next.next
			return nil
		}
		node = node.next
	}
	return nil
}

// implement a pop method - O(1) 
func (l *LinkedList) pop() (int, bool) {
	// Return the first node's value and remove it from the list.
	if l.head == nil {
		return 0, false
	}
	node := l.head
	l.head = l.head.next
	return node.value, true
}

func (l *LinkedList) insert (value, pos int) error {
	//Insert value at pos position in the list. If pos is larger than the length of the list, append to the end of the list.
	if pos == 0 {
		l.prepend(value)
		return nil
	}

	index := 0
	node := l.head
	var new_node *Node
	n := &Node{value: value}

	for node.next != nil && index <= pos {
		if pos - 1 == index {
			new_node = n
			new_node.next = node.next
			node.next = new_node
			return nil
		}
		index += 1
		node = node.next
	}
	l.append(value)
	return nil
} 

// implement a search method - O(N) 
func (l *LinkedList) search(value int) *Node {
	// Search the linked list for a node with the requested value and return the node.
	if l.head == nil {
		return nil
	}

	node := l.head 
	for node != nil {
		if node.value == value {
			return node
		}
		node = node.next
	}
	return nil
}

// Add a method to_array() to LinkedList that converts a linked list back into an array. - O(N)

func (l *LinkedList) to_array() []int {

	arr := []int{}
	node := l.head
	for node != nil {
		arr = append(arr, node.value)
		node = node.next
	}	
	return arr
}

func (l *LinkedList) display() {
	node := l.head
	for node != nil {
		fmt.Printf("%v ->", node.value)
		node = node.next
	}
	fmt.Printf("\n===================================\n")
}

func main() {

	// Test node
	head := &Node{value: 2}
	head.next = &Node{value: 1}
	
	fmt.Println(head.value)
	fmt.Println(head.next.value)

	// test append
	linked_list := LinkedList{}
	list := []int{1, 2, 1, 4, 5, 6, 7}
	for _, v := range list{
		linked_list.append(v)
	}

	linked_list.display()

	// test append
	linked_list.prepend(-1)

	linked_list.display()

	// test remove
	linked_list.remove(-1)

	linked_list.display()

	// test pop
	value, not_empty := linked_list.pop()
	fmt.Println(value, not_empty)
	linked_list.display()

	// test insert
	linked_list.insert(150, 5)

	linked_list.display()

	// test search 
	search_node := linked_list.search(5)
	fmt.Println(search_node.value)

	linked_list_array := linked_list.to_array()
	fmt.Println(linked_list_array)
}

