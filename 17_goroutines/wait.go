package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	fmt.Printf("Worker %d ended...\n", i)
}

func main() {
	var wg sync.WaitGroup

	fmt.Println("Main function....")
	for i := 1; i <= 4; i++ {
		go worker(i, &wg)
		wg.Add(1)
	}

	wg.Wait()
}
