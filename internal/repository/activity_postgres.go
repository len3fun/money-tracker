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
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	newActivityQuery := fmt.Sprintf("INSERT INTO %s(user_id, source_id, activity_type, change, label, activity_date) "+
		"VALUES ($1, $2, $3, $4, $5, $6)", activitiesTable)
	logger.Debug(newActivityQuery)
	_, err = r.db.Exec(newActivityQuery, activity.UserId, activity.SourceId, activity.Type,
		activity.Change, activity.Label, activity.ActivityDate)
	if err != nil {
		tx.Rollback()
		return err
	}

	updateBalanceQuery := r.generateUpdateBalanceQuery(activity.Type)

	_, err = r.db.Exec(updateBalanceQuery, activity.Change, activity.SourceId, activity.UserId)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

// generateUpdateBalanceQuery generates db request for updating balance depending on activity type.
func (r *ActivityPostgres) generateUpdateBalanceQuery(activityType string) string {
	var sign string
	if activityType == "income" {
		sign = "+"
	} else if activityType == "expense" {
		sign = "-"
	}
	return fmt.Sprintf("UPDATE %s SET balance = balance %s $1 WHERE id = $2 AND user_id = $3;",
		sourcesTable, sign)
}
