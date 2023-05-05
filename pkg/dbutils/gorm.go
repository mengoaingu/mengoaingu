package dbutils

import (
	"context"
	"database/sql"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/opentelemetry/tracing"
)

type GormOption func(db *gorm.DB)

func WithLogger(l logger.Interface) GormOption {
	return func(db *gorm.DB) {
		db.Config.Logger = l
	}
}

func WithPlugIn(plugin gorm.Plugin) GormOption {
	return func(db *gorm.DB) {
		_ = db.Use(plugin)
	}
}

func WithTracing() GormOption {
	p := tracing.NewPlugin()
	return WithPlugIn(p)
}

func WithDefaultLogger() GormOption {
	l := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             200 * time.Millisecond, // Slow SQL threshold
			LogLevel:                  logger.Info,            // Log level
			IgnoreRecordNotFoundError: true,                   // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,                   // Disable color
		},
	)
	return WithLogger(l)
}

func NewMySQLGorm(db *sql.DB, opts ...GormOption) (*gorm.DB, error) {
	mysqld := mysql.New(mysql.Config{
		Conn: db,
	})
	gormDB, err := gorm.Open(mysqld, &gorm.Config{
		Logger: logger.Discard,
	})
	if err != nil {
		return nil, err
	}
	opts = append(opts, WithTracing())
	for _, o := range opts {
		o(gormDB)
	}
	return gormDB, nil
}

func NewGormConnPool(db *sql.DB) (gorm.ConnPool, error) {
	gormDb, err := NewMySQLGorm(db)
	return &GormRawLogger{DB: gormDb}, err
}

type GormRawLogger struct {
	*gorm.DB
}

func (g *GormRawLogger) ExecContext(ctx context.Context, s string, i ...interface{}) (sql.Result, error) {
	query := g.DB.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Exec(s, i...)
	})
	g.DB.Logger.Info(ctx, query)
	return g.DB.ConnPool.ExecContext(ctx, s, i)
}

func (g *GormRawLogger) PrepareContext(ctx context.Context, s string) (*sql.Stmt, error) {
	query := g.DB.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Exec(s)
	})
	g.DB.Logger.Info(ctx, query)
	return g.DB.ConnPool.PrepareContext(ctx, s)
}

func (g *GormRawLogger) QueryContext(ctx context.Context, s string, i ...interface{}) (*sql.Rows, error) {
	query := g.DB.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Exec(s, i...)
	})
	logger.Default.Info(ctx, query)
	return g.DB.ConnPool.QueryContext(ctx, s, i...)
}

func (g *GormRawLogger) QueryRowContext(ctx context.Context, s string, i ...interface{}) *sql.Row {
	query := g.DB.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Exec(s, i...)
	})
	logger.Default.Info(ctx, query)
	return g.DB.WithContext(ctx).Exec(s, i...).Row()
}
