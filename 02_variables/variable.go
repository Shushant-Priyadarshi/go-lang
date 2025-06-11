package main

import "fmt"

var MyName string = "Shushant"

func main(){
	
	var  age = 18
	MyName = "Rajneet"

	const PI float64 = 3.143245
	// PI = 34 error

	role := "SDE"
	fmt.Println("Hello,",MyName)
	fmt.Println("Your age is:",age)
	fmt.Println("Role: ",role)


}