package models

type TeamDetail struct {
	ID           uint `gorm:"primaryKey;autoIncrement"`
	TeamName     string
	ProjectName  string
	IsFinished   bool
	StartTime    string
	FinishedTime string
}
