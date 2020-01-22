/*
Problem statement

You are given a non-negative number in the form of list elements. 
For example, the number 123 would be provided as arr = [1, 2, 3]. 
Add one to the number and return the output in the form of a new list. 
*/

package main 

import (
	"fmt"
)

func add_one(arr []int) []int{
	output := 1
	var borrow int
	for i := len(arr)-1; i >= 0; i-- {

		output += arr[i]
		borrow = output / 10
		if borrow == 0 {
			arr[i] = output
			break
		} else {
			arr[i] = output % 10
			output = borrow
		}
	}
	arr = append([]int{borrow}, arr...)

	index := 0
	for arr[index] == 0 {
		index += 1
	}   
	return arr[index:]
}
	
func test_function(test_case [][]int) string {
	arr := test_case[0]
	solution := test_case[1]
	status := "Pass"
	output := add_one(arr)
	for i, v := range output {
		if v != solution[i]{
			status = "Fail"
			fmt.Println(status)
			return status
		}
	}
	fmt.Println(status)
	return status
}

func main(){

	arr := []int{0}
	solution := []int{1}
	test_case := [][]int{arr, solution}
	test_function(test_case)

	arr = []int{1, 2, 3}
	solution = []int{1, 2, 4}
	test_case = [][]int{arr, solution}
	test_function(test_case)

	arr = []int{9, 9, 9}
	solution = []int{1, 0, 0, 0}
	test_case = [][]int{arr, solution}
	test_function(test_case)
}