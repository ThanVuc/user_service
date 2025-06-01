//go:build wireinject

package wire

import (
	"authentication_service/internal/controller"
	"authentication_service/internal/repos"
	"authentication_service/internal/services"

	"github.com/google/wire"
)

func InjectAuthController() *controller.AuthController {
	wire.Build(
		repos.NewAuthRepo,
		services.NewAuthService,
		controller.CreateAuthController,
	)
	return new(controller.AuthController)
}
