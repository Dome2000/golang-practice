package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

// P1
type product struct {
	ProductName string
	Price       float64
	Weight      float64
}

// P2
func Fibonacci(x int) {
	a := 0
	b := 1
	c := b
	fmt.Printf("%d %d", a, b)
	for {
		c = b
		b = a + b
		if b >= x {
			fmt.Println("")
			break
		}
		a = c
		fmt.Printf(" %d", b)
	}
}

// P3
func getInput(promt string) float64 {
	fmt.Printf("%v", promt)
	input, _ := reader.ReadString('\n')
	value, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(input)
	} else if value == 0 {
		fmt.Println(input)
	}
	return value
}
func calcu(value1, value2 float64) {
	fmt.Println("Additon : ", value1+value2)
	fmt.Println("Subtraciton : ", value1-value2)
	fmt.Println("Multiplication : ", value1*value2)
	fmt.Println("Division : ", value1/value2)
}

// P5
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func main() {
	var product1 product
	product1.ProductName = "Iphone"
	product1.Price = 26900
	product1.Weight = 0.21

	fmt.Println("P1")
	fmt.Println("Product Name: ", product1.ProductName)
	fmt.Println("Price: ", product1.Price)
	fmt.Println("Weight: ", product1.Weight, " Kg")

	fmt.Println("P2")
	Fibonacci(10)

	fmt.Println("P3")
	value1 := getInput("Enter first number = ")
	value2 := getInput("Enter second  number = ")
	calcu(value1, value2)

	fmt.Println("P4")
	for i := 1; i < 13; i++ {
		for j := 1; j < 13; j++ {
			fmt.Print(i*j, " ")
		}
		fmt.Println("")
	}

	fmt.Println("P5")
	fmt.Println("LCM is: ", LCM(2, 4, 6, 8, 10))

}
