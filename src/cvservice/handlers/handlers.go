//Multiplexer buit-by Van Tu
package handlers

import (
	"context/application/cvservice/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type CVModel struct {
	Id      int
	Title   string
	Content []string
}
type CVModels []CVModel

func HandleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/cvmodel", createCVModel).Methods("POST")
	myRouter.HandleFunc("/cvmodel", readCVModel).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func createCVModel(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Fprintf(w, "%+v", string(reqBody))
}
func readCVModel(w http.ResponseWriter, r *http.Request) {
	var cv models.CVModel
	cv.Setter(1, "b", []string{"a", "b", "c"})
	json.NewEncoder(w).Encode(cv)
}
