//go:build wireinject
// +build wireinject

package routers

import (
	"github.com/Hsun-Weng/human-resource-service/internal/controllers"
	"github.com/Hsun-Weng/human-resource-service/internal/middleware"
	"github.com/Hsun-Weng/human-resource-service/internal/repository"
	"github.com/Hsun-Weng/human-resource-service/internal/services"
	"github.com/Hsun-Weng/human-resource-service/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var Set = wire.NewSet(
	util.NewDb,
	util.NewRedisClient,
	repository.ProviderSet,
	services.ProviderSet,
	controllers.ProviderSet,
	middleware.NewAdminAuthenticationMiddleware,
	NewRouter,
)

func InitializeEngine() (*gin.Engine, error) {
	wire.Build(
		Set,
	)
	return nil, nil
}
