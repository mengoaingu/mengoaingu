package usecase

import (
	"backend/internal/quizzes/domain"
	"context"
)

type QuizzesUsecase interface {
	FindAllTopic(ctx context.Context) ([]*domain.Topic, error)
	FindAllQuiz(ctx context.Context, topic_id int64) ([]*domain.Quiz, error)
	FindAllQuestionByTopicID(ctx context.Context, quiz_id int64) ([]*domain.Question, error)
	CreateQuizReport(ctx context.Context, quizReport *domain.QuizReport) (*domain.QuizReport, error)
}
