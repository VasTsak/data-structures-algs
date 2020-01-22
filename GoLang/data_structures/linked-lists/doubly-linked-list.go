// This type of list has connections backwards and forwards through the list.
package main

import "fmt"

type Node struct {
	value int
	next *Node
	previous *Node
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
}

// implement an append method - O(1) 
func (dl *DoublyLinkedList) append (i int) error {
	
	n := &Node{value: i}

	if dl.head == nil {
		dl.head = n
		dl.tail = dl.head
		return nil
	} else {
		dl.tail.next = n
		dl.tail.next.previous = dl.tail
		dl.tail = dl.tail.next
		return nil
	}
}

// implement a method that transforms a DoublyLinkedList into an array - O(N) 
func (dl *DoublyLinkedList) to_array() []int {

	arr := []int{}
	node := dl.head
	for node != nil {
		arr = append(arr, node.value)
		node = node.next
	}	
	return arr
}

func (dl *DoublyLinkedList) display() {
	node := dl.head
	for node != nil {
		fmt.Printf("%v ->", node.value)
		node = node.next
	}
}

// implement a method that reverses the DoublyLinkedList - O(N) 
func (dl *DoublyLinkedList) reverse() {

	curr := dl.head 
	var next *Node
	var prev *Node
	dl.tail = dl.head

	for curr != nil {
		next = curr.next
		curr.next = prev
		prev = curr 
		curr = next
	}
	dl.head = prev
}

func main() {

	dl := DoublyLinkedList{}
	dl.append(5)
	dl.append(4)
	dl.append(3)
	dl.append(1)

	fmt.Println("Going forward through the list, should print 5, 4, 3, 1")
	node := dl.head
	for node != nil{
		fmt.Println(node.value)
		node = node.next
	}

	fmt.Println("\nGoing backward through the list, should print 1, 3, 4, 5")
	node = dl.tail
	for node != nil{
		fmt.Println(node.value)
		node = node.previous
	}

	linked_list_array := dl.to_array()
	fmt.Println(linked_list_array)

	dl.display()
}