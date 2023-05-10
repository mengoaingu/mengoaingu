//go:generate enumer -type=Level -trimprefix=Level -json -yaml -text -sql
//go:generate enumer -type=QuestionType -trimprefix=QuestionType -json -yaml -text -sql

package domain

import "gorm.io/gorm"

type Level int

const (
	InvalidLevel Level = iota
	Level_A1
	Level_A2
	Level_B1
	Level_B2
	Level_C1
	Level_C2
)

type QuestionType int

const (
	InvalidType QuestionType = iota
	QuestionType_PART1
	QuestionType_PART2
	QuestionType_PART3
	QuestionType_PART4
	QuestionType_PART5
	QuestionType_PART6
	QuestionType_PART7
)

type Question struct {
	QuestionText string
	Answers      string
	Explain      string
	Level        Level
	QuestionType QuestionType
	Image        string
	Audio        string
	gorm.Model
}
