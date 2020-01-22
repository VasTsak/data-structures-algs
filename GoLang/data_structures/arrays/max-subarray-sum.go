/*
Problem Statement

You have been given an array containg numbers. 
Find and return the largest sum in a contiguous subarray within the input array.
*/

package main

import "fmt"

func max_sum_subarray(arr[]int) int {

	max_sum := arr[0]
	current_sum := arr[0]

	for i:=1; i <= len(arr)-1; i++ {
		current_sum = max(current_sum + arr[i], arr[i])
		max_sum = max(current_sum, max_sum)
	}
	return max_sum
}

func max(x, y int) int {

	if x >= y {
		return x
	} else {
		return y
	}
}
		
func test_function(test_case [][]int) {
	
	arr := test_case[0]
	solution := test_case[1]
	output := max_sum_subarray(arr)
	
	if output == solution[0]{
		fmt.Println("Pass")
	} else {
		fmt.Println("Fail")
	}
}

func main() {
	
	arr := []int{1, 2, 3, -4, 6}
	solution := []int{8} // sum of array
	test_case := [][]int{arr, solution}

	test_function(test_case)

	arr = []int{1, 2, -5, -4, 1, 6}
	solution = []int{7}  // sum of last two elements

	test_case = [][]int{arr, solution}
	test_function(test_case)

	arr = []int{-12, 15, -13, 14, -1, 2, 1, -5, 4}
	solution = []int{18}  //# sum of subarray = [15, -13, 14, -1, 2, 1]

	test_case = [][]int{arr, solution}
	test_function(test_case)
}
