package repositories

import (
	"context/application/cvservice/entities"
	"fmt"
	"log"
)

func FindAllCVModel() (entities.CVEntities, entities.InfoEntities) {
	db, err := GetPosgresqlConnection("rzipymtodmnerm", "fabfcfd23551517f5af39c105e53f3f447b718ac0cc4694663b37ca98da5ae30", "ec2-34-193-113-223.compute-1.amazonaws.com", 5432, "d83gnn9lniar5o")
	// connStr := "postgres://rzipymtodmnerm:fabfcfd23551517f5af39c105e53f3f447b718ac0cc4694663b37ca98da5ae30@ec2-34-193-113-223.compute-1.amazonaws.com:5432/d83gnn9lniar5o?sslmode=require"
	// db, err := sql.Open("postgres", connStr)
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
	defer CVEntityResults.Close()
	for CVEntityResults.Next() {
		var cvEntity entities.CVEntity
		err := CVEntityResults.Scan(&cvEntity.Id, &cvEntity.Title)
		if err != nil {
			log.Fatal(err)
		}
		cvEntities = append(cvEntities, cvEntity)
	}
	fmt.Printf("%v", cvEntities)
	// GET DATA FROM INFO TABLE
	var infoEntities entities.InfoEntities
	queryInfos := "SELECT * FROM info WHERE delete_flag=true"
	infoResults, err := db.Query(queryInfos)
	if err != nil {
		log.Fatal(err)
	}
	defer infoResults.Close()
	for infoResults.Next() {
		var info entities.InfoEntity
		err := infoResults.Scan(&info.Id, &info.Txt, &info.Progress, &info.CvmodelId, &info.DeleteFlag)
		if err != nil {
			log.Fatal(err)
		}
		infoEntities = append(infoEntities, info)
	}
	fmt.Printf("%v", infoEntities)
	return cvEntities, infoEntities
}
