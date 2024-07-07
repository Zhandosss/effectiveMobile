package entity

import (
	"effectiveMobileTestProblem/internal/model"
	"time"
)

// TODO: Delete json???
type UserDB struct {
	ID             string `db:"id"`
	PassportSeries string `db:"passport_series"`
	PassportNumber string `db:"passport_number"`
	Name           string `db:"name"`
	Surname        string `db:"surname"`
	Address        string `db:"address"`
}

func (u *UserDB) ToUser() *model.User {
	return &model.User{
		ID:                      u.ID,
		PassportSeriesAndNumber: u.PassportSeries + " " + u.PassportNumber,
		Name:                    u.Name,
		Surname:                 u.Surname,
		Address:                 u.Address,
	}
}

type WorkDB struct {
	ID        string     `db:"id"`
	Name      string     `db:"name"`
	UserID    string     `db:"user_id"`
	StartTime time.Time  `db:"start_time"`
	EndTime   *time.Time `db:"end_time"`
}

func (w *WorkDB) ToWork() *model.Work {
	var workTime time.Duration
	var isWorking bool
	if w.EndTime == nil {
		workTime = time.Now().Sub(w.StartTime)
		isWorking = true
	} else {
		workTime = w.EndTime.Sub(w.StartTime)
		isWorking = false
	}
	return &model.Work{
		ID:        w.ID,
		Name:      w.Name,
		StartTime: w.StartTime,
		EndTime:   w.EndTime,
		UserID:    w.UserID,
		WorkTime:  workTime.String(),
		IsWorking: isWorking,
	}
}
