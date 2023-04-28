package main

import "fmt"

var numberInt, numberInt2 int = 1000, 2000
var text string = "Hello"

func main() {
	// ประกาศตัวแปรแบบสั้น
	numberfloat := 25.4

	fmt.Println(numberInt)
	fmt.Println(numberInt2)
	fmt.Println(numberfloat)

	fmt.Println(text)

	fmt.Println(numberInt + numberInt2)
	fmt.Println(float64(numberInt) + numberfloat)
	fmt.Println(text+"World")
	fmt.Println("Sum = ",numberInt," บาท")
}

