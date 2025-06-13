//go:build wireinject

package wire

import (
	"user_service/internal/controller"
	"user_service/internal/repos"
	"user_service/internal/services"

	"github.com/google/wire"
)

func InjectUserController() *controller.UserController {
	wire.Build(
		repos.NewUserRepo,
		services.NewUserService,
		controller.NewUserController,
	)
	return new(controller.UserController)
}
