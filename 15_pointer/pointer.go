package main

import "fmt"

func main() {
	val := 115
	fmt.Println("Value:", val, "Address:", &val)

	var ptr1 *int = &val
	fmt.Println("ptr1:", ptr1)
	fmt.Println("value held by ptr1:", *ptr1)

	ptr2 := &val
	fmt.Println("ptr2:", ptr2)
	fmt.Println("value held by ptr2:", *ptr1)

	speed := new(int)
	fmt.Println("Speed before increas():", *speed)
	increase(speed)
	fmt.Println("Speed after increas():", *speed)
}

func increase(speedPtr *int) {
	*speedPtr = 80
}
