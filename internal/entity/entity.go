package entity

import (
	"effectiveMobileTestProblem/internal/model"
	"time"
)

// TODO: Delete json???
type UserDB struct {
	ID             string `json:"id,omitempty" db:"id"`
	PassportSeries string `json:"passport_series" db:"passport_series"`
	PassportNumber string `json:"passport_number" db:"passport_number"`
	Name           string `json:"name" db:"name"`
	Surname        string `json:"surname" db:"surname"`
	Address        string `json:"address" db:"address"`
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
	ID        string     `json:"id,omitempty" db:"id"`
	Name      string     `json:"name" db:"name"`
	UserID    string     `json:"user_id" db:"user_id"`
	StartTime time.Time  `json:"start_time" db:"start_time"`
	EndTime   *time.Time `json:"end_time" db:"end_time"`
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
