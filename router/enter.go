package router

import (
	"myGo/router/customerRouter"
)

type RouterGroup struct {
	Customer customerRouter.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
