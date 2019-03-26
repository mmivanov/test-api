package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"test-api/internal/model"
)

func main() {
	log.Print("Service started")
	router := mux.NewRouter()
	router.HandleFunc("/", GetIndex).Methods("GET")
	router.HandleFunc("/posts", GetPosts).Methods("GET")
	router.HandleFunc("/post/{id}", GetPost).Methods("GET")
	router.HandleFunc("/post", CreatePost).Methods("POST")
	router.HandleFunc("/post/{id}", UpdatePost).Methods("PUT")
	router.HandleFunc("/post/{id}", DeletePost).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func GetIndex(w http.ResponseWriter, r *http.Request) {
	log.Print("GetIndex")

}
func GetPosts(w http.ResponseWriter, r *http.Request) {
	var posts []model.Post
	log.Print("GetPosts and script updated")
	posts = append(posts, model.Post{ID: 1, Email: "test@text.ru", Text: "This is a simple text", Status: 0})
	posts = append(posts, model.Post{ID: 2, Email: "test2@mail.ru", Text: "This is a simple text", Status: 0})
	posts = append(posts, model.Post{ID: 3, Email: "test3@gmail.com", Text: "This is a simple text", Status: 0})
	posts = append(posts, model.Post{ID: 4, Email: "test4@gmail.com", Text: "This is a simple text", Status: 0})
	json.NewEncoder(w).Encode(posts)
}
func GetPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	log.Printf("GetPost: %s", params["id"])
}
func CreatePost(w http.ResponseWriter, r *http.Request) {}
func UpdatePost(w http.ResponseWriter, r *http.Request) {}
func DeletePost(w http.ResponseWriter, r *http.Request) {}
