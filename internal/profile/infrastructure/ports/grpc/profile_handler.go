package grpc

import (
	"backend/internal/profile/domain"
	"backend/internal/profile/usecase"
	"backend/pkg"
	"backend/pkg/middleware"
	"backend/pkg/token"
	pb "backend/proto/gen"
	context "context"
	"errors"
	"fmt"
	"time"
)

type ProfileGRPCServer struct {
	profileUC  usecase.ProfileUsecase
	authClient *middleware.Authorization
	jwtMaker   token.Maker
}

// GetProfile implements gen.ProfileServiceServer
func (h *ProfileGRPCServer) GetProfile(ctx context.Context, request *pb.GetProfileRequest) (*pb.GetProfileResponse, error) {
	tokenJwt, err := h.authClient.AuthorizeJWT(ctx)
	if err != nil {
		return nil, fmt.Errorf("invalid access token: %s", err)
	}

	resp, err := h.profileUC.GetProfile(ctx, &usecase.GetProfileRequest{
		UserID: tokenJwt.ID,
	})
	if err != nil {
		return nil, errors.New("something wrong")
	}

	return &pb.GetProfileResponse{
		Data: &pb.GetProfileResponse_Data{
			Profile: &pb.User{
				Id:          int64(resp.Profile.ID),
				Provider:    resp.Profile.Provider,
				Uid:         resp.Profile.UID,
				DisplayName: resp.Profile.DisplayName,
				Email:       resp.Profile.Email,
				Avatar:      resp.Profile.Avatar,
			},
		},
	}, nil

}

// Authorize implements gen.ProfileServiceServer
func (h *ProfileGRPCServer) Authorize(ctx context.Context, request *pb.AuthorizeRequest) (*pb.AuthorizeResponse, error) {
	token := request.GetIdToken()
	firebaseToken, err := h.authClient.VerifyToken(ctx, token)
	if err != nil {
		return nil, fmt.Errorf("invalid access token: %s", err)
	}

	resp, err := h.profileUC.GetProfile(ctx, &usecase.GetProfileRequest{
		UserID: firebaseToken.UID,
	})
	if err != nil {
		return nil, errors.New("something wrong")
	}

	if resp.Profile.UID == "" {
		resp, err := h.profileUC.CreateProfile(ctx, &usecase.CreateProfileRequest{
			Profile: &domain.User{
				UID:         firebaseToken.UID,
				DisplayName: firebaseToken.Claims["name"].(string),
				Email:       firebaseToken.Claims["email"].(string),
				Avatar:      firebaseToken.Claims["picture"].(string),
				Provider:    firebaseToken.Firebase.SignInProvider,
			},
		})
		if err != nil {
			return nil, errors.New("something wrong")
		}
		accessToken, _, err := h.jwtMaker.CreateToken(resp.Profile.UID, resp.Profile.DisplayName, resp.Profile.Email, time.Duration(30*time.Minute))
		if err != nil {
			return nil, errors.New("something wrong")
		}
		return &pb.AuthorizeResponse{
			Data: &pb.AuthorizeResponse_Data{
				Token: accessToken,
			},
		}, nil
	}

	accessToken, _, err := h.jwtMaker.CreateToken(resp.Profile.UID, resp.Profile.DisplayName, resp.Profile.Email, time.Duration(30*time.Minute))
	if err != nil {
		return nil, errors.New("something wrong")
	}
	return &pb.AuthorizeResponse{
		Data: &pb.AuthorizeResponse_Data{
			Token: accessToken,
		},
	}, nil
}

func NewProfileGRPCServer(app *pkg.App, profileUC usecase.ProfileUsecase, authClient *middleware.Authorization, tokenMaker token.Maker) *ProfileGRPCServer {
	profileGrpcServer := &ProfileGRPCServer{
		profileUC:  profileUC,
		authClient: authClient,
		jwtMaker:   tokenMaker,
	}

	pb.RegisterProfileServiceServer(app.GrpcServer, profileGrpcServer)

	return profileGrpcServer
}

var _ pb.ProfileServiceServer = (*ProfileGRPCServer)(nil)
