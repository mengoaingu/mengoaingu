package tasks

import (
	grpcHandler "backend/internal/tasks/infrastructure/ports/grpc"
	"backend/internal/tasks/repository"
	"backend/internal/tasks/usecase"
	"backend/pkg/dbutils"
	"database/sql"
	"errors"
	"log"

	"github.com/spf13/viper"
	"go.uber.org/fx"
	gormoteltracing "gorm.io/plugin/opentelemetry/tracing"
)

var (
	moduleName = "TASKS"

	Module = fx.Options(
		fx.Provide(CreateMySQLDB),
		fx.Invoke(grpcHandler.NewTaskGRPCServer),
		fx.Provide(usecase.NewTaskUsecase),
		fx.Provide(repository.NewTaskRepository),
	)
)

func CreateMySQLDB() repository.TasksDB {
	var dsn string
	v := viper.GetString(moduleName + "_MYSQL_DSN")
	if v != "" {
		dsn = v
	} else {
		panic(errors.New(moduleName + "_MYSQL_DSN is not set"))
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
