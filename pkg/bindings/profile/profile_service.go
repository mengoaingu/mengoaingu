package profile

import (
	grpcHandler "backend/internal/profile/infrastructure/ports/grpc"
	"backend/internal/profile/repository"
	"backend/internal/profile/usecase"
	"backend/pkg/config"
	"backend/pkg/dbutils"
	"database/sql"
	"log"

	DB "backend/internal/profile/infrastructure/db"

	"go.uber.org/fx"
	gormoteltracing "gorm.io/plugin/opentelemetry/tracing"
)

var (
	moduleName = "PROFILE"

	Module = fx.Options(
		fx.Provide(CreateMySQLDB),
		fx.Invoke(RunMigration),
		fx.Invoke(grpcHandler.NewProfileGRPCServer),
		fx.Provide(usecase.NewProfileUsecase),
		fx.Provide(repository.NewProfileRepository),
	)
)

func CreateMySQLDB(cfg *config.Config) repository.ProfileDB {

	db, err := sql.Open("mysql", cfg.PROFILE_MYSQL_DSN)
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

func RunMigration(cfg *config.Config) error {
	db, err := sql.Open("mysql", cfg.PROFILE_MYSQL_DSN)
	if err != nil {
		panic(err)
	}
	err = DB.Migrate(db, -1)
	if err == nil || err.Error() == "no change" {
		return nil
	}
	return nil
}
