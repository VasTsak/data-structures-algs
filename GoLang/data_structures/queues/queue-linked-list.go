/*In this script we implement a queue data structure using a linked list*/
package main 

import "fmt"

type Node struct {
	value int
	next *Node
}

type Queue struct {
	head *Node
	tail *Node
	num_elements int
}

func (q *Queue) size() int {
	return q.num_elements
}

func (q *Queue) is_empty() bool {
	return q.num_elements == 0
}

// Create the enqueue method - O(1)
func (q *Queue) enqueue(item int) error {
	new_node := &Node{value: item}
	if q.head == nil {
		q.head = new_node
		q.tail = q.head
	} else {
		q.tail.next = new_node //add data to the next attribute of the tail (i.e. the end of the queue)
		q.tail = q.tail.next //shift the tail (i.e., the back of the queue)
	}
	q.num_elements += 1
	return nil
}

// Create the dequeue method - O(1)
func (q *Queue) dequeue() (int, bool ){
	if q.is_empty() == true {
		return 0, false
	} else {
		value := q.head.value
		q.head = q.head.next
		q.num_elements -= 1
		return value, true
	}
}

// for testing purposes
func (q *Queue) display() {
	node := q.head
	for node != nil {
		fmt.Printf("%v ->", node.value)
		node = node.next
	}
	fmt.Printf("\n===================================\n")
}

// for testing purposes
func (q *Queue) to_array() []int {

	arr := []int{}
	node := q.head
	for node != nil {
		arr = append(arr, node.value)
		node = node.next
	}	
	return arr
}


func main() {
	queue_test := []int{1, 2, 3, 4, 5}
	q := &Queue{}

	// test enqueue 
		
	for _, v := range queue_test {
		q.enqueue(v)
	}
	q.display()
	
	queue_to_array :=  q.to_array()
	status := "Pass"
	for i, v := range queue_to_array {
		if v != queue_test[i] {
			status = "Fail"
			break
		}
	}
	fmt.Println(status)


	// Test size
	if q.size() == len(queue_test) {
		fmt.Println("Pass")
	} else {
		fmt.Println("Fail")
	}

	// Test dequeue
	dequeue_value, _ := q.dequeue()
	if dequeue_value== 1 {
		fmt.Println("Pass")
	} else {
		fmt.Println("Fail")
	}
}

 