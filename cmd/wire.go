//go:build wireinject
// +build wireinject

package main

import (
	"app/application/controllers"
	bzs "app/core/bizs"
	config "app/core/configs"
	ops "app/infrastructure/opensearch"
	mr "app/infrastructure/repositories/mysql"
	// ks "app/pkg/kafkasrv"
	log "app/pkg/logger"
	ms "app/pkg/mysql"
	opspk "app/pkg/opensearch"
	"github.com/google/wire"
)

func NewApp(cfg *config.Config, logger *log.ApiLogger) (controllers.AppServer, error) {
	panic(wire.Build(
		// ks.Wired,
		ms.Wired,
		ops.Wired,
		opspk.Wired,

		// Job Repo.
		mr.Wired,

		// Job Biz.
		bzs.Wired,

		// Http server.
		controllers.Wired,
		// wire.Bind(new(controllers.AppServer), new(*controllers.RestfulServer)),
	))
}
