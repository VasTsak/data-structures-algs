/*How the algorithm works:
    1. Find the center of the list 
    2. Check to see if the element at the center is your target.
    3. If it is, return the index.
    4. If not, check if the target is greater or less than that element.
    5. If greater, move the lower bound to just above the current center
    6. If less, move the upper bound to just below the current center
    7. Repeat steps 1-6 until you find the target or until the bounds are the same or cross (the upper bound is less than the lower bound).
*/

package main 

import "fmt"

func binary_search_iterative(arr []int, target int) int {
	/*Write a function that implements the binary search algorithm using iteration
   
    args:
      array: a sorted array
      target: the element you're searching for
   
    returns:
      int: the index of the target, if found, otherwise -1
	*/
	start_index := 0
	end_index := len(arr) - 1
	var center_index int
	for start_index <= end_index {
		center_index = (start_index + end_index) / 2
		if target == arr[center_index] {
			return center_index
		} else if target > arr[center_index] {
			start_index = center_index + 1
		} else {
			end_index = center_index - 1
		}
	}  
	return -1
}

func test_function(arr[]int, target int, expected int){
	solution := binary_search_iterative(arr, target)
	if solution == expected{
		fmt.Println("Pass")
	} else {
		fmt.Println("Fail")
	}
}

func main(){
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 6
	expected := 6
	test_function(arr, target , expected)

	arr = []int{0, 1, 2, 3, 4}
	target = 6
	expected = -1
	test_function(arr, target , expected)

	arr = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	target = 11
	expected = 11
	test_function(arr, target , expected)

	arr = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	target = 10
	expected = 10
	test_function(arr, target , expected)
}