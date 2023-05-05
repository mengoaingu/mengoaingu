package grpc

import (
	"backend/internal/profile/domain"
	"backend/internal/profile/usecase"
	"backend/pkg/middleware"
	pb "backend/proto/gen"
	context "context"
	"errors"
	"strings"

	"google.golang.org/grpc"
)

type ProfileGRPCServer struct {
	profileUC  usecase.ProfileUsecase
	authClient *middleware.Authorization
}

func NewProfileGRPCServer(grpcServer *grpc.Server, profileUC usecase.ProfileUsecase, authClient *middleware.Authorization) *ProfileGRPCServer {
	profileGrpcServer := &ProfileGRPCServer{
		profileUC:  profileUC,
		authClient: authClient,
	}

	pb.RegisterProfileServiceServer(grpcServer, profileGrpcServer)

	return profileGrpcServer
}

var _ pb.ProfileServiceServer = (*ProfileGRPCServer)(nil)

// Authorize implements gen.ProfileServiceServer
func (h *ProfileGRPCServer) Authorize(ctx context.Context, request *pb.AuthorizeRequest) (*pb.AuthorizeResponse, error) {
	token, err := h.authClient.AuthorizeN(ctx)
	if err != nil {
		return nil, errors.New("you don't have permission")
	}

	resp, err := h.profileUC.GetProfile(ctx, &usecase.GetProfileRequest{
		UserID: token.UID,
	})
	if err != nil {
		return nil, errors.New("something wrong")
	}

	if resp.Profile.UserId == "" {
		resp, err := h.profileUC.CreateProfile(ctx, &usecase.CreateProfileRequest{
			Profile: &domain.Profile{
				UserId:      token.UID,
				DisplayName: strings.Split(token.Claims["email"].(string), "@")[0],
				Email:       token.Claims["email"].(string),
			},
		})
		if err != nil {
			return nil, errors.New("something wrong")
		}
		return &pb.AuthorizeResponse{
			Data: &pb.AuthorizeResponse_Data{
				Profile: &pb.Profile{
					UserId:      resp.Profile.UserId,
					DisplayName: resp.Profile.DisplayName,
					Email:       resp.Profile.Email,
				},
			},
		}, nil
	}

	return &pb.AuthorizeResponse{
		Data: &pb.AuthorizeResponse_Data{
			Profile: &pb.Profile{
				UserId:      resp.Profile.UserId,
				DisplayName: resp.Profile.DisplayName,
				Email:       resp.Profile.Email,
			},
		},
	}, nil
}
