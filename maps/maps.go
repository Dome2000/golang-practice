package main

import "fmt"

var product = make(map[string]float64)

func main() {
	fmt.Println("product =", product)

	//add
	product["Macbook"] = 40000
	product["iPhone"] = 30000
	product["iPad"] = 25000
	fmt.Println("product =", product)

	//delete
	delete(product, "iPad")
	fmt.Println("product =", product)

	//update
	product["iPhone"] = 20000
	fmt.Println("product =", product)

	value1 := product["iPhone"]
	fmt.Println("value1 =", value1)

	courseName := map[string]string{"101": "jave", "102": "python"}
	fmt.Println("courseName =", courseName)
}
