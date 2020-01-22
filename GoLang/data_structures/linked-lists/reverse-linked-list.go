/*Given a singly linked list, return another linked list that is the reverse of the first.*/

package main 

import "fmt"

type Node struct {
	value int
	next *Node
}

type LinkedList struct {
	head *Node
}

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

func (l *LinkedList) display() {
	node := l.head
	for node != nil {
		fmt.Printf("%v ->", node.value)
		node = node.next
	}
	fmt.Printf("\n===================================\n")
}


func (l *LinkedList) to_array() []int {

	arr := []int{}
	node := l.head
	for node != nil {
		arr = append(arr, node.value)
		node = node.next
	}	
	return arr
}

// implement a reverse method - O(N)
func (l *LinkedList) reverse() *LinkedList{
	new_list := l
	var prev_node *Node

	// We want to take the node from the original linked list and prepend it to the new list
	for l.head != nil {
		new_node := &Node{value: l.head.value}
		new_node.next = prev_node
		prev_node = new_node
		l.head = l.head.next
	}
	new_list.head = prev_node
	return new_list
}

func main() {

	linked_list := LinkedList{}
	list := []int{1, 2, 3, 4, 5, 6, 7}
	for _, v := range list{
		linked_list.append(v)
	}

	linked_list.display()
	linked_list.reverse().display()
}
