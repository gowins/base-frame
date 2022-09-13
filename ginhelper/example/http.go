package main

import (
	"base-frame/ginhelper"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/zeromicro/go-zero/rest"
)

func main() {
	r := ginhelper.NewZeroGinRouter(rest.RestConf{
		Host: "0.0.0.0",
		Port: 9999,
	})
	r.Use(gin.Logger())
	r.GET("test", func(c *gin.Context) {
		c.JSON(200, ginhelper.Response{
			Code: 200,
			Msg:  "请求成功",
			Data: time.Now().Unix(),
		})
	})
	ag := r.Group("admin/v1")
	ag.POST("get", func(c *gin.Context) {
		var tt = struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		}{}
		if err := c.ShouldBind(&tt); err != nil {
			c.JSON(200, ginhelper.Response{
				Code: 400,
				Msg:  "请求成功",
				Data: time.Now().Unix(),
			})
		}
		fmt.Printf("%v\n", tt)
		c.JSON(200, ginhelper.Response{
			Code: 200,
			Msg:  "请求成功",
			Data: time.Now().Unix(),
		})
	})

	defer r.Shutdown()
	r.Run()
}
