//go:generate enumer -type=State -trimprefix=State -json -yaml -text -sql

package domain

import "gorm.io/gorm"

type State int

const (
	Invalid State = iota
	Doing
	Done
)

type QuizReport struct {
	UserID        int64
	TopicID       int64
	QuizID        int64
	Percentage    string
	Duration      int64
	TotalDuration int64
	Result        bool
	StartTime     int64
	TotalQuiz     int64
	MaskObtained  int64
	State         State
	Answers       string
	gorm.Model
}
