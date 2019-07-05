package calc

import "errors"

// Add func
func Add(nums ...int) (int, error) {
	if len(nums) < 2 {
		return -1, errors.New("Please provide atleast 2 input arguements")
	}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum, nil
}

// Sub func
func Sub(num1 int, num2 int) int {
	return num1 - num2
}

// Mul func
func Mul(num1 int, num2 int) int {
	return num1 * num2
}

// Div func
func Div(num1 int, num2 int) int {
	if num2 == 0 {
		return -1
	}
	return num1 / num2
}
