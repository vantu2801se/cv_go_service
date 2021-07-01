package repositories

import (
	"context/application/cvservice/entities"
	"log"
)

func FindAllCVModel() (entities.CVEntities, entities.InfoEntities) {
	db, err := GetMySqlConnection("root", "123qwe!@#", "127.0.0.10", 3306, "cv")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// GET DATA FROM CVMODEL TABLE
	var cvEntities entities.CVEntities
	queryCVEntities := "SELECT * FROM cvmodel"
	CVEntityResults, err := db.Query(queryCVEntities)
	if err != nil {
		log.Fatal(err)
	}
	for CVEntityResults.Next() {
		var cvEntity entities.CVEntity
		err := CVEntityResults.Scan(&cvEntity.Id, &cvEntity.Title)
		if err != nil {
			log.Fatal(err)
		}
		cvEntities = append(cvEntities, cvEntity)
	}
	// GET DATA FROM INFO TABLE
	var infoEntities entities.InfoEntities
	queryInfos := "SELECT * FROM info WHERE delete_flag=true"
	infoResults, err := db.Query(queryInfos)
	if err != nil {
		log.Fatal(err)
	}
	for infoResults.Next() {
		var info entities.InfoEntity
		err := infoResults.Scan(&info.Id, &info.Txt, &info.Progress, &info.CvmodelId, &info.DeleteFlag)
		if err != nil {
			log.Fatal(err)
		}
		infoEntities = append(infoEntities, info)
	}
	return cvEntities, infoEntities
}
