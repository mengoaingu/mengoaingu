package profile

import (
	grpcHandler "backend/internal/profile/infrastructure/ports/grpc"
	"backend/internal/profile/repository"
	"backend/internal/profile/usecase"
	"backend/pkg/dbutils"
	"database/sql"
	"errors"
	"log"

	"github.com/spf13/viper"
	"go.uber.org/fx"
	gormoteltracing "gorm.io/plugin/opentelemetry/tracing"
)

var (
	moduleName = "profile"

	Module = fx.Options(
		fx.Provide(CreateMySQLDB),
		fx.Invoke(grpcHandler.NewProfileGRPCServer),
		fx.Provide(usecase.NewProfileUsecase),
		fx.Provide(repository.NewProfileRepository),
	)
)

func CreateMySQLDB() repository.ProfileDB {
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
