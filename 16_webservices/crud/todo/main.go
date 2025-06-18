package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Todo struct {
	Id        int    `json:"id"`
	Todo      string `json:"todo"`
	Completed bool   `json:"completed"`
	UserID    int    `json:"userId"`
}

// get request
func getTodo(id string) {
	url := "https://dummyjson.com/todos/"+id

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching API: ", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK{
		fmt.Println("API returned non-200 status")
		return
	}

	var todo Todo
	if err := json.NewDecoder(res.Body).Decode(&todo)
	err != nil{
		fmt.Println("Error Parsing the data: ",err)
	}
	fmt.Println("GET Response: ",todo)
}

// post request
func postTodo() {
	myTodo := Todo{
		Todo: "YOYO",
    	Completed: false,
    	UserID: 32,
	}

	//convert this myTodo struct to json
	jsonData,err := json.Marshal(myTodo)
	if err != nil{
		fmt.Println("Error converting struct to json: ",err)
		return
	}


	//calling the post the request
	res,err := http.Post("https://dummyjson.com/todos/add","application/json",bytes.NewReader(jsonData))
	if err != nil{
		fmt.Println("Error sending Post request: ",err)
		return
	}

	defer res.Body.Close()

	data,err := io.ReadAll(res.Body)
	if err != nil{
		fmt.Println("Error Reading the byte response data")
	}
	fmt.Println("Post ResPonse: ",string(data))

}


func updateTodo() {
	updatedTodo := Todo{
		Completed: false,
	}

	url := "https://dummyjson.com/todos/2"

	//todo struct => json by marshalling
	jsonData,err :=json.Marshal(updatedTodo)
	if err != nil{
		fmt.Println("Error converting the data to json", err)
		return
	}

	//Creating the put request
	req,err := http.NewRequest(http.MethodPut,url,bytes.NewReader(jsonData))
	if err != nil{
		fmt.Println("Error creating the put request:",err)
		return
	}
	req.Header.Set("Content-type","application/json")

	//Sending the created request
	client := http.Client{}
	res,err := client.Do(req)
	if err != nil{
		fmt.Println("Error Sending the put request: ",err)
		return
	}
	defer res.Body.Close()

	data,err :=io.ReadAll(res.Body)
	if err!=nil{
		fmt.Println("Error converting the http response data into byte format")
	} 
	fmt.Println("PUT Response: ",string(data))
}

func deleteTodo() {
	baseUrl := "https://dummyjson.com/todos/2"

	req,err := http.NewRequest(http.MethodDelete,baseUrl,nil)
	if err != nil{
		fmt.Println("Error Creating put request:",err)
		return
	}

	client := http.Client{}
	res,err := client.Do(req)
	if err != nil{
		fmt.Println("Error sending put request: ",err)
		return
	}
	defer res.Body.Close()

	data,err := io.ReadAll(res.Body)
	if err!=nil{
		fmt.Println("Error converting the http response to byte data:",err)
		return
	}
	fmt.Println("DELETE RESPONSE: ",string(data))

}

func main() {
	//getTodo("2")
	//postTodo()
	//updateTodo()
	deleteTodo()
}
