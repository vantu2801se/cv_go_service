package api

import (
	"context/application/cvservice/models"
	"context/application/cvservice/services"
	"encoding/json"
	"net/http"
)

func ReadCVModel(w http.ResponseWriter, r *http.Request) {
	cvEntities, infoEntities := services.FindAllCVModel()
	results := []models.CVModel{}
	for i := 0; i < len(cvEntities); i++ {
		content := []models.Info{}
		for j := 0; j < len(infoEntities); j++ {
			if infoEntities[j].CvmodelId == cvEntities[i].Id {
				content = append(content,
					models.Info{
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
