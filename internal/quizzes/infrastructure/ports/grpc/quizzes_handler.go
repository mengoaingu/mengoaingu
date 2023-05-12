package grpc

import (
	"backend/internal/quizzes/usecase"
	"backend/pkg"
	"backend/pkg/middleware"
	"backend/proto/gen"
	"context"
)

type QuizzesGRPCServer struct {
	quizzesUC  usecase.QuizzesUsecase
	authClient *middleware.Authorization
}

func NewQuizzesGRPCServer(grpcServer *pkg.App, taskUC usecase.QuizzesUsecase, authClient *middleware.Authorization) *QuizzesGRPCServer {
	quizzesGRPCServer := &QuizzesGRPCServer{
		authClient: authClient,
	}

	gen.RegisterQuizzesServiceServer(grpcServer.GrpcServer, quizzesGRPCServer)

	return quizzesGRPCServer
}

var _ gen.QuizzesServiceServer = (*QuizzesGRPCServer)(nil)

// GetListQuizzes implements gen.QuizzesServiceServer
func (h *QuizzesGRPCServer) GetListQuizzes(ctx context.Context, request *gen.GetListQuizzesRequest) (*gen.GetListQuizzesResponse, error) {
	topicID := request.TopicId
	resp, err := h.quizzesUC.FindAllQuiz(context.Background(), topicID)
	if err != nil {
		return nil, err
	}
	return &gen.GetListQuizzesResponse{
		Data: &gen.GetListQuizzesResponse_Data{
			Quizzes: quizzes2Pb(resp),
		},
	}, nil
}

// GetListTopics implements gen.QuizzesServiceServer
func (h *QuizzesGRPCServer) GetListTopics(context.Context, *gen.GetListTopicsRequest) (*gen.GetListTopicsResponse, error) {
	resp, err := h.quizzesUC.FindAllTopic(context.Background())
	if err != nil {
		return nil, err
	}
	return &gen.GetListTopicsResponse{
		Data: &gen.GetListTopicsResponse_Data{
			Topics: topics2Pb(resp),
		},
	}, nil
}

// StartQuiz implements gen.QuizzesServiceServer
func (h *QuizzesGRPCServer) StartQuiz(context.Context, *gen.StartQuizRequest) (*gen.StartQuizResponse, error) {
	panic("unimplemented")
}

// SubmitQuiz implements gen.QuizzesServiceServer
func (h *QuizzesGRPCServer) SubmitQuiz(context.Context, *gen.SubmitQuizRequest) (*gen.SubmitQuizResponse, error) {
	panic("unimplemented")
}
