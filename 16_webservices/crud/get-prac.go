package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)



func main(){
	url_str := "https://api.freeapi.app/api/v1/public/randomjokes/1"

	//paring the string url to url type for id extraction
	parsed_url,err := url.Parse(url_str)
	if err != nil{
		fmt.Println("Error Paring the url string",err)
		return
	}
	
	//extracting the id from the parsel string url
	jokePath := strings.Split(parsed_url.Path,"/")
	jokeId := jokePath[len(jokePath)-1]

	res,err := http.Get(url_str)
	if(err != nil){
		fmt.Println("Error Fetching API: ", err)
		return
	}

	defer res.Body.Close()

	//handle fetching error
	if(res.StatusCode != http.StatusOK){
		var errorResponse ErrorResponse
		fmt.Println("API returned non-200 status code")
		//parsing error to ErrorResponse struct
		if err := json.NewDecoder(res.Body).Decode(&errorResponse)
		err != nil{
			fmt.Println("Error Response: ",err)
			return
		}
		fmt.Printf("API Error: %s\n", errorResponse.Message)
		return
	}

	//parsing if there is not fetching error 
	var jokeResponse JokeResponse
	if err := json.NewDecoder(res.Body).Decode(&jokeResponse)
	err != nil{
		fmt.Println("Error Parsing the json data",err)
		return
	}

	//printing after parsing
	fmt.Printf("Joke at id %s : %s",jokeId,jokeResponse.Data.Content)
	
}


type ErrorResponse struct {
	StatusCode int      `json:"statusCode`
	Data       string   `json:"data"`
	Success    bool     `json:"success"`
	Errors     []string `json:"errors"`
	Message    string   `json:"message"`
}

type JokeResponse struct {
	StatusCode int    `json:"statusCode`
	Data       Joke   `json:"data"`
	Message    string `json:"message"`
	Success    bool   `json:"success"`
}

type Joke struct {
	Categories []string `json:"categories"`
	ID         int      `json:"id"`
	Content    string   `json:"content"`
}