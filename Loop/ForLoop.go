package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}
	// for like while
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}
	//nested loop
	for i = 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Println(i, j)
		}
	}
}
