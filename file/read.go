package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	flie, err := os.Open("indexinfo.csv")
	if err != nil {
		panic(err)
	}
	count := 1
	scanner := bufio.NewScanner(flie)
	for scanner.Scan() {
		line := scanner.Text()
		item := strings.Split(line, ",")

		fmt.Println(item[3])
		// fmt.Printf("Line %d : %s \n", count, line)
		count++
	}
}
