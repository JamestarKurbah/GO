package main

import "fmt"

type Rectangle struct {
	height float64
	width  float64
}

func main() {
	rect1 := Rectangle{10, 5}
	fmt.Println(rect1.height)
	fmt.Println(rect1.width)
	fmt.Println(rect1.area())
}
func (rect *Rectangle) area() float64 {
	return rect.height * rect.width
}
