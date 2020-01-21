package main

import "fmt"

func main() {
	var name string = "Subhadeep Sen"
	var age int = 26
	var salary float64 = 25316.45
	var isEmployeed bool = true

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Salary:", salary)
	fmt.Println("Is Employeed?", isEmployeed)

	fmt.Println("\n* Variable declared with Type-Inferenced, type of the variable will automatically be defined by the compiler from the value assigned to it.")
	fmt.Println("* Once the value is assigned, type cannot be changed, hence statically typed language.")
	fmt.Println()
	empName := "Subhadeep Sen"
	empAge := 26
	empSalary := 25316.45
	empIsEmployeed := true

	fmt.Println("Employee Name:", empName)
	fmt.Println("Employee Age:", empAge)
	fmt.Println("Employee Salary:", empSalary)
	fmt.Println("Employee Is Employeed?", empIsEmployeed)

	const PI float64 = 3.14
	fmt.Println("Constant value of PI", PI)
}
