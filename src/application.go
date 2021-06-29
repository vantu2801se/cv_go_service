package main

import (
	"context/application/cvservice/handlers"
)

// type singleton struct {
// 	multiplexer mux.Route
// }

// var instance *singleton

// func GetInstance() *singleton {
// 	if instance == nil {
// 		instance = new(singleton)
// 	}
// 	return instance
// }

// type Article struct {
// 	id   int
// 	name string
// }
// type Articles []Article

// func createNewArticle(w http.ResponseWriter, r *http.Request) {
// 	reqBody, _ := ioutil.ReadAll(r.Body)
// 	fmt.Fprintf(w, "%+v", string(reqBody))
// }
// func readArticle(w http.ResponseWriter, r *http.Request) {
// 	json.NewEncoder(w).Encode(Article{id: 1, name: "Hello world"})
// }
// func HandleRequest() {
// 	myRouter := mux.NewRouter().StrictSlash(true)
// 	myRouter.HandleFunc("/articles", createNewArticle).Methods("POST")
// 	myRouter.HandleFunc("/articles", readArticle).Methods("GET")
// 	log.Fatal(http.ListenAndServe(":8080", myRouter))
// }

func main() {
	handlers.HandleRequest()
}
