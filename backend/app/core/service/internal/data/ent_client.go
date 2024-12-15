package data

import (
	"context"

	"go-cms/app/core/service/internal/data/ent"
	"go-cms/app/core/service/internal/data/ent/migrate"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/tx7do/go-utils/entgo"
	conf "github.com/tx7do/kratos-bootstrap/api/gen/go/conf/v1"
)

// NewEntClient 创建Ent ORM数据库客户端
func NewEntClient(cfg *conf.Bootstrap, logger log.Logger) *entgo.EntClient[*ent.Client] {
	l := log.NewHelper(log.With(logger, "module", "ent/data/core-service"))

	drv, err := entgo.CreateDriver(
		cfg.Data.Database.GetDriver(),
		cfg.Data.Database.GetSource(),
		cfg.Data.Database.GetEnableTrace(),
		cfg.Data.Database.GetEnableMetrics(),
	)
	if err != nil {
		l.Fatalf("failed opening connection to db: %v", err)
		return nil
	}

	client := ent.NewClient(
		ent.Driver(drv),
		ent.Log(func(a ...any) {
			l.Debug(a...)
		}),
	)

	// 运行数据库迁移工具
	if cfg.Data.Database.GetMigrate() {
		if err = client.Schema.Create(context.Background(), migrate.WithForeignKeys(true)); err != nil {
			l.Fatalf("failed creating schema resources: %v", err)
		}
	}

	cli := entgo.NewEntClient(client, drv)

	cli.SetConnectionOption(
		int(cfg.Data.Database.GetMaxIdleConnections()),
		int(cfg.Data.Database.GetMaxOpenConnections()),
		cfg.Data.Database.GetConnectionMaxLifetime().AsDuration(),
	)

	return cli
}
