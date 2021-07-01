//Multiplexer buit-by Van Tu
package handlers

import (
	"context/application/cvservice/api"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/cvmodel", api.UpdateCV).Methods("POST")
	myRouter.HandleFunc("/cvmodel", api.UpdateCV).Methods("PUT")
	myRouter.HandleFunc("/cvmodel", api.ReadCVModel).Methods("GET")
	fmt.Println("GO SERVER IS RUNNING. . . ")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func createCVModel(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Fprintf(w, "%+v", string(reqBody))
}
