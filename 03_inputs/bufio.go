package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){

	fmt.Print("Please Enter your name: ")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Printf("Your name is %s",name)
}