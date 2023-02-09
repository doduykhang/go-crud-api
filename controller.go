package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const dns = "root:123456@tcp(127.0.0.1:3306)/articledb"

type Article struct {
	Id          string `json:"id`
	Title       string `json:"title`
	Description string `json:"description`
	Content     string `json:"content`
}

func InitialMigration() {
	DB, err = gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}
	DB.AutoMigrate(&Article{})
}

func GetArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var articles []Article
	DB.Find(&articles)
	json.NewEncoder(w).Encode(articles)
}

func GetArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var article Article
	DB.First(&article, params["id"])
	json.NewEncoder(w).Encode(article)
}

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var article Article
	json.NewDecoder(r.Body).Decode(&article)
	DB.Create(&article)
	json.NewEncoder(w).Encode(article)
}

func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var article Article
	DB.First(&article, params["id"])
	json.NewDecoder(r.Body).Decode(&article)
	DB.Save(&article)
	json.NewEncoder(w).Encode("Article is updated!")
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var article Article
	DB.Delete(&article, params["id"])
	json.NewEncoder(w).Encode("Article is deleted!")
}
