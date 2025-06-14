package main

import "fmt"

func passByValue( num int){
	num += 10
}

func passByRef(num *int){
	*num += 10
}

func main(){
	num := 10

	passByValue(num)
	fmt.Println("After Pass By Value:",num)

	passByRef(&num)
	fmt.Println("After Pass By Ref:",num)

}