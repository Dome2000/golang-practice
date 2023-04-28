package main

import "fmt"

func main() {
	i := 1
	fmt.Println("i =", i)

	zerovalue(i)
	fmt.Println("i from function zerovalue =", i)

	// ต้องมี & เพื่อเข้าถึง address
	zeropointer(&i)
	fmt.Println("i value from function zeropointer =", i)

	fmt.Println("i address function zeropointer =", &i)
}

func zerovalue(ivalue int) {
	ivalue = 0
}

//ตัวแปรเราเป็น pointer ต้องใส่ *ตัวแปรด้วย
func zeropointer(ipointer *int) {
	//ตัวแปรที่เป็น pointer
	*ipointer = 0
}
