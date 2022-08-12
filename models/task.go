package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model

	ID          uint `gorm:"primaryKey;autoIncrement:true"`
	Assignee    string
	Description string
	Date        string
	IsDone      bool
}
