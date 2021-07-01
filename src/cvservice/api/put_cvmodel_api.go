package api

import (
	"context/application/cvservice/models"
	"context/application/cvservice/services"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func UpdateCV(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var cv []models.CVModel
	json.Unmarshal(reqBody, &cv)
	services.UpdateCV(cv)
	json.NewEncoder(w).Encode(cv)
}
