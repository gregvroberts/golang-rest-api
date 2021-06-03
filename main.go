package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
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

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	//get body of post request
	//return the string response containing the request body
	requestBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(requestBody, &article)

	Articles = append(Articles, article)

	json.NewEncoder(w).Encode(article)
	//fmt.Fprintf(w, "%+v", string(requestBody))

}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	// once again, we will need to parse the path parameters
	vars := mux.Vars(r)
	// we will need to extract the `id` of the article we
	// wish to delete
	id := vars["id"]

	// we then need to loop through all our articles
	for index, article := range Articles {
		// if our id path parameter matches one of our
		// articles
		if article.Id == id {
			// updates our Articles array to remove the
			// article
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}

}

//todo
//func updateArticle(w http.ResponseWriter, r *http.Request) {
//	requestBody, _ := ioutil.ReadAll(r.Body)
//	vars := mux.Vars(r)
//	id := vars["id"]
//	var article Article
//	json.Unmarshal(requestBody, &article)
//	for index, article := range Articles {
//		if article.Id == id {
//
//		}
//	}
//}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	gorillaRouter := mux.NewRouter().StrictSlash(true)
	//HomePage
	gorillaRouter.HandleFunc("/", homePage)
	// return all articles
	gorillaRouter.HandleFunc("/articles", readAllArticles)
	//create a single article
	//MUST BE BEFORE THE OTHER /article ENDPOINT
	gorillaRouter.HandleFunc("/article", createNewArticle).Methods("POST")
	// delete article
	gorillaRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	//return a single article
	gorillaRouter.HandleFunc("/article/{id}", readSingleArticle)
	// Start up server
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
