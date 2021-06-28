package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// type singleton struct{
// 	multiplexer int
// }
// var instance *singleton
// func GetInstance() *singleton{
// 	if instance ==nil{
// 		instance=new(singleton)
// 	}
// 	return instance
// }
type Article struct {
	id   int
	name string
}
type Articles []Article

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Fprintf(w, "%+v", string(reqBody))
}
func readArticle(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Article{id: 1, name: "Hello world"})
}
func HandleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/articles", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/articles", createNewArticle).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	fmt.Println("Say Hello")
	HandleRequest()
}
