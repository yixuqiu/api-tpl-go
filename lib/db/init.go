package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"api/ent"
	"api/lib/config"
	"api/lib/log"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

var cli *ent.Client

// Config 数据库初始化配置
type Config struct {
	// Driver 驱动名称
	Driver string

	// DSN 数据源名称
	// [-- MySQL] username:password@tcp(localhost:3306)/dbname?timeout=10s&charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=True&loc=Local
	// [Postgres] host=localhost port=5432 user=root password=secret dbname=test connect_timeout=10 sslmode=disable
	// [- SQLite] file::memory:?cache=shared
	DSN string `json:"dsn"`

	// Options 配置选项
	Options *Options `json:"options"`
}

// Options 数据库配置选项
type Options struct {
	// MaxOpenConns 设置最大可打开的连接数
	MaxOpenConns int `json:"max_open_conns"`

	// MaxIdleConns 连接池最大闲置连接数
	MaxIdleConns int `json:"max_idle_conns"`

	// ConnMaxLifetime 连接的最大生命时长
	ConnMaxLifetime time.Duration `json:"conn_max_lifetime"`

	// ConnMaxIdleTime 连接最大闲置时间
	ConnMaxIdleTime time.Duration `json:"conn_max_idle_time"`
}

// Init 初始化Ent实例
func Init(cfg *Config) error {
	db, err := sql.Open(cfg.Driver, cfg.DSN)
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		db.Close()
		return err
	}

	if cfg.Options != nil {
		db.SetMaxOpenConns(cfg.Options.MaxOpenConns)
		db.SetMaxIdleConns(cfg.Options.MaxIdleConns)
		db.SetConnMaxLifetime(cfg.Options.ConnMaxLifetime)
		db.SetConnMaxIdleTime(cfg.Options.ConnMaxIdleTime)
	}

	cli = ent.NewClient(
		ent.Driver(dialect.DebugWithContext(
			entsql.OpenDB(cfg.Driver, db),
			func(ctx context.Context, v ...any) {
				if config.ENV.Debug {
					log.Info(ctx, "SQL info", zap.String("SQL", fmt.Sprint(v...)))
				}
			}),
		),
	)

	return nil
}

// Client 返回Ent实例
func Client() *ent.Client {
	return cli
}
