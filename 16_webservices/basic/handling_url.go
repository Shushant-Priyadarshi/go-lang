package main

import (
	"fmt"
	"net/url"
)

func main(){
	myUrl := "https://jsonplaceholder.typicode.com/todos/1?key=12&search=hello"
	fmt.Printf("Type: %T\n",myUrl)

	parsed_utl,err := url.Parse(myUrl)
	if(err != nil){
		return
	}
	fmt.Printf("Type Parsed: %T\n",parsed_utl)

	fmt.Println("Scheme: ",parsed_utl.Scheme)
	fmt.Println("Host: ",parsed_utl.Host)
	fmt.Println("Path: ",parsed_utl.Path)
	fmt.Println("RawQuery: ",parsed_utl.RawQuery)
}