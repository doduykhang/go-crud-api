package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func initializeRouter() {
	r := mux.NewRouter()

	r.HandleFunc("/articles", GetArticles).Methods("GET")
	r.HandleFunc("/article/{id}", GetArticle).Methods("GET")
	r.HandleFunc("/article", CreateArticle).Methods("POST")
	r.HandleFunc("/article/{id}", UpdateArticle).Methods("PUT")
	r.HandleFunc("/article/{id}", DeleteArticle).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", r))
}

func main() {
	InitialMigration()
	initializeRouter()
}
