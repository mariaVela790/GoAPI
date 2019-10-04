package main

import (
	"fmt"
)

// PART ONE /////////////////////////////////////////////////////////
//////Creating the base for our REST API
//////Function values may be used as function arguments and return values.
////func homePage(w http.ResponseWriter, r *http.Request){
////	fmt.Fprintf(w, "Welcome to the HomePage!")
////	fmt.Println("Endpoint Hit: homePage")
////}
////
////// Go functions may be closures. A closure is a function value that references variables from outside its body.
////// The function may access and assign to the referenced variables; in this sense the function is "bound" to the
////// variables.
////func handleRequests() {
////	http.HandleFunc("/", homePage)
////	// Fatal functions come from the package "log", works like Print, and calls os.Exit(1)
////	// after writing the log message
////	log.Fatal(http.ListenAndServe(":10000", nil))
////}
////
////func main() {
////	handleRequests()
////}
//
//// PART TWO ///////////////////////////////////////////////////////////
//// switch to mux by using `go get -u github.com/gorilla/mux`
//// We can use Go structs to define our Article structure. Structs are collections of fields
//// we can then access the fields using a dot
//type Article struct {
//	Id string `json:"Id"`
//	Title string `json:"Title"`
//	Desc string `json:"desc"`
//	Content string `json:"content"`
//}
//// The above is an example of type definition, which is where we use the type keyword
//// to create a new type (in this case a new type called Article) with the same type as the underlying structure
//
//// Using this struct, we can create a global array named Articles
//// consisting of Article structs to simulate a database
//var Articles []Article
//
//func homePage(w http.ResponseWriter, r *http.Request){
//	fmt.Fprintf(w, "Welcome to the HomePage!")
//	fmt.Println("Endpoint Hit: homePage")
//}
//
//func returnAllArticles(w http.ResponseWriter, r *http.Request) {
//	fmt.Println("Endpoint Hit: returnAllArticles")
//	json.NewEncoder(w).Encode(Articles)
//}
//
//func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
//	fmt.Println("Endpoint Hit: returnSingleArticle")
//	vars := mux.Vars(r)
//	key := vars["id"]
//
//	// printing just the article id
//	//fmt.Fprintf(w, "Key: " + key)
//
//	for _, article := range Articles {
//		if article.Id == key {
//			json.NewEncoder(w).Encode(article)
//		}
//	}
//
//}
//
//func createNewArticle(w http.ResponseWriter, r *http.Request) {
//	// We can test this by sending in a request body as raw data using post man
//	reqBody, _ := ioutil.ReadAll(r.Body)
//	//fmt.Fprintf(w, "%+v", string(reqBody))
//	var article Article
//	json.Unmarshal(reqBody, &article)
//	Articles = append(Articles, article)
//
//	json.NewEncoder(w).Encode(article)
//
//	// Test with the following new article
//	//{
//	//	"Id": "3",
//	//	"Title": "Newly Created Post",
//	//	"desc": "The description for my new post",
//	//	"content": "my articles content"
//	//}
//}
//
//func deleteArticle(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//
//	id := vars["id"]
//
//	// The range form of the for loop iterates over a slice or map
//	// When ranging over a slice, the index and the copy of the element
//	// at that index are returned. In this case, we use the index i but when
//	// we don't care about the index we can replace it with _. If we don't want to
//	// use the copy of the element at that index, we can just omit the variable
//	for i, article := range Articles {
//		if article.Id == id {
//			// Here we assign a slice to Articles that consists of Articles up to
//			// and not including i and Articles after i up to the length of Articles
//			// if the underlying array is to small to fit the new values, a bigger array
//			// will be allocated
//			Articles = append(Articles[:i], Articles[i+1:]...)
//		}
//	}
//	json.NewEncoder(w).Encode(Articles)
//
//
//	//	 Test with http://localhost:10000/article/2
//}
//
//func updateArticle(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//	id := vars["id"]
//
//	// this should only contain the new info but the id shouldn't be changed
//	reqBody, _ := ioutil.ReadAll(r.Body)
//	var a Article
//	// This method basically breaks down the request body and puts the info in the
//	// variable we are pointing to, in this case article
//	json.Unmarshal(reqBody, &a)
//
//	for i, article := range Articles {
//		if article.Id == id {
//			// once we find the matching Article we replace it with the new one
//			// in our slice of articles
//			Articles[i] = a
//		}
//	}
//	json.NewEncoder(w).Encode(Articles)
//
////	Test data
////{
//	//    "Id": "2",
//	//    "Title": "New Boo In Town",
//	//    "desc": "The description for my new post",
//	//    "content": "my articles content"
//	//}
//}
//
//func handleRequests() {
//	myRouter := mux.NewRouter().StrictSlash(true)
//
//	myRouter.HandleFunc("/", homePage)
//	myRouter.HandleFunc("/articles", returnAllArticles)
//
//	// We tack on Methods("POST") so that we call the createNewArticle function for POST requests only
//	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
//	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
//
//	myRouter.HandleFunc("/article/{id}", updateArticle).Methods("PUT")
//	myRouter.HandleFunc("/article/{id}", returnSingleArticle)
//	log.Fatal(http.ListenAndServe(":10000", myRouter))
//}
//
//func main() {
//	fmt.Println("Rest API v2.0 - Mux Routers")
//
//	// The following is a small example of Go pointers. A pointer holds the memory address of a value
//	// The type *T is a pointer to a T value. Its zero value is nil
//	// & generates a pointer to its operand * denotes the pointer's underlying value so we can assign values
//	// However, * is not necessary for get field values from pointers to structs because (*A).Id would be cumbersome
//	A := &Articles
//
//	// [n]T is an array of n values of type T. An array's length is part of its type so they can't be resized
//
//	// The following assigns an slice of type Article. A slice is a dynamically-sized, flexible view into the elements
//	// of an array by being references to arrays
//	*A = []Article{
//		Article{Id:"1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
//		Article{"2", "Hello 2", "Article Description", "Article Content"},
//	}
//	// The above notation of Article{...} is known as a struct literal which denotes a newly allocated struct value by listing
//	// values of its fields. Named fields can be in any order and be a subset of fields. & can be used in conjunction with a struct literal
//	// to return a pointer to the struct value
//
//	//The struct literals are in in a slice literal
//
//	//Usually slices have a lower and upper bound that includes the lower bound and excludes the upper. By default the
//	// bound is 0 to length of the slice
//
//	//capacity of a slice is the length of the underlying array, found by cap(s) where s is a slice
//	// slices can be resliced as long as the slice doesn't go past capacity. Slices have zero val of nil, length 0, cap 0
//	// Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.
//
//
//	handleRequests()
//}
//

// Extra topics ///////////////////////////////////////////////////////////////////////////////////////////////////////////
/////////////////// Methods and More on Pointers
// Methods - Go doesn't have classes but we can define methods on types
// A method is a function but it has a special receiver argument that is in its
// own argument list between the func keyword and method name
// The receiver argument can be of any type defined in the same package but not
// types from others including built in types like int

// Example
type Article struct {
	Id string `json:"Id"`
	Title string `json:"Title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func (a Article) getArticleTitle() string {
	return a.Title
}

// Broken Example
// This won't work because int is defined in another package
//func (a int) returnA() int {
//	return a
//}

// We can create non-struct types that can be used for the methods reciever argument
type Int int

func (a Int) returnInt() Int {
	return a
}

// Methods are just functions with a receiver argument

// Pointer Receivers
// We can also send in pointers to change the receiver itself
func (a *Article) changeTitleTo(title string) {
	a.Title = title
}

func main() {
	//a := Article{Id: "5", Title: "Title of a", Desc: "This is a desc", Content: "I'm where it's at"}
	//fmt.Println(a.getArticleTitle())

	////// Non struct type
	//b := Int(1)
	//fmt.Println("The following should be one...is it???")
	//fmt.Println(b.returnInt())

	//// Pointer receiver example
	art := Article{Id:"1", Title: "Hello", Desc: "Article Description", Content: "Article Content"}
	art.changeTitleTo("Hello World")
	// What happens when we take away the * from the def of the method
	fmt.Println(art)

}

//////////// Interfaces
//An interface type is defined as a set of method signatures.
//A value of interface type can hold any value that implements those methods.
//type Abser interface {
//	Abs() float64
//}
//
//func main() {
//	var a Abser
//	f := MyFloat(-math.Sqrt2)
//	v := Vertex{3, 4}
//
//	a = f  // a MyFloat implements Abser
//	a = &v // a *Vertex implements Abser
//
//	// In the following line, v is a Vertex (not *Vertex)
//	// and does NOT implement Abser.
//	//This example is broken because v is not a pointer and thus can't implement the method in the interface
//	//a = v
//
//	fmt.Println(a.Abs())
//}
//
//type MyFloat float64
//
//func (f MyFloat) Abs() float64 {
//	if f < 0 {
//		return float64(-f)
//	}
//	return float64(f)
//}
//
//type Vertex struct {
//	X, Y float64
//}
//
//func (v *Vertex) Abs() float64 {
//	return math.Sqrt(v.X*v.X + v.Y*v.Y)
//}
