/*
Problem Statement
Find and return the nth row of Pascal's triangle in the form an array. n is 0-based.

For exmaple, if n = 4, then output = [1, 4, 6, 4, 1].

To know more about Pascal's triangle: https://www.mathsisfun.com/pascals-triangle.html
*/

package main 

import (
	"fmt"
)


func nth_row_pascal(n int) []int {
	if n == 0{
		return []int{1}
	}
	current_row := []int{1}
	var next_number int
	for i := 1; i <= n; i++{
		previous_row := current_row
		current_row = []int{1}
		for j := 1; j < i; j++ {
			next_number = previous_row[j] + previous_row[j - 1]
			current_row = append(current_row, next_number)
		}
		current_row = append(current_row, 1)
	}
	return current_row
}

func test_function(test_case [][]int) string{
    n := test_case[0]
    solution := test_case[1]
	output := nth_row_pascal(n[0])
	status := "Pass"
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

func main() {

	n := []int{0}
	solution := []int{1}

	test_case := [][]int{n, solution}
	test_function(test_case)

	n = []int{1}
	solution = []int{1, 1}

	test_case = [][]int{n, solution}
	test_function(test_case)

	n = []int{2}
	solution = []int{1, 2, 1}

	test_case = [][]int{n, solution}
	test_function(test_case)

	n = []int{3}
	solution = []int{1, 3, 3, 1}

	test_case = [][]int{n, solution}
	test_function(test_case)


	n = []int{4}
	solution = []int{1, 4, 6, 4, 1}

	test_case = [][]int{n, solution}
	test_function(test_case)
}

