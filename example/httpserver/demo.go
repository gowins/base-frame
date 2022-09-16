package main

import (
	"base-frame/cmd"
	"net/http"

	base "base-frame"

	"github.com/gowins/dionysus/ginx"

	"github.com/gin-gonic/gin"
)

func main() {
	ginCmd := cmd.NewGinCommand()
	engine := ginCmd.GetEngine()
	addRoute(engine)
	base.Start("structure", "msg-srv", ginCmd)
}

func addRoute(engine ginx.ZeroGinRouter) {
	engine.Use(gin.Logger())
	adminGroup := engine.Group("admin/v1")
	adminGroup.Handle(http.MethodGet, "user/get", userGet)
	adminGroup.Handle(http.MethodPost, "user/post", userPost)

	webGroup := engine.Group("web/v1")
	webGroup.Handle(http.MethodGet, "user/get", userGet)
	webGroup.Handle(http.MethodPost, "user/post", userPost)
}

func userGet(c *gin.Context) ginx.Render {
	return ginx.Success("获取用户成功")
}

func userPost(c *gin.Context) ginx.Render {
	var user = struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{}

	if err := c.ShouldBind(&user); err != nil {
		return ginx.Error(err)
	}

	return ginx.Success("修改用户成功")
}

func customError(c *gin.Context) ginx.Render {
	return ginx.Error(ginx.NewGinError(350001, "自定义错误"))
}

func customPanic(c *gin.Context) ginx.Render {
	panic("自定义panic")
}
