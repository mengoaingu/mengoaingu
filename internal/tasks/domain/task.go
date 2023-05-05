//go:generate enumer -type=TaskStatus -trimprefix=TaskStatus -json -yaml -text -sql
//go:generate enumer -type=Priority -trimprefix=Priority -json -yaml -text -sql

package domain

import "gorm.io/gorm"

type TaskStatus int

const (
	TaskStatusInvalid TaskStatus = iota
	TaskStatusInProgress
	TaskStatusCompleted
	TaskStatusCancelled
)

type Priority int

const (
	PriorityInvalid Priority = iota
	PriorityLow
	PriorityMedium
	PriorityHigh
)

type Task struct {
	Name        string
	Description string
	Status      TaskStatus
	Priority    Priority
	UserID      string
	gorm.Model
}
