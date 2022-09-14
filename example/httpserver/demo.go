package main

import (
	base "base-frame"
	"base-frame/cmd/zerogin"
	"net/http"

	"github.com/gowins/dionysus/ginhelper"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
)

func main() {
	ginCmd := zerogin.NewGinCommand()
	engine := ginCmd.GetEngine()
	addRoute(engine)
	base.Start("structure", "msg-srv", ginCmd)
}

func addRoute(engine ginhelper.ZeroGinRouter) {
	engine.Use(gin.Logger())
	adminGroup := engine.Group("admin/v1")
	adminGroup.Handle(http.MethodGet, "user/get", userGet)
	adminGroup.Handle(http.MethodPost, "user/post", userPost)
	adminGroup.Handle(http.MethodPost, "error", customError)
	adminGroup.Handle(http.MethodPost, "panic", customPanic)
}

func userGet(c *gin.Context) render.Render {
	return ginhelper.Success("获取用户成功")
}

func userPost(c *gin.Context) render.Render {
	var user = struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{}

	if err := c.ShouldBind(&user); err != nil {
		return ginhelper.Error(err)
	}

	return ginhelper.Success("修改用户成功")
}

func customError(c *gin.Context) render.Render {
	return ginhelper.Error(ginhelper.NewGinError(350001, "自定义错误"))
}

func customPanic(c *gin.Context) render.Render {
	panic("自定义panic")
}
