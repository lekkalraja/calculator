package main

// Add func
func Add(num1 int, num2 int) int {
	return num1 + num2
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
