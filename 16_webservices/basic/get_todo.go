package main

import (
	"fmt"
	"io"
	"net/http"
)

func main(){
	 res,err := http.Get("https://jsonplaceholder.typicode.com/todos")
	 if(err != nil){
		fmt.Println("Error fetching todos",err)
		return
	 }
	 defer res.Body.Close()

	 data,err := io.ReadAll(res.Body) //read till end of line
	 if(err != nil){
		fmt.Println("Error getting reading response",err)
		return
	 }

	 //fmt.Println(data) //byte data
	 fmt.Println(string(data)) //byte to string

	 
}