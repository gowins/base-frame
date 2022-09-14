package main

import (
	"base-frame/ginhelper"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"github.com/zeromicro/go-zero/rest"
)

func main() {
	r := ginhelper.NewZeroGinRouter(rest.RestConf{
		Host: "0.0.0.0",
		Port: 9999,
	})
	r.Use(gin.Logger())
	r.Handle(http.MethodGet, "test", func(c *gin.Context) render.Render {
		return ginhelper.Success(time.Now().Unix())
	})
	ag := r.Group("admin/v1")
	ag.Handle(http.MethodPost, "get", func(c *gin.Context) render.Render {
		var tt = struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		}{}
		if err := c.ShouldBind(&tt); err != nil {
			return ginhelper.Error(err)
		}
		fmt.Printf("%v\n", tt)
		return ginhelper.Success(time.Now().Unix())
	})

	ag.Handle(http.MethodPost, "panic", func(c *gin.Context) render.Render {
		panic("hehe")
	})

	ag.Handle(http.MethodPost, "error", func(c *gin.Context) render.Render {
		return ginhelper.Error(ginhelper.NewGinError(350001, "请重新登陆"))
	})

	defer r.Shutdown()
	r.Run()
}
