package main

import "fmt"

func main() {
	arr := []string{"one", "two", "three", "four", "five"}

	//slice of an array in range
	slice1 := arr[1:4]
	printSlice(slice1)
	fmt.Println()

	//omiting the lower bound
	slice2 := arr[:3]
	printSlice(slice2)
	fmt.Println()

	//omiting the upper bound
	slice3 := arr[1:]
	printSlice(slice3)
	fmt.Println()

	//creating a slice and then changing the content of slice, original array content should also be changed
	slice := arr[1:3]
	printSlice(slice)
	fmt.Println()
	slice[0] = "ABC"
	printSlice(slice)
	fmt.Println()
	printSlice(arr)
	fmt.Println()

	lenAndCapacity(arr)
	fmt.Println()

	sliceWithMakeKeyword()
	fmt.Println()

	sliceWithStructKeyword()
}

func printSlice(arr []string) {
	for i := 0; i < len(arr); i++ {
		fmt.Printf("Element[%d] = %s\n", i, arr[i])
	}
}

func lenAndCapacity(arr []string) {
	printSliceLenAndCap(arr)

	slice1 := arr[:0]
	printSliceLenAndCap(slice1)

	slice2 := arr[:4]
	printSliceLenAndCap(slice2)

	slice3 := arr[2:]
	printSliceLenAndCap(slice3)
}

func sliceWithMakeKeyword() {
	slice := make([]string, 10)
	printSliceLenAndCap(slice)
	slice1 := make([]string, 0, 10)
	printSliceLenAndCap(slice1)
	slice2 := slice1[:5]
	printSliceLenAndCap(slice2)
	slice3 := slice2[2:5]
	printSliceLenAndCap(slice3)
}

func printSliceLenAndCap(slice []string) {
	fmt.Printf("length =%d capacity=%d %v\n", len(slice), cap(slice), slice)
}

func sliceWithStructKeyword() {
	slice := []struct {
		i int
		b bool
	}{{1, false}, {2, true}, {3, false}, {4, true}, {5, false}}
	fmt.Println(slice)
}
