package domain

import "gorm.io/gorm"

type Topic struct {
	Title     string
	Thumbnail string
	gorm.Model
}
