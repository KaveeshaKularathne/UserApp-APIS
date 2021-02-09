package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"


)
type Article struct {
	Title string`json:"Title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
	
}
type Articles []Article

func allArticles(w http.ResponseWriter,r *http.Request)  {
	articles:=Articles{
		Article{Title:"Title",Desc: "Test Description",Content: "Hello"},
	}
	fmt.Println("Endpoint Hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)

}
func testPostArticles(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w,"Test POST Endpoint worked")

}

func homePage(w http.ResponseWriter,r *http.Request)   {
 fmt.Fprint(w,"Homepage")
}
func handleRequest()  {
	myRouter:=mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/",homePage)
	myRouter.HandleFunc("/articles",allArticles).Methods("GET")
	myRouter.HandleFunc("/articles",testPostArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":8082",myRouter))
}
func main()  {
	handleRequest()
}