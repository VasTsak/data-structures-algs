/*A palindrome is a word that is the reverse of itself—that is, it is the same word when read forwards and backwards.
For example:

"madam" is a palindrome
"abba" is a palindrome
"cat" is not
"a" is a trivial case of a palindrome */

package main 

import "fmt"

func is_palindrome(input string) bool {
	/* 
    Return True if input is palindrome, False otherwise.
    
    Args:
       input: input to be checked if it is palindrome
	*/
	if len(input) <= 1 {
		return true
	} else {
		first_char := input[0]
		last_char := input[(len(input)-1)]

		// sub_input is input with first and last char removed
		sub_input := input[1:(len(input)-1)]
		return (first_char == last_char) && is_palindrome(sub_input)
	}
}

func test_function(test_case []string) {
	str := test_case[0]
	solution_str := test_case[1]
	output := is_palindrome(str)
	solution := true
	if solution_str == "false" {
		solution = false
	} 
	if output == solution {
		fmt.Println("Pass")
	} else {
		fmt.Println("Fail")
	}
}

func main() {
	test_case := []string{"", "true"}
	test_function(test_case)
	test_case = []string{"a", "true"}
	test_function(test_case)
	test_case = []string{"madam", "true"}
	test_function(test_case)
	test_case = []string{"abba", "true"}
	test_function(test_case)
	test_case = []string{"programming", "false"}
	test_function(test_case)
	test_case = []string{"data structures", "false"}
	test_function(test_case)
}