package repository

import (
	"context"
	"effectiveMobileTestProblem/internal/entity"
	"effectiveMobileTestProblem/internal/model"
	"errors"
	"fmt"
	"slices"
	"time"
)

func (r *WorkRepository) NewWork(ctx context.Context, work *model.Work) (string, error) {
	tx, err := r.conn.Begin()
	if err != nil {
		return "", err
	}
	defer tx.Rollback()

	query := `INSERT INTO works (start_time, name, user_id) VALUES ($1, $2, $3) RETURNING id`
	var id string

	err = r.conn.GetContext(ctx, &id, query, work.StartTime, work.Name, work.UserID)
	if err != nil {
		return "", err
	}

	if tx.Commit() != nil {
		return "", err
	}

	return id, nil
}

func (r *WorkRepository) StopWork(ctx context.Context, id string) error {
	tx, err := r.conn.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	endTime := ctx.Value("end_time").(time.Time)
	if endTime.IsZero() {
		return errors.New("end time is not set")
	}

	query := `SELECT start_time, time_period_in_minute FROM works WHERE id = $1`
	var startTime time.Time
	var timePeriodString string
	err = r.conn.QueryRowContext(ctx, query, id).Scan(&startTime, &timePeriodString)
	if err != nil {
		return err
	}

	workDuration := endTime.Sub(startTime)

	oldTime, err := time.ParseDuration(timePeriodString)
	if err != nil {
		return err
	}

	newTime := oldTime + workDuration

	hours := int(newTime.Hours())
	minute := int(newTime.Minutes()) % 60

	timeAns := fmt.Sprintf("%dh%dm", hours, minute)

	query = `UPDATE works SET end_time = $1, time_period_in_minute = $2 WHERE id = $3`
	_, err = r.conn.ExecContext(ctx, query, endTime, timeAns, id)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *WorkRepository) GetWorkById(ctx context.Context, id string) (*entity.WorkDB, error) {
	query := `SELECT id, start_time, end_time, name, user_id FROM works WHERE id=$1`
	work := make([]*entity.WorkDB, 0, 1)
	err := r.conn.SelectContext(ctx, &work, query, id)
	if err != nil {
		return &entity.WorkDB{}, err
	}
	return work[0], nil
}

func (r *WorkRepository) GetWorks(ctx context.Context, user string) ([]*entity.WorkDB, error) {
	query := `SELECT id, start_time, end_time, name, user_id FROM works WHERE user_id=$1`
	works := make([]*entity.WorkDB, 0)
	err := r.conn.SelectContext(ctx, &works, query, user)
	endTime := ctx.Value("end_time").(time.Time)
	fmt.Println(1)
	slices.SortFunc(works, func(i, j *entity.WorkDB) int {
		iEndTime := i.EndTime
		jEndTime := j.EndTime
		if i.EndTime == nil {
			iEndTime = &endTime
		}
		if j.EndTime == nil {
			jEndTime = &endTime
		}
		if iEndTime.Before(*jEndTime) {
			return 1
		} else if iEndTime.After(*jEndTime) {
			return -1
		} else {
			return 0
		}
	})
	fmt.Println(err)
	if err != nil {
		return nil, err
	}
	return works, nil
}
