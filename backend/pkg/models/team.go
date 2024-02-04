package models

import "time"

type TeamDetail struct {
	ID           uint `gorm:"primaryKey;autoIncrement"`
	TeamName     string
	ProjectName  string
	IsFinished   bool
	StartTime    time.Time
	FinishedTime *time.Time
}
