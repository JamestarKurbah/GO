package main

import "fmt"

func main() {
	var a int = 5
	var b float32 = 4.32
	const pi float64 = 3.14151319475
	x, y := 5, 6
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(pi)
	fmt.Println(x, " , ", y)
	isbool := true
	hate := false
	fmt.Println(isbool && hate)
	fmt.Println(isbool || hate)
	fmt.Println(!isbool)
}
