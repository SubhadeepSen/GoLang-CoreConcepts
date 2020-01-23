package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := 115
	y := 1645.546
	fmt.Println("Type:", reflect.TypeOf(x))
	fmt.Printf("Type: %T\n", y)
}
