package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/len3fun/money-tracker/internal/models"
	"github.com/len3fun/money-tracker/pkg/logger"
)

type ActivityRepository struct {
	db *sqlx.DB
}

func NewActivityRepository(db *sqlx.DB) *ActivityRepository {
	return &ActivityRepository{db: db}
}

func (r *ActivityRepository) GetAllActivities(userId int) ([]models.ActivitiesOut, error) {
	var activities []models.ActivitiesOut
	query := fmt.Sprintf("SELECT change, activity_type, label, activity_date " +
		"FROM %s " +
		"WHERE user_id = $1", activitiesTable)
	logger.Debug(query)
	err := r.db.Select(&activities, query, userId)
	return activities, err
}

