package grpc

import (
	"backend/internal/quizzes/domain"
	"backend/proto/gen"

	"google.golang.org/protobuf/types/known/wrapperspb"
)

func topic2Pb(topic *domain.Topic) *gen.Topic {
	return &gen.Topic{
		Id:        int64(topic.ID),
		Title:     topic.Title,
		Thumbnail: topic.Thumbnail,
	}
}

func topics2Pb(topics []*domain.Topic) []*gen.Topic {
	var result []*gen.Topic
	for _, topic := range topics {
		result = append(result, topic2Pb(topic))
	}
	return result
}

func quiz2Pb(quiz *domain.Quiz) *gen.Quiz {
	return &gen.Quiz{
		Id:          int64(quiz.ID),
		TopicId:     int64(quiz.TopicID),
		Name:        quiz.Name,
		Description: quiz.Description,
		IsFree:      &wrapperspb.BoolValue{Value: quiz.IsFree},
	}
}

func quizzes2Pb(quizzes []*domain.Quiz) []*gen.Quiz {
	var result []*gen.Quiz
	for _, quiz := range quizzes {
		result = append(result, quiz2Pb(quiz))
	}
	return result
}
