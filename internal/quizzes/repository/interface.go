package repository

import (
	"backend/internal/quizzes/domain"
	"context"
)

type TopicRepository interface {
	Save(ctx context.Context, topic *domain.Topic) error
	Find(ctx context.Context) ([]*domain.Topic, int64, error)
}

type QuizRepository interface {
	Save(ctx context.Context, quiz *domain.Quiz) error
	FindByTopicID(ctx context.Context, topic_id int64) ([]*domain.Quiz, error)
}

type QuestionRepository interface {
	Save(ctx context.Context, question *domain.Question) error
	FindByQuizID(ctx context.Context, quiz_id int64) ([]*domain.Question, error)
}

type QuizReportRepository interface {
	Save(ctx context.Context, quizReport *domain.QuizReport) error
	Update(ctx context.Context, id int64, quizReport *domain.QuizReport) error
	FindByID(ctx context.Context, id int64) (*domain.QuizReport, error)
}
