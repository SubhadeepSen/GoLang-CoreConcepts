package main

import "fmt"

func main() {
	var num int
	fmt.Println("Enter range: ")
	fmt.Scanln(&num)

	printByCounterForLoop(num)
	fmt.Println()
	printByConditionalForLoop(num)
	fmt.Println()
	infiniteLoop()
	fmt.Println()
	forRangeLoop()
}

// Example of counter-controlled loop
func printByCounterForLoop(num int) {
	fmt.Println("Printing numbers with counter-controlled for-loop, similar to natural for loop.")
	for i := 1; i <= num; i++ {
		fmt.Println(i)
	}
}

// Example of condition-controlled loop
func printByConditionalForLoop(num int) {
	fmt.Println("Printing numbers with condition-controlled for-loop, similar to while loop.")
	i := 1
	for i <= num {
		fmt.Println(i)
		i++
	}
}

// Example of infinite loop
func infiniteLoop() {
	fmt.Println("infinite loop.")
	var num int
	for {
		fmt.Println("Enter 15 to exit from infinite loop: ")
		fmt.Scanln(&num)
		if num == 15 {
			break
		}
	}
}

// Example of infinite loop
func forRangeLoop() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for index, value := range nums {
		fmt.Println("Index:", index, "Value:", value)
	}

	fmt.Println()

	name := "Subhadeep Sen"
	for index, value := range name {
		fmt.Println("Index:", index, "Char ASCII Value:", value)
	}

	fmt.Println()

	nameAge := map[int]string{1: "One", 2: "Two", 3: "Three", 4: "Four", 5: "Five"}
	for key, value := range nameAge {
		fmt.Println("Key:", key, "Value:", value)
	}
}
