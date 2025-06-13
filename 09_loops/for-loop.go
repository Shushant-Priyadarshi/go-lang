package main

import "fmt"


func infiniteLoop(){
	counter := 0
	for{
		fmt.Println("Hello world")
		counter++
		if counter == 10 {return}
	}
}

func loop(){
	for i := 1 ; i <= 5 ;i++{
		fmt.Printf("i: %d\n",i)
	}
}

func rangeLoop(numbers []int){
	for index,value := range numbers {
		fmt.Printf("The Value at index %d is %d\n",index,value)
	}
}

func rangeLoopStr(str string){
	for index,value := range str {
		fmt.Printf("The Value at index %d is %c\n",index,value)
	}
}

func main(){
	numbers := []int{21,34,11,43}
	rangeLoop(numbers)

	myStr := "Shushant"
	rangeLoopStr(myStr)

	
}