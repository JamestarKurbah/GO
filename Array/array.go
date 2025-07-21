package main

import "fmt"

func main() {
	EvenNum := [5]int{0, 2, 4, 6, 8}
	//EvenNum[0]=0
	for i, value := range EvenNum {
		fmt.Println(value, i)
	}
	//slice array
	numSlice := []int{5, 4, 3, 2, 1}
	sliced := numSlice[3:5]
	fmt.Println(sliced)
	sliced2 := numSlice[0:]
	fmt.Println(sliced2)
	sliced3 := make([]int, 5, 10)
	fmt.Println(sliced3)
	copy(sliced2, numSlice)
	fmt.Println(numSlice)
	slice4 := append(numSlice, 3, 4, 5, 0)
	fmt.Println(slice4)
}
