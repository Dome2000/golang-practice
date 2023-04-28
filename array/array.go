package main

import "fmt"

var productName [4]string
var price [4]float32

func main() {
	productName[0] = "Mac"
	productName[1] = "Mac2"
	productName[2] = "Mac3"
	productName[3] = "Mac3"

	price := [4]float32{100,200,300,400}

	fmt.Println(productName)
	fmt.Println(price)
}