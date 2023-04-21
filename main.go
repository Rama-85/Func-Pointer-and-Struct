package main

import (
	"fmt"
)

func main() {
	var x = 5
	var y = 6
	a := &x
	b := &y
	fmt.Println("before swap:", x, y)
	x = *a
	y = *b
	temp := *a
	*a = *b
	*b = temp
	fmt.Println("After swap,x and y values are:", x, y)

}
