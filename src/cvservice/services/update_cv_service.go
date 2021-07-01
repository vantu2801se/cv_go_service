package services

import (
	"context/application/cvservice/models"
	"context/application/cvservice/repositories"
)

func UpdateCV(cv []models.CVModel) {
	repositories.UpdateCV(cv)
}
