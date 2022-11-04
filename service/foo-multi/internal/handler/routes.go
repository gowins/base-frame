package handler

import (
	"base-frame/service/foo-multi/internal/handler/example"
	"base-frame/service/foo-multi/internal/handler/middleware"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gowins/dionysus/ginx"
)

func RegisterHandler(routers ginx.ZeroGinRouter) {
	routers.Handle(http.MethodGet, "/test", func(c *gin.Context) ginx.Render {
		return ginx.Success(time.Now().Unix())
	})

	//访问路径 /config
	routers.Handle(http.MethodGet, "/config", example.GetConfig)

	//注意这时返回默认http code状态码200, 这里的code不是http code
	routers.Handle(http.MethodGet, "/test/error", example.TestError)

	demogroup := routers.Group("/demogroup")
	demogroup.Use(middleware.Auth, middleware.ContextM)

	//访问路径 /demogroup/demoroute
	demogroup.Handle(http.MethodGet, "/demoroute", example.DemoRoute)
	//会返回http code 504错误码
	demogroup.Handle(http.MethodGet, "/error", example.HttpCode)

	//访问路径 /empty
	emptyGroup := routers.Group("")
	emptyGroup.Use(middleware.ContextM2)
	emptyGroup.Handle(http.MethodGet, "/empty", func(c *gin.Context) ginx.Render {
		//会拿到ContextM2中间件塞入gin.Contex的值
		va, ok := c.Get("contextMiddle2")
		if !ok {
			return ginx.Error(ginx.NewGinError(500100, "内部错误"))
		}
		return ginx.Success(va)
	})
}
