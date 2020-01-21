package main

import (
	"fmt"
	"strconv"
)

func main() {
	i := 10
	f := 6.44
	str1 := "101"
	str2 := "10.123"

	fmt.Println("Casting int to float", float64(i))
	fmt.Println("Casting float to int", int(f))
	val, _ := strconv.ParseInt(str1, 0, 64)
	fmt.Println("Parsing int-string to int", val)

	fVal, _ := strconv.ParseFloat(str2, 64)
	fmt.Println("Parsing float-string to float64", fVal)
}
