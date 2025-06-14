package main

import "fmt"

func main(){
	name := "Shushant"
	ptrName := &name

	fmt.Println("Name: ",name)
	fmt.Println("Pointer Name: ",ptrName)

	fmt.Println("Value at pointer name: ", *ptrName)
}