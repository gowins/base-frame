package handler

import (
	"base-frame/service/bar-single/internal/handler/example"
	"base-frame/service/bar-single/internal/handler/middleware"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gowins/dionysus/ginx"
)

func RegisterHandler(routers ginx.ZeroGinRouter) {
	routers.Handle(http.MethodGet, "/test", func(c *gin.Context) ginx.Render {
		return ginx.Success(time.Now().Unix())
	})

	routers.Handle(http.MethodGet, "/config", example.GetConfig)

	//注意这时返回默认http code状态码200, 这里的code不是http code
	routers.Handle(http.MethodGet, "/test/error", example.TestError)

	demogroup := routers.Group("/demogroup")
	demogroup.Use(middleware.Auth, middleware.ContextM)
	demogroup.Handle(http.MethodGet, "/demoroute", example.DemoRoute)
	//会返回http code 504错误码
	demogroup.Handle(http.MethodGet, "/error", example.HttpCode)
}
