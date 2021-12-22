package main

import "fmt"

func main() {

	var num1 int                      // variable declaration
	var name string = "Sundar Pichai" // initialization
	var isVaccinated bool = true      // bool initialization
	var weight float32 = 55.6         // float
	var a, b int = 10, 20             // multiple assignment
	var val = "hello world"           // inferred type
	num1 = 9
	fmt.Println("Number1 ", num1)
	num1 = 11.0
	fmt.Println("Number2 ", num1)
	fmt.Println("Name : ", name)
	fmt.Println("Vaccinated : ", isVaccinated)
	fmt.Println("Weight ", weight)
	fmt.Println("A ", a, " B ", b)
	fmt.Println("Inferred Type ", val)

	// multiple type declaration

	var (
		eyeColor   = "black"
		weights    = 30
		height     = 5.2
		bloodGroup = "A1B+ve"
	)
	fmt.Println("------------------Features------------------")
	fmt.Println(" Eye Color : ", eyeColor, "\n Weight: ", weights, "\n Height: ", height, "\n Blood Group: ", bloodGroup)

	// short hand declaration

	var1 := 10
	var2 := 25.6
	var3 := var1 + int(var2)

	fmt.Println("Value of var3 : ", var3)
}
