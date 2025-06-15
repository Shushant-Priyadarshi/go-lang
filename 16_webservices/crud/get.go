// package main

// import (
// 	"encoding/json"
// 		fmt "fmt"
// 	"io"
// 	"net/http"
// )

// // Define structures matching your API's response format
// type JokeResponse struct {
// 	StatusCode int    `json:"statusCode"`
// 	Data       Joke  `json:"data"`
// 	Message    string `json:"message"`
// 	Success    bool   `json:"success"`
// }

// type Joke struct {
// 	Categories []string `json:"categories"`
// 	ID         int      `json:"id"`
// 	Content    string   `json:"content"`
// }

// func main() {
// 	url := "https://api.freeapi.app/api/v1/public/randomjokes/2323423"

// 	// Perform the GET request
// 	res, err := http.Get(url)
// 	if err != nil {
// 		fmt.Println("Error fetching joke:", err)
// 		return
// 	}
// 	defer res.Body.Close()

// 	// Check for non-200 status codes first
// 	if res.StatusCode != http.StatusOK {
// 		fmt.Printf("API returned non-200 status code: %d\n", res.StatusCode)
// 		body, _ := io.ReadAll(res.Body)
// 		fmt.Printf("Response body: %s\n", body)
// 		return
// 	}

// 	// Parse the JSON into our struct
// 	var jokeResponse JokeResponse
// 	if err := json.NewDecoder(res.Body).Decode(&jokeResponse)
// 	err != nil {
// 	fmt.Println("Error parsing json:", err)
// 		return
// 	}

// 	// Access the parsed data safely
// 	fmt.Printf("Joke: %s\n", jokeResponse.Data.Content)
// }
