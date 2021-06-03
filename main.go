package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

/**
 * @Description: Declare global articles array that we can populate in
 * our main function to simulate a database
 */
var Articles []Article

func readSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func readAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint his: readAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	gorillaRouter := mux.NewRouter().StrictSlash(true)

	gorillaRouter.HandleFunc("/", homePage)
	gorillaRouter.HandleFunc("/article/{id}", readSingleArticle)
	gorillaRouter.HandleFunc("/all", readAllArticles)
	log.Fatal(http.ListenAndServe(":3000", gorillaRouter))
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	//populate Articles with dummy data
	Articles = []Article{
		Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Id: "2", Title: "Hello 2", Desc: "Article Description 2", Content: "Article Content 2"},
	}

	handleRequests()
}
