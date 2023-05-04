package main

import (
	"encoding/json"
	"fmt"
)

type employee struct {
	Id           int
	EmployeeName string
	Tel          string
	Email        string
}

func main() {
	// json.Marshal อ่านแค่key ตัวพิมใหญ่
	data, _ := json.Marshal(&employee{101, "peeraphat charoenะ้ฟร", "1111111111111", "dome@gmail.com"})
	fmt.Print(string(data))
}
