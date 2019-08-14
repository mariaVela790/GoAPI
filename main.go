package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

// PART ONE /////////////////////////////////////////////////////////
//// Creating the base for our REST API
//// Function values may be used as function arguments and return values.
//func homePage(w http.ResponseWriter, r *http.Request){
//	fmt.Fprintf(w, "Welcome to the HomePage!")
//	fmt.Println("Endpoint Hit: homePage")
//}
//
//// Go functions may be closures. A closure is a function value that references variables from outside its body.
//// The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.
////For example, the adder function returns a closure. Each closure is bound to its own sum variable.
//func handleRequests() {
//	http.HandleFunc("/", homePage)
//	// Fatal functions come from the package "log", works like Print, and calls os.Exit(1)
//	// after writing the log message
//	log.Fatal(http.ListenAndServe(":10000", nil))
//}
//
//func main() {
//	handleRequests()
//}

// PART TWO ///////////////////////////////////////////////////////////
// We can use Go structs to define our Article structure. Structs are collections of fields
// we can then access the fields using a dot
type Article struct {
	Id string `json:"Id"`
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

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnSingleArticle")
	vars := mux.Vars(r)
	key := vars["id"]

	// printing just the article id
	//fmt.Fprintf(w, "Key: " + key)

	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}

}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	// We can test this by sending in a request body as raw data using post man
	reqBody, _ := ioutil.ReadAll(r.Body)
	//fmt.Fprintf(w, "%+v", string(reqBody))
	var article Article
	json.Unmarshal(reqBody, &article)
	Articles = append(Articles, article)

	json.NewEncoder(w).Encode(article)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	// The range form of the for loop iterates over a slice or map
	// When ranging over a slice, the index and the copy of the element
	// at that index are returned. In this case, we use the index i but when
	// we don't care about the index we can replace it with _. If we don't want to
	// use the copy of the element at that index, we can just omit the variable
	for i, article := range Articles {
		if article.Id == id {
			// Here we assign a slice to Articles that consists of Articles up to
			// and not including i and Articles after i up to the length of Articles
			// if the underlying array is to small to fit the new values, a bigger array
			// will be allocated
			Articles = append(Articles[:i], Articles[i+1:]...)
		}
	}
}

//func updateArticle(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//	id := vars["id"]
//
//	reqBody, _ := ioutil.ReadAll(r.Body)
//
//}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", returnAllArticles)
	// We tack on Methods("POST") so that we call the createNewArticle function for POST requests only
	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	//myRouter.HandleFunc("/article/{id}", updateArticle).Methods("PUT")
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	// The following is an example use of shorthand syntax := for declaring and initiating variables which
	// can only be used in functions
	fmt.Println("Rest API v2.0 - Mux Routers")

	// The following is a small example of Go pointers. A point holds the memory address of a value
	// The type *T is a pointer to a T value. Its zero value is nil
	// & generates a pointer to its operand * denotes the pointer's underlying value so we can assign values
	// However, * is not necessary for get field values from pointers to structs because (*A).Id would be cumbersome
	A := &Articles

	// [n]T is an array of n values of type T. An array's length is part of its type so they can't be resized

	// The following assigns an slice of type Article. A slice is a dynamically-sized, flexible view into the elements
	// of an array by being references to arrays
	*A = []Article{
		Article{Id:"1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{"2", "Hello 2", "Article Description", "Article Content"},
	}
	// The above notation of Article{...} is known as a struct literal which denotes a newly allocated struct value by listing
	// values of its fields. Named fields can be in any order and be a subset of fields. & can be used in conjunction with a struct literal
	// to return a pointer to the struct value

	//The struct literals are in in a slice literal

	//Usually slices have a lower and upper bound that includes the lower bound and excludes the upper. By default the
	// bound is 0 to length of the slice

	//capacity of a slice is the length of the underlying array, found by cap(s) where s is a slice
	// slices can be resliced as long as the slice doesn't go past capacity. Slices have zero val of nil, length 0, cap 0
	// Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.
	handleRequests()
}
