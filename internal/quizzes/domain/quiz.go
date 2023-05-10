package domain

import "gorm.io/gorm"

type Quiz struct {
	Name        string
	Description string
	TopicID     int64
	IsFree      bool
	gorm.Model
}
