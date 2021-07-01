package api

import (
	"context/application/cvservice/models"
	"context/application/cvservice/repositories"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ReadCVModel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=ascii")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
	cvEntities, infoEntities := repositories.FindAllCVModel()
	results := []models.CVModel{}
	for i := 0; i < len(cvEntities); i++ {
		content := []models.Info{}
		for j := 0; j < len(infoEntities); j++ {
			if infoEntities[j].CvmodelId == cvEntities[i].Id {
				content = append(content,
					models.Info{
						Id:       infoEntities[j].Id,
						Text:     infoEntities[j].Txt,
						Progress: infoEntities[j].Progress})
			}
		}
		cvModel := models.CVModel{
			Id:      cvEntities[i].Id,
			Title:   cvEntities[i].Title,
			Content: content,
		}
		results = append(results, cvModel)
	}
	json.NewEncoder(w).Encode(results)
}

func UpdateCV(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var cv []models.CVModel
	json.Unmarshal(reqBody, &cv)
	repositories.UpdateCV(cv)
	json.NewEncoder(w).Encode(cv)
}
