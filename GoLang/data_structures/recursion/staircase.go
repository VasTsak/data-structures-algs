/*Problem Statement
Suppose there is a staircase that you can climb in either 1 step, 2 steps, or 3 steps. In how many possible ways can you climb the 
staircase if the staircase has n steps? Write a recursive function to solve the problem.

Example:

n = 3
output = 4
The output is 4 because there are four ways we can climb the staircase:

1. 1 step +  1 step + 1 step
2. 1 step + 2 steps 
3. 2 steps + 1 step
4. 3 steps
*/

package main 

import "fmt"

func staircase(n int) int {
	if n <= 0 {
		return 1
	}
	if n == 1{
		return 1
	} else if n == 2 {
		return 2
	} else if n == 3 {
		return 4
	}

	return staircase(n - 1) + staircase(n - 2) + staircase(n - 3)
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

	test_case := []int{3, 4}
	test_function(test_case)
	test_case = []int{4, 7}
	test_function(test_case)
	test_case = []int{7, 44}
	test_function(test_case)
}