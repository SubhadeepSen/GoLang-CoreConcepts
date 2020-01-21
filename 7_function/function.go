package main

import "fmt"

func main() {

	sayHello()
	sum := sum(15, 5)
	fmt.Println("sum is:", sum)

	add, sub := sumAndSub(15, 5)
	fmt.Println("sum is:", add, "sub is:", sub)

	lambdaFunction()

	fmt.Println("Factorial of 5 is:", factorial(5))
}

//example of normal function
func sayHello() {
	fmt.Println("Hello !")
}

//example of function with parameters and single return type
func sum(a int, b int) int {
	return a + b
}

//example of function with parameters and dual return type
func sumAndSub(a int, b int) (int, int) {
	return (a + b), (a - b)
}

//example of anonymous function or lambda function or closure
func lambdaFunction() {
	num := 10
	squareNum := func() int {
		num *= num
		return num
	}

	fmt.Println(squareNum())
	fmt.Println(squareNum())

	square := func(n int) int {
		return n * n
	}
	fmt.Println(square(5))
	fmt.Println(square(5))
}

//example of recursion
func factorial(n int) int {
	if n == 1 || n == 0 {
		return 1
	}

	return n * factorial(n-1)
}
