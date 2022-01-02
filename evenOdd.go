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

func evaluateBool() {
	// assignment the value and evaluating the condition
	if cond := true; cond == true {
		fmt.Println("I am in true Block")
		fmt.Println(&cond)
	} else {
		fmt.Println("I am in false Block", cond)
	}

	//	fmt.Println(cond) cond variable can't be accessed outside if... else block
}

func findLargestNumber(num1, num2, num3 int) int {
	if num1 > num2 && num1 > num3 {
		return num1
	} else if num2 > num1 && num2 > num3 {
		return num2
	} else {
		return num3
	}
}

func main() {

	evenOdd(20)
	posNegNum(12)
	posNegNum(0)
	posNegNum(-10)
	printStudDetails("diyashri", 85.5, 98.3, 87.5)
	evaluateBool()

	a, b, c := 22, 50, 45
	largestNumber := findLargestNumber(a, b, c)
	fmt.Printf("%d is the largest number of %d,%d,%d", largestNumber, a, b, c)
}
