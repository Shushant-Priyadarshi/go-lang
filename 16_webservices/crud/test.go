package main

import (
 "strings"
 "fmt"
)

func main() {
	str := "https://api.freeapi.app/api/v1/public/randomjokes/2"
	//some := strings.Trim(str, "/")
	somemore := strings.Split(str,"/")
	fmt.Println(somemore[len(somemore) -1])
}