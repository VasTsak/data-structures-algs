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


/*
// first solution 
func binary_search_recursive(arr[]int, target int, start_index int, end_index int) int{
	if start_index > end_index {
		return -1
	}
	mid_index := (start_index + end_index) / 2
	if target == arr[mid_index] {
		return mid_index
	} else if target > arr[mid_index] {
		return binary_search_recursive(arr, target,  mid_index + 1, end_index)
	} else {
		return binary_search_recursive(arr, target, start_index, mid_index - 1)
	}
}
*/

// second and cleaner solution
func binary_search_recursive(target int, arr[]int, left int) int{
	if len(arr) == 0{
        return -1
	}
    center := (len(arr)-1) / 2
    if arr[center] == target{
        return center + left
	}else if arr[center] < target{
        return binary_search_recursive(target, arr[center+1:], left+center+1)
	}else{
        return binary_search_recursive(target, arr[:center], left)
	}
}

func test_function(arr[]int, target int, expected int){
	solution := binary_search_recursive(target, arr, 0)
	//solution := binary_search_recursive(arr, target, 0, len(arr)-1)
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

