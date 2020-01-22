/*def merge(left, right):
    merged = []
    left_index = 0
    right_index = 0
    
    # Move through the lists until we have exhausted one
    while left_index < len(left) and right_index < len(right):
        # If left's item is larger, append right's item
        # and increment the index
        if left[left_index] > right[right_index]:
            merged.append(right[right_index])
            right_index += 1
        # Otherwise, append left's item and increment
        else:
            merged.append(left[left_index])
            left_index += 1
     
    # Append any leftovers. Because we've broken from our while loop,
    # we know at least one is empty, and the remaining:
    # a) are already sorted
    # b) all sort past our last element in merged
    merged += left[left_index:]
    merged += right[right_index:]
        
    # return the ordered, merged list
    return merged

# Test this out
merged = merge([1,3,7], [2,5,6])
print(merged)*/

package main 

import "fmt"

func merge(left, right []int) []int{
	merged := []int{}
	left_index := 0
	right_index := 0
	// Move through the lists until we have exhausted one
	for left_index < len(left) && right_index < len(right) {
		// if left's item is larger , append right's item
		// and increment the index
		if left[left_index] > right[right_index] {
			merged = append(merged, right[right_index])
			right_index += 1
		} else {
			merged = append(merged, left[left_index])
			left_index += 1
		}
	}
	// Append any leftovers. Because we've broken from our while loop,
	// we know at least one is empty, and the remaining: 
	// a) are already sorted
	// b) all sort past our last element in merged
	for _, v := range left[left_index:]{
		merged = append(merged, v)
	}
	for _, v :=  range right[right_index:]{
		merged = append(merged, v)
	}
	return merged
}

func main() {
	merged := merge([]int{1,3,7}, []int{2,5,6})
	fmt.Println(merged)
}