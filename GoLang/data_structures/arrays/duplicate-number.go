/*
Problem Statement

You have been given an array of length = n. The array contains integers from 0 to n - 2. 
Each number in the array is present exactly once except for one number which is present twice. 
Find and return this duplicate number present in the array
*/

package main 

import (
	"fmt"
)

func duplicate_number(arr []int) int{
	current_sum := 0
	expected_sum := 0
	
	for _, num := range arr{
		current_sum += num
	}
	
	for i := 0; i <= len(arr)-2; i++ {
		expected_sum += i
	}
	return current_sum - expected_sum
}

func test_function(test_case [][]int) {
	arr := test_case[0]
	solution := test_case[1]
	output := duplicate_number(arr)
	if output == solution[0]{
		fmt.Println("Pass")
	} else {
		fmt.Println("Fail")
	}
}

func main() {

	arr := []int{0, 0}
	solution := []int{0}

	test_case := [][]int{arr, solution}
	test_function(test_case)

	arr = []int{0, 2, 3, 1, 4, 5, 3}
	solution = []int{3}

	test_case = [][]int{arr, solution}
	test_function(test_case)

	arr = []int{0, 1, 5, 4, 3, 2, 0}
	solution = []int{0}

	test_case = [][]int{arr, solution}
	test_function(test_case)		

	arr = []int{0, 1, 5, 5, 3, 2, 4}
	solution = []int{5}

	test_case = [][]int{arr, solution}
	test_function(test_case)
		
}