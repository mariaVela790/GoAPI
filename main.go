package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// PART ONE /////////////////////////////////////////////////////////
// Creating the base for our REST API

//func homePage(w http.ResponseWriter, r *http.Request){
//	fmt.Fprintf(w, "Welcome to the HomePage!")
//	fmt.Println("Endpoint Hit: homePage")
//}
//
//func handleRequests() {
//	http.HandleFunc("/", homePage)
//	log.Fatal(http.ListenAndServe(":10000", nil))
//}
//
//func main() {
//	handleRequests()
//}

// PART TWO ///////////////////////////////////////////////////////////
// We can use Go structs to define our Article structure
type Article struct {
	Title string `json:"Title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

// Using this struct, we can create a global array named Articles
// consisting of Article structs to simulate a database
var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	// This is the second route we add to use the function returnAllArticles
	http.HandleFunc("/articles", returnAllArticles)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func main() {
	// The following is an example use of shorthand syntax := for declaring and initiating variables which
	// can only be used in functions
	Articles = []Article{
		Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	handleRequests()
}
