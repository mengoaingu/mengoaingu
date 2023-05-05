package bindings

import (
	"net/http"

	gwruntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type App struct {
	Mux        *gwruntime.ServeMux
	GrpcServer *grpc.Server
}

func New(s *http.ServeMux) *App {
	mux := gwruntime.NewServeMux()
	grpc := grpc.NewServer()

	return &App{
		Mux:        mux,
		GrpcServer: grpc,
	}
}
