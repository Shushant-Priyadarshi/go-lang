package main

import (
	"fmt"
	"time"
)

//channels helps goroutines to communicate with each other


func worker(ch chan string){
	fmt.Println("Inside Worker: Worker started...")
	time.Sleep(time.Second * 2)
	ch <- "Worker is working on ....."
}


func main(){
	ch := make(chan string)

	go worker(ch)

	workerMsg := <- ch

	fmt.Println("Main: ",workerMsg)
}
