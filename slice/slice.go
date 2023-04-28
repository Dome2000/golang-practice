package main

import "fmt"

func main() {
	var courseName []string
	courseName = []string{"jave", "Python"}
	fmt.Println(courseName)
	courseName = append(courseName, "C","HTML","CSS")
	fmt.Println(courseName)

	courseWeb := courseName[3:5]
	fmt.Println(courseWeb)

	courseWeb = courseName[:3]
	fmt.Println(courseWeb)
}