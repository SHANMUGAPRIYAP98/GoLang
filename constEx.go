package main

import (
	"fmt"
)

func main() {
	//constants

	const a = 20 // value of variable cannot be changed once assigned
	fmt.Println("Value of a : ", a)

	const (
		val1 = 5
		val2 = false
		val3 = "green"
		val4 = 3.14
	)

	data := val1
	data = 25
	fmt.Println(data)

	fmt.Println("Value of val1 ", val1, "Value of val2 ", val2, "Value of val3 ", val3, "Value of val4 ", val4)

	// const val = math.Cbrt(8) cannot assign the value to const variable during func call

}
