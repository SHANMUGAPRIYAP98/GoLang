package main

import (
	"fmt"
)

func findSquare(a float64) float64 {
	return a * a
}

func traingleArea(b, h float64) float64 {
	return 0.5 * b * h
}

func bmiCalc(w, h float64) float64 {
	return w / (h * h)
}

func circle(rad float64) (float64, float64) {
	area := 3.14 * rad * rad
	peri := 2 * 3.14 * rad
	return area, peri
}

func printDetails(name string, weight, height float64) (fname string, bmi float64) {
	fmt.Println("First Name : ", name)
	bmi = weight / (height * height)
	return
}

func main() {
	squareRes := findSquare(25)
	fmt.Println("Square Result : ", squareRes)
	areaTriangleRes := traingleArea(5, 10)
	fmt.Println("Area of Triangle : ", areaTriangleRes)
	bmiRes := bmiCalc(55, 5.5)
	fmt.Println("BMI Calc : ", bmiRes)
	area, peri := circle(2.5)
	fmt.Print("\nArea of Circle : ", area, "\nPerimeter of Circle : ", peri)
	_, res := circle(5) // _ is blank identifier which is used to discard a variable
	fmt.Println()
	fmt.Println("Perimeter : ", res)
	val1, val2 := printDetails("diyashri", 22.5, 3.4)
	fmt.Println(val1)
	fmt.Println(val2)

}
