package domain

import "gorm.io/gorm"

type Provider int

type User struct {
	UID         string
	DisplayName string
	Email       string
	Avatar      string
	Provider    string
	gorm.Model
}
