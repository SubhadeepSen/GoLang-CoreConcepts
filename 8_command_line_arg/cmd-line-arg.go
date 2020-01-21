package main

import (
	"fmt"
	"os"
)

func main() {
	// First arg is the current path of the .exe
	// Second to last are the values passed
	for i := 0; i < len(os.Args); i++ {
		fmt.Println(os.Args[i])
	}

	fmt.Println()

	args := os.Args[1:]
	for i := 0; i < len(args); i++ {
		fmt.Println(args[i])
	}
}
