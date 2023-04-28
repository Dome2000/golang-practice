package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {

	value1 := getInput("value 1 = ")
	value2 := getInput("value 2 = ")

	var result float64
	switch operator := getOperator(); operator {
	case "+":
		result = add(value1, value2)
	case "-":
		result = subtract(value1, value2)
	case "*":
		result = mulitiply(value1, value2)
	case "/":
		result = divide(value1, value2)
	default:
		panic("invalid operator")
	}
	fmt.Printf("result is %v", result)
}
func getInput(promt string) float64 {
	fmt.Printf("%v", promt)
	input, _ := reader.ReadString('\n')
	value, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println("check" + input)
	}
	return value
}
func getOperator() string {
	fmt.Print("operator is (+ - * /) : ")
	op, _ := reader.ReadString('\n')
	return strings.TrimSpace(op)
}

func add(value1, value2 float64) float64 {
	return value1 + value2
}

func subtract(value1, value2 float64) float64 {
	return value1 - value2
}

func mulitiply(value1, value2 float64) float64 {
	return value1 * value2
}

func divide(value1, value2 float64) float64 {
	return value1 / value2
}
