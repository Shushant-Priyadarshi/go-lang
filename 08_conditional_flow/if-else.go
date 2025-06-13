package main

import "fmt"

func main(){
	var age int
	fmt.Print("Please Enter your age: ")
	fmt.Scan(&age)

	if age <= 0 || age > 120{
		fmt.Println("Please enter a valid age")
	}else if age >= 18{
		fmt.Println("You are allowed to vote")
	}else{
		fmt.Println("You are not allowed to vote")
	}
}