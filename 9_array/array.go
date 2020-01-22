package main

import "fmt"

func main() {
	array1DWithInLineValues()
	fmt.Println()
	array1DWithMakeKeyword()
	fmt.Println()
	array1DByVarKeyword()
	fmt.Println()
	capAndLenOfArray()
	fmt.Println()
	array2DWithInLineValues()
}

func array1DWithInLineValues() {
	arr := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		fmt.Printf("Element[%d] = %d\n", i, arr[i])
	}
}

func array1DWithMakeKeyword() {
	arr := make([]int, 10)
	for i := 0; i < len(arr); i++ {
		arr[i] = i + 1
	}
	for i := 0; i < len(arr); i++ {
		fmt.Printf("Element[%d] = %d\n", i, arr[i])
	}
}

func array1DByVarKeyword() {
	var arr [10]int
	for i := 0; i < len(arr); i++ {
		arr[i] = i + 1
	}
	for i := 0; i < len(arr); i++ {
		fmt.Printf("Element[%d] = %d\n", i, arr[i])
	}
}

func capAndLenOfArray() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println("Capacity =", cap(arr))
	fmt.Println("Length =", len(arr))
}

func array2DWithInLineValues() {
	arr := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}
}
