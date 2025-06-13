package routers

import (
	"user_service/internal/routers/users"
)

type RouterGroup struct {
	UserRouterEnter *users.UserRouterGroup
}

var RouterGroupApp *RouterGroup = &RouterGroup{
	UserRouterEnter: &users.UserRouterGroup{},
}
