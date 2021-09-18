package service

import (
	"github.com/len3fun/money-tracker/internal/models"
	"github.com/len3fun/money-tracker/internal/repository"
)

type ActivityService struct {
	repo repository.Activity
}


func NewActivityService(repo repository.Activity) *ActivityService {
	return &ActivityService{repo: repo}
}

func (s *ActivityService) GetAllActivities(userId int) ([]models.ActivitiesOut, error) {
	return s.repo.GetAllActivities(userId)
}
