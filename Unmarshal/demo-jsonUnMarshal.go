package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type employee2 struct {
	Id           int
	EmployeeName string
	Tel          string
	Email        string
}

func main() {

	e := employee2{}
	err := json.Unmarshal([]byte(`{"ID":101,"EmployeeName":"Dome","Tel":"2222222","Email":"dome@gmail.com"}`), &e)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(e.EmployeeName)
	
}
