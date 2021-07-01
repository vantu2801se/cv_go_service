package repositories

import (
	"context/application/cvservice/entities"
	"context/application/cvservice/models"
	"fmt"
	"log"
)

func UpdateCV(cv []models.CVModel) {
	cvEntities := []entities.CVEntity{}
	infoEntities := []entities.InfoEntity{}
	for i := 0; i < len(cv); i++ {
		cvEntities = append(cvEntities, entities.CVEntity{
			Id:    cv[i].Id,
			Title: cv[i].Title,
		})
		if len(cv[i].Content) > 0 {
			for j := 0; j < len(cv[i].Content); j++ {
				infoEntities = append(infoEntities, entities.InfoEntity{
					Id:         cv[i].Content[j].Id,
					Progress:   cv[i].Content[j].Progress,
					Txt:        cv[i].Content[j].Text,
					CvmodelId:  cv[i].Id,
					DeleteFlag: false,
				})
			}
		}
	}
	fmt.Printf("%v", cvEntities)
	fmt.Println("")
	fmt.Printf("%v", infoEntities)
	db, err := GetMySqlConnection("root", "123qwe!@#", "127.0.0.10", 3306, "cv")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Println(cvEntities[0].Title)

	for i := 0; i < len(cvEntities); i++ {
		query := fmt.Sprintf("UPDATE cvmodel SET title='%s' WHERE id=%d;", cvEntities[i].Title, cvEntities[i].Id)
		update, err := db.Query(query)
		if err != nil {
			log.Fatal(err)
		}
		update.Close()
	}

	for i := 0; i < len(infoEntities); i++ {
		query := fmt.Sprintf("UPDATE info SET txt='%s',progress=%d WHERE id=%d;", infoEntities[i].Txt, infoEntities[i].Progress, infoEntities[i].Id)
		update, err := db.Query(query)
		if err != nil {
			log.Fatal(err)
		}
		update.Close()
	}

}
