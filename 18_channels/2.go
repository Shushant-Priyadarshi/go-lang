package main

import (
	"fmt"
	"math/rand"
	"time"
)

func processNum(numChan chan int) {
	for num := range numChan {
		fmt.Println("Processing Number: ", num)
		time.Sleep(time.Second)
	}

}

func main() {
	chanInt := make(chan int)

	go processNum(chanInt)
	
	for {
		chanInt <- rand.Intn(100)
	}
}
