package tasks

import (
	grpcHandler "backend/internal/tasks/infrastructure/ports/grpc"
	"backend/internal/tasks/repository"
	"backend/internal/tasks/usecase"
	"backend/pkg/dbutils"
	"backend/proto/gen"
	"context"
	"database/sql"
	"errors"
	"log"

	gwruntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	gormoteltracing "gorm.io/plugin/opentelemetry/tracing"
)

var (
	moduleName = "tasks"

	Module = fx.Options(
		fx.Provide(CreateMySQLDB),
		fx.Invoke(grpcHandler.NewTaskGRPCServer),
		fx.Provide(usecase.NewTaskUsecase),
		fx.Provide(repository.NewTaskRepository),
		// fx.Provide(GRPCHandler),
	)
)

func GRPCHandler(ctx context.Context, mux *gwruntime.ServeMux) *gwruntime.ServeMux {
	// gen.RegisterProfileServiceHandlerFromEndpoint(ctx, mux, "localhost:9090", []grpc.DialOption{grpc.WithInsecure()})
	gen.RegisterTaskServiceHandlerFromEndpoint(ctx, mux, "localhost:9090", []grpc.DialOption{grpc.WithInsecure()})
	return mux
}

func CreateMySQLDB() repository.TasksDB {
	var dsn string
	v := viper.GetString(moduleName + ".mysql.dsn")
	if v != "" {
		dsn = v
	} else {
		panic(errors.New(moduleName + ".mysql.dsn is not set"))
	}
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	gormDB, err := dbutils.NewMySQLGorm(
		db,
		dbutils.WithDefaultLogger(),
		dbutils.WithPlugIn(gormoteltracing.NewPlugin()),
	)
	if err != nil {
		log.Fatal("can not init gormDB ", err)
	}
	return gormDB
}
