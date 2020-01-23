package main

import "fmt"

type employee struct {
	id        int
	firstName string
	lastName  string
	addr      address
}

type address struct {
	state string
	city  string
}

func main() {
	address := address{"Telangana", "Hyderabad"}
	emp := employee{id: 123, firstName: "Subhadeep", lastName: "Sen", addr: address}
	fmt.Println(emp)
	fmt.Println(emp.firstName)
	fmt.Println(emp.addr.city)

	emp.id = 645
	fmt.Println(emp)

	emp.details()
	address.details()
}

func (addr address) details() {
	fmt.Println("Address Details:", addr)
}

func (emp employee) details() {
	fmt.Println("Employee Details:", emp)
}
