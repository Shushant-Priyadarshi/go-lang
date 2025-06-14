package main

import "fmt"


func  main(){
	
	var x int
	x = 23

	var ptrX *int
	ptrX = &x

	y := 34
	ptrY := &y

	fmt.Println("X: ",x)
	fmt.Println("Pointer X: ",ptrX)

	fmt.Println("Y: ",y)
	fmt.Println("Pointer Y: ",ptrY)
}