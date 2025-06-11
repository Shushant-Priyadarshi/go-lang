package main

import "fmt"

func main(){
	var name = "Shushant"
	age := 19

	fmt.Println("name: ",name)
	fmt.Print("age: ",age)
	fmt.Println("Heelo")

	fmt.Printf("Name is %s and age is %d",name,age )
	fmt.Println()
	fmt.Printf("Type of name is %T and age is %T",name,age)
}