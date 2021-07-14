//We'll create a simple https server which handle HTTP requests.
//A Home page function that will handle all requests to our root URL.
//A Handle request function that will match the URL path 
//A main function that will act as an entry point for our server.
/*package main

import (
	"fmt"
	"net/http"
	"log"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func handleRequests(w http.ResponseWriter, r *http.Request) {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests(nil,nil)
}
//After this we'll create a our first API endpoint that returns a JSON response of a list of articles that we've to find in memory

//We need to define article struct that will hold our article information*/
package main

import (
	"fmt"
	"net/http"
	"log"
	"encoding/json"
)

type Article struct {
	Title string `json:"Title"`
	Desc string `json:"Description"`
	Content string `json:"Content"`
}

type Articles  []Article//now we can start mocking up some API endpoints which can head to retrive some data.

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles:= Articles{
		Article{Title:"First Article",Desc:"Article Description",Content:"Article Content"},
		//http://localhost:8080/articles
	}
	fmt.Fprintf(w, "Endpoint Hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func handleRequests(w http.ResponseWriter, r *http.Request) {
	http.HandleFunc("/", homePage)
	//We need to register our function 
	http.HandleFunc("/articles", allArticles)//We wanted to return with the function all articles.
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests(nil,nil)
}