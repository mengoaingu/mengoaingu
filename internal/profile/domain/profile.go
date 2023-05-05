package domain

import "gorm.io/gorm"

type Profile struct {
	ID          int64
	UserId      string
	DisplayName string
	Email       string
	gorm.Model
}
