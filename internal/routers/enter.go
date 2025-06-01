package routers

import "user_service/internal/routers/auth"

type RouterGroup struct {
	AuthRouterEnter *auth.AuthRouterGroup
}

var RouterGroupApp *RouterGroup = &RouterGroup{
	AuthRouterEnter: &auth.AuthRouterGroup{},
}
