package main

import "fmt"


func main(){
	numbers := []int{0,1,2,3}
	

	fmt.Println(numbers)
	fmt.Printf("%T \n",numbers)

	numbers =append(numbers,4,5,6,7)
	fmt.Println(numbers)
}