package profile

// import (
// 	"backend/internal/profile/infrastructure/ports/grpc"
// 	"backend/internal/profile/repository"
// 	"backend/internal/profile/usecase"
// 	"backend/pkg/dbutils"
// 	"database/sql"
// 	"errors"
// 	"log"

// 	_ "github.com/go-sql-driver/mysql"
// 	"github.com/google/wire"
// 	"github.com/spf13/viper"
// 	gormoteltracing "gorm.io/plugin/opentelemetry/tracing"
// )

// var (
// 	moduleName = "profile"
// )

// func CreateMySQLDB() repository.ProfileDB {
// 	var dsn string
// 	v := viper.GetString(moduleName + ".mysql.dsn")
// 	if v != "" {
// 		dsn = v
// 	} else {
// 		panic(errors.New(moduleName + ".mysql.dsn is not set"))
// 	}
// 	db, err := sql.Open("mysql", dsn)
// 	if err != nil {
// 		panic(err)
// 	}
// 	gormDB, err := dbutils.NewMySQLGorm(
// 		db,
// 		dbutils.WithDefaultLogger(),
// 		dbutils.WithPlugIn(gormoteltracing.NewPlugin()),
// 	)
// 	if err != nil {
// 		log.Fatalf("can not init gormDB ", err)
// 	}
// 	return gormDB
// }

// var ProfileSet = wire.NewSet(
// 	grpc.NewProfileGRPCServer,
// 	usecase.NewProfileUsecase,
// 	repository.NewProfileRepository,
// 	CreateMySQLDB,
// )
