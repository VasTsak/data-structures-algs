package main 

import "fmt"

func factorial(n int) int {
	var fact_result int
	if n == 0 {
		fact_result = 1
	} else {
		fact_result = n * factorial(n - 1)
	}
	return fact_result
}

func test_function(test_case []int) {
	n := test_case[0]
	solution := test_case[1]
	output := staircase(n)
	if output == solution {
		fmt.Println("Pass")
	} else {
		fmt.Println("Fail")
	}
}

func main() {
	test_case := []int{0, 1}
	test_function(test_case)
	test_case = []int{1, 1}
	test_function(test_case)
	test_case = []int{5, 120}
	test_function(test_case)
}
