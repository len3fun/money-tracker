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
	query := fmt.Sprintf("SELECT id, change, activity_type, label, activity_date "+
		"FROM %s "+
		"WHERE user_id = $1 "+
		"ORDER BY activity_date DESC", activitiesTable)
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
		"SELECT $1, s.id, $3, $4, $5, $6 "+
		"FROM %s s "+
		"WHERE user_id = $1 AND id = $2", activitiesTable, sourcesTable)
	logger.Debug(newActivityQuery)
	res, err := r.db.Exec(newActivityQuery, activity.UserID, activity.SourceID, activity.Type,
		activity.Change, activity.Label, activity.ActivityDate)
	if err != nil {
		tx.Rollback()
		return err
	}

	n, err := res.RowsAffected()
	if err != nil {
		tx.Rollback()
		return err
	}

	if n == 0 {
		tx.Rollback()
		return fmt.Errorf("user doesn't have such source id")
	}

	updateBalanceQuery := r.generateUpdateBalanceQuery(activity.Type)

	_, err = r.db.Exec(updateBalanceQuery, activity.Change, activity.SourceID, activity.UserID)
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
