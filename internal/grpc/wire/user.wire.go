//go:build wireinject

package wire

import (
	"user_service/internal/grpc/controller"
	"user_service/internal/grpc/mapper"
	"user_service/internal/grpc/repos"
	"user_service/internal/grpc/services"

	"github.com/google/wire"
)



func InjectUserController() *controller.UserController {
		wire.Build(
			repos.NewUserRepo,
			mapper.NewUserMapper,
			services.NewUserService,
			controller.NewUserController,
		)

		return nil
}