package main

import (
	"fmt"
	"math"
)

func evenOdd(num int) {
	if num%2 == 0 {
		fmt.Printf("%d number is a even number", num)
		fmt.Println()
	} else {
		fmt.Printf("%d number is a odd number", num)
		fmt.Println()
	}
}

func posNegNum(val int) {
	if val > 0 {
		fmt.Println(val, " is a positive number")
	} else if val < 0 {
		fmt.Println(val, " is a negative number")
	} else {
		fmt.Println(val, " is a neutral value")
	}
}

func printStudDetails(name string, num1, num2, num3 float64) {

	total := num1 + num2 + num3
	fmt.Println("Student Name : " + name)
	fmt.Println("Total Marks : ", total)
	avg := total / 3
	fmt.Println("Average : ", math.Round(avg))
	var grade string
	if avg >= 90 && avg < 100 {
		grade = "O"
	} else if avg >= 80 && avg < 90 {
		grade = "A"
	} else if avg >= 70 && avg < 80 {
		grade = "B"
	} else {
		grade = "E"
	}
	fmt.Println("Grade : " + grade)
}

func main() {

	evenOdd(20)
	posNegNum(12)
	posNegNum(0)
	posNegNum(-10)
	printStudDetails("diyashri", 85.5, 98.3, 87.5)
}
