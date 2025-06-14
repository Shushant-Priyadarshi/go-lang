package main

import "fmt"

func main() {
	var number int = 34
	var numberF float64 = float64(number)

	fmt.Printf("number type : %T\n", number)
	fmt.Printf("numberF type : %T\n", numberF)
}
