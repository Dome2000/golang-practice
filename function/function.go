package main

import "fmt"

func main() {
	hello()
	plus(1, 2)
	result := plusReturn(1, 3)
	fmt.Println("result with return =", result)

	result2 := plus3value(5,5,10)
	fmt.Println("result 3 value =", result2)
}

func hello() {
	fmt.Println("Hello function")
}

func plus(value1 int, value2 int) {
	result := value1 + value2
	fmt.Println("result =", result)
}

func plusReturn(value1 int, value2 int) int {
	return value1 + value2
}

func plus3value(value1, value2, value3 int) int {
	return value1+value2+value3
}
