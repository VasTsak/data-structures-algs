package main 

import "fmt"

// Native implementation of binary search in the `contains` function.
func contains(target string, array []string) bool {
	if len(array) == 0 {
		return false
	}
	center := (len(array) - 1) / 2
	if array[center] == target {
		return true
	} else if array[center] < target {
		return contains(target, array[center + 1:])
	} else {
		return contains(target, array[:center])
	}
}

func test_function(array[]string, target string, expected bool) {
	solution := contains(target, array)
	if solution == expected {
		fmt.Println("Pass")
	} else {
		fmt.Println("Fail")
	}
}

func main() {
	array := []string{"a", "b", "c", "d", "e", "f"}
	target := "c"
	test_function(array, target, true)
	target = "z"
	test_function(array, target, false)
	target = "b"
	test_function(array, target, true)
}