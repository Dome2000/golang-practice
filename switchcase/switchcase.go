package main

import "fmt"


var input string
func main() {
	fmt.Scan(&input)

	switch input {
	case "blue":
		fmt.Println("#0000FF")
	case "green":
		fmt.Println("#008000")
	case "pink":
		fmt.Println("#FFC0CB")
	case "yellow":
		fmt.Println("#FFFF00")
	default:
		fmt.Println("invalid value")
	}

}
