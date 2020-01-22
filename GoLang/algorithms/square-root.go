/*
Finding the Square Root of an Integer
Find the square root of the integer without using any GoLang module. You have to find the floor value of the square root.
The expected time complexity is O(log(n))
*/

package main

import(
	"fmt"
)

func sqrt(number int) (int, bool){

	var value int
	var imaginary_root bool

	if number < 0 {
		imaginary_root = true
		value = -1
	} else {
		imaginary_root = false
		value = 1
	}

	if number == 0{
		imaginary_root = true
		value = 0 
		return value, imaginary_root
	}

	for true{
		if imaginary_root == false {
			// number / value is expeted to return an integer
			if number/value == value {
				return value, imaginary_root
			} else if number / value > value {
				value += ((number / value) - value) / 2 // integer is expected back
			} else {
				value -= (value - (number / value)) / 2
			}
		} else {
			if number / value == -1 * value {
				return -1 * value, imaginary_root
			} else if number / value < value{
				value += ((number / value) + value ) / 2
			} else {
				value -= ((number / value) + value ) / 2
			}
		}
	}
	return 0, true
}

func test(input, expected int, imaginary_root bool){
	solution, imaginary_root := sqrt(input)
	if imaginary_root{
		solution = -1 * solution
	} 
	if expected == solution { 
		fmt.Println("Pass") 
	} else { 
		fmt.Println("Fail") 
	}
}
func main() {
	test(9, 3, false)
	test(0, 0, false)
	test(4, 2, false)
	test(27, 5, false)
	test(9, 3, false)
	test(-1, -1, true)
	test(-25, -5, true)
	test(-27, -5, true)
}
