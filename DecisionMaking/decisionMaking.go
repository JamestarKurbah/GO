package main

import "fmt"

func main() {
	age := 18
	if age < 18 {
		fmt.Println("you can't vote")
	} else {
		fmt.Println("you can vote")
	}
	age = 30
	switch age {
	case 16:
		fmt.Println("Prepare for college")
	case 18:
		fmt.Println("Don't indulge in bad habit")
	case 20:
		fmt.Println("Get youself a job")
	default:
		fmt.Println("Are you even alive?")
	}
}
