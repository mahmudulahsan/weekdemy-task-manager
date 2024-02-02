package models

type BookDetail struct {
	ID          uint `gorm:"primaryKey;autoIncrement"`
	BookName    string
	AuthorID    uint
	Publication string
}
