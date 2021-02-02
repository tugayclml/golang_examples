package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

type Article struct {
	Title string `json:"Title"`
	Desc string `json:"Desc"`
	Content string `json:"Content"`
}

type Articles []Article

func getAllArticles(w http.ResponseWriter, r *http.Request)  {
	articles := Articles{
		Article{Title: "Test title", Desc: "Test Description", Content: "Test Content"},
	}
	fmt.Println("Endpoint hit all articles")
	json.NewEncoder(w).Encode(articles)

}

func testPostArticles(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Test post hit")
}


func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Home page endpoint hit!")
}

func handleRequest()  {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage)
	router.HandleFunc("/articles", getAllArticles).Methods("GET")
	router.HandleFunc("/articles", testPostArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", router))
}

func main(){
	handleRequest()
}
