package main

import (
    "fmt"
    "sync"
)

var counter int
var mu sync.Mutex // Create a Mutex

func increment() {
    mu.Lock()   // Lock before accessing shared data
    counter++
   mu.Unlock() // Unlock after done
}

func main() {
    var wg sync.WaitGroup

    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            increment()
        }()
    }

    wg.Wait()
    fmt.Println("Final Counter:", counter)
}

