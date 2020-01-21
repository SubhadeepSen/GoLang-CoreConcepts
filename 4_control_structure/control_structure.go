package main

import "fmt"

func main() {
	fmt.Print("if-else example, ")
	ifElse()
	fmt.Print("switch-case example, ")
	switchCase()
	fmt.Print("switch-case-fallthrough example, ")
	switchCaseFallThrough()
	fmt.Print("goto example, ")
	gotoExample()
}

func ifElse() {
	var age int
	fmt.Println("Enter your age:")
	fmt.Scanln(&age)

	if age < 18 {
		fmt.Println("Your are below 18.")
	} else {
		fmt.Println("You are eligible for vote.")
	}
}

func switchCase() {
	var val int
	fmt.Println("Enter an integer:")
	fmt.Scanln(&val)

	switch val {
	case 10:
		fmt.Println("value is 10")
	case 20:
		fmt.Println("value is 20")
	case 30:
		fmt.Println("value is 30")
	default:
		fmt.Println("you have entered", val)
	}
}

func switchCaseFallThrough() {
	var val int
	fmt.Println("Enter an integer:")
	fmt.Scanln(&val)

	switch val {
	case 10:
		fmt.Println("value is 10")
		fallthrough
	case 20:
		fmt.Println("value is 20")
		fallthrough
	case 30:
		fmt.Println("value is 30")
		fallthrough
	default:
		fmt.Println("you have entered", val)
	}
}

func gotoExample() {
	var age int
AGE_LABEL:
	fmt.Println("Enter your age:")
	fmt.Scanln(&age)

	if age < 18 {
		fmt.Println("Your are below 18.")
		goto AGE_LABEL
	} else {
		fmt.Println("You are eligible for vote.")
	}
}
