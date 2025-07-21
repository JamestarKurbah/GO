package main

import "fmt"

func main() {
	x, y := 5, 6
	fmt.Println(add(x, y))
	fmt.Println(factorial(5))
	// multiple parameter
	fmt.Println(addemup(10, 20, 30, 19, 1))
}
func add(num1, num2 int) int {
	return num1 + num2
}
func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}
func addemup(args ...int) int {
	sum := 0
	for _, value := range args {
		sum += value
	}
	return sum
}
