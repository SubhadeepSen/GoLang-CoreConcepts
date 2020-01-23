package main

import "fmt"

type vehicle interface {
	accelerate()
}

func vehicleDetails(v vehicle) {
	fmt.Println(v)
}

type indica struct {
	model string
	color string
	speed int
}

func (i indica) accelerate() {
	fmt.Println("Indica accelerating at", i.speed)
}

type suzuki struct {
	model string
	color string
	speed int
}

func (s suzuki) accelerate() {
	fmt.Println("Indica accelerating at", s.speed)
}

func main() {
	indica := indica{"C546Z", "White", 80}
	suzuki := suzuki{"DZire", "Blue", 120}

	vehicleDetails(indica)
	indica.accelerate()

	vehicleDetails(suzuki)
	suzuki.accelerate()
}
