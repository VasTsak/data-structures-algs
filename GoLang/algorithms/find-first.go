package main 

import "fmt"

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

func find_first(arr[]int, target int) int{
	/*In this function we want to find the index of the first element having the target value*/
	index := binary_search_recursive(target, arr, 0)
	if index < 0 {
		return -1 
	}
	for arr[index] == target {
		if index == 0 {
			return 0
		} 
		if arr[index-1] == target{
			index -= 1
		} else {
			return index
		}
	}
	return index
}

func test_function(arr[]int, target int, expected int) {
	solution := find_first(arr, target)
	if solution == expected {
		fmt.Println("Pass")
	} else {
		fmt.Println("Fail")
	}
}
func main() {
	array := []int{1, 2, 3, 4, 5, 6, 6, 7,  7, 7, 7, 8, 8, 8, 8, 8, 9, 10, 12, 12, 13, 14, 15}
	target := 7
	expected := 7
	test_function(array, target, expected)
	target = 9
	expected = 16
	test_function(array, target, expected)
	target = 11
	expected = -1
	test_function(array, target, expected)

}