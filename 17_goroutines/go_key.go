package main

import (
	"fmt"
	"time"
)

func hello(){
	time.Sleep(time.Millisecond * 1000)
	fmt.Println("Hello")
	
}

func bonjjour(){

	time.Sleep(1500 * time.Millisecond)
	fmt.Println("Bonjour")
}

func main(){
	fmt.Println("Main function starting...")
	go hello()
	go bonjjour()
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("Main function ending...")
}
