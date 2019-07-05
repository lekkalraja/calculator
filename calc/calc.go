package calc

// Add func
func Add(nums ...int) int {
	if len(nums) < 2 {
		return -1
	}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
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
