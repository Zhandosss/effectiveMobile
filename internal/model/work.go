package model

import "time"

type Work struct {
	ID        string     `json:"id,omitempty"`
	Name      string     `json:"name"`
	StartTime time.Time  `json:"start_time"`
	EndTime   *time.Time `json:"end_time"`
	WorkTime  string     `json:"work_time"`
	IsWorking bool       `json:"is_working"`
	UserID    string     `json:"user_id"`
}
