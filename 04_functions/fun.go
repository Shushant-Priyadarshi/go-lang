package main

import "fmt"


func add(a int ,b int) (int) {
	return a+b
}

func greetUser(name string) (string){
	return "Hello, "+name
}

func main(){
	//Greet("shushant")
	fmt.Println("5 + 3 =",add(5,3))
	fmt.Println(greetUser("shushant"))
}