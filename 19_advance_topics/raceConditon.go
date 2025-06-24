package main

import (
	"fmt"
	"sync"
)

var counter int
var mu sync.Mutex

func increment(wg *sync.WaitGroup){

	mu.Lock()
	counter++
	mu.Unlock()
	defer wg.Done()
}

func main(){

	var wg sync.WaitGroup

	for i:= 1 ; i<=1000;i++{
		wg.Add(1)
		go increment(&wg)
	}

	wg.Wait()
	fmt.Println("Counter: ", counter)
}
