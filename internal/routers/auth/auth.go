package auth

import (
	"user_service/internal/helper"
	"user_service/internal/wire"

	"github.com/gin-gonic/gin"
)

// login, refresh token, logout, revoke token, forgot password
type AuthRouter struct{}

var version_1 string = "v1/api"

func (a *AuthRouter) InitAuthRouter(Router *gin.RouterGroup) {
	controller := wire.InjectAuthController()
	// public router
	productRouterPublic := Router.Group("/auth")
	{
		productRouterPublic.POST("/login", controller.Login)                    // user login
		productRouterPublic.POST("/forgot-password", controller.ForgotPassword) // forgot password
		productRouterPublic.POST("/refresh-token", controller.RefreshToken)     // refresh token
	}

	// private router
	productRouterPrivate := Router.Group("/auth")
	{
		productRouterPrivate.POST("/logout", controller.Logout) // user logout
	}

	a.registerAuthRoute()
}

// register routes for documentation, backend usage and backup
func (ar *AuthRouter) registerAuthRoute() {
	helper.RegisterRoute("POST", version_1+"/auth/login")
	helper.RegisterRoute("POST", version_1+"/auth/forgot-password")
	helper.RegisterRoute("POST", version_1+"/auth/refresh-token")
	helper.RegisterRoute("POST", version_1+"/auth/logout")
}
