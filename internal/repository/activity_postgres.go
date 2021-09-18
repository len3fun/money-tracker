package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/len3fun/money-tracker/internal/models"
	"github.com/len3fun/money-tracker/pkg/logger"
)

type ActivityPostgres struct {
	db *sqlx.DB
}

func NewActivityRepository(db *sqlx.DB) *ActivityPostgres {
	return &ActivityPostgres{db: db}
}

func (r *ActivityPostgres) GetAllActivities(userId int) ([]models.ActivitiesOut, error) {
	activities := make([]models.ActivitiesOut, 0)
	query := fmt.Sprintf("SELECT change, activity_type, label, activity_date "+
		"FROM %s "+
		"WHERE user_id = $1", activitiesTable)
	logger.Debug(query)
	err := r.db.Select(&activities, query, userId)
	return activities, err
}

func (r *ActivityPostgres) CreateActivity(activity models.Activity) error {
	query := fmt.Sprintf("INSERT INTO %s(user_id, activity_type, change, label, activity_date) "+
		"VALUES ($1, $2, $3, $4, $5)", activitiesTable)
	logger.Debug(query)
	_, err := r.db.Exec(query, activity.UserId, activity.Type,
		activity.Change, activity.Label, activity.ActivityDate)
	return err
}
