package main

import "fmt"

func main(){
	var names[5] string
	names[0] = "one"
	names[1] = "two"
	names[2] = "three"

	numbers := []int {1,2,3,4,5}

	fmt.Println(names)
	fmt.Println(names[4])
	fmt.Println(numbers)
	fmt.Println("Length is: ", len(names))
	



}