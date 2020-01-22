/*In this script we implement a queue using an array*/

package main

import "fmt"

type Queue struct {
	arr []int
	next_index int
	front_index int
	queue_size int
}

func (q *Queue) initialize(initial_size int) {
	q.arr = make([]int, initial_size)
	q.next_index = 0
	q.front_index = -1
	q.queue_size = 0
}

func (q *Queue) enqueue(item int) {
	// if queue is already full --> increase capacity
	if q.queue_size == len(q.arr){
		q.handle_queue_full_capacity()
	}

	// enqueue new element
	q.arr[q.next_index] = item
	q.queue_size += 1
	q.next_index = (q.next_index + 1) % len(q.arr)
	if q.front_index == -1 {
		q.front_index = 0
	}
}

func (q *Queue) size() int{
	return q.queue_size
}

func (q *Queue ) is_empty() bool {
	return q.size() == 0
}

func (q *Queue) dequeue() (int, bool) {
	if q.is_empty() {
		q.front_index = -1 // resetting pointers
		q.next_index = 0
		return 0, false
	} else {
		value := q.arr[q.front_index]
		q.front_index = (q.front_index + 1) % len(q.arr)
		q.queue_size -= 1
		return value, true
	}
}

/* However, array has a big disadvantage compared to a linked list. You need to define the array size, so we need to handle 
the cases when the capacity is full*/

func (q *Queue) handle_queue_full_capacity() {
	old_arr := q.arr
	q.arr = make([]int, 2 * len(old_arr))

	index := 0
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
	// copy all elements from front of queue (front-index) until end
	for i := q.front_index; i <= len(old_arr); i++ {
		q.arr[index] = old_arr[i]
		index += 1
	}
	// case: when front-index is ahead of next index
	for i := 0; i <= q.front_index; i++ {
		q.arr[i] = old_arr[i]
		index += 1
	}
	// reset pointers
	q.front_index = 0
	q.next_index = index
}

func main() {
	queue_test := []int{1, 2, 3, 4, 5}
	q := &Queue{}
	q.initialize(10)

	for _, v := range queue_test {
		q.enqueue(v)
	}

	// Test size
	if q.size() == len(queue_test) {
		fmt.Println("Pass")
	} else {
		fmt.Println("Fail")
	}

	// Test dequeue
	v, val := q.dequeue()
	if v == 1 {
		fmt.Println("Pass")
	} else {
		fmt.Println("Fail")
	}

	// Test enqueue
	v, val = q.dequeue()
	if v == 2 {
		fmt.Println("Pass")
	} else {
		fmt.Println("Fail")
	}
	v, val = q.dequeue()
	if v == 3 {
		fmt.Println("Pass")
	} else {
		fmt.Println("Fail")
	}
	v, val = q.dequeue()
	if v == 4 {
		fmt.Println("Pass")
	} else {
		fmt.Println("Fail")
	}
	q.enqueue(6)
	v, val = q.dequeue()
	if v == 5 {
		fmt.Println("Pass")
	} else {
		fmt.Println("Fail")
	}
	v, val = q.dequeue()
	if v == 6 {
		fmt.Println("Pass")
	} else {
		fmt.Println("Fail")
	}
	v, val = q.dequeue()
	if val == false {
		fmt.Println("Pass")
	} else {
		fmt.Println("Fail")
	}
}