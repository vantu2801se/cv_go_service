package services

import (
	"context/application/cvservice/entities"
	"context/application/cvservice/repositories"
)

func FindAllCVModel() (entities.CVEntities, entities.InfoEntities) {
	return repositories.FindAllCVModel()
}
