package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a := true
	fmt.Println("Value of A : ", a)
	var a1 = false
	fmt.Println("Value of A1 : ", a1)
	var b1 bool
	b1 = false
	fmt.Println("Not Operator : ", !b1)

	// logical operators -> &&, || , ! (return bool)
	// relational operator -> <,>,<=,>=,==,!= (return bool)
	// conditional operator -> ?:
	result := a && a1
	fmt.Println("Logical AND Result : ", result)

	output := b1 || a
	fmt.Println("Logical OR Result : ", output)

	op := !a
	fmt.Println("Logical NOT Result : ", op)

	val1, val2 := 234, 565
	val3 := val1 + val2
	fmt.Println("Result : ", val3)
	fmt.Printf("Type of %T", val1) // returns the type of variable using %T format specifier
	fmt.Println()
	fmt.Printf("Type of %T", a)
	fmt.Println()
	fmt.Printf("SizeOf of %d", unsafe.Sizeof(val3)) // returns the size of the variable in bytes

	var data1 float32
	data1 = 98.5
	data2 := 100.2
	fmt.Println()
	fmt.Printf("Type of data2 %T", data2)
	fmt.Println()
	data3 := data1 * float32(data2)
	fmt.Println("Val is ", data3)

	d1 := complex(10, 20) // complex number
	fmt.Println("Complex Value ", d1)
	d2 := 25.8 + 100.4i
	d3 := d1 + d2
	fmt.Println("Sum of Complex Value ", d3)

	// byte (uint8) and rune(int32)

	var str string
	str = "diyashri"
	str2 := "nagaraj"
	str3 := str + str2

	fmt.Println("Name : " + str3)
}
