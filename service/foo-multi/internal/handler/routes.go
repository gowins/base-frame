package handler

import (
	"net/http"

	foov1 "base-frame/service/foo-multi/internal/handler/foo/v1"

	"github.com/gin-gonic/gin"
	"github.com/gowins/dionysus/ginx"
)

func RegisterHandlers(gr ginx.ZeroGinRouter) {
	adminv1Group := gr.Group("foo/v1")
	adminv1Group.Use(ginx.TimeoutMiddleware(1000))
	adminv1Group.Use(ginx.LimiterMiddleware(1000))
	adminv1Group.Handle(http.MethodPost, "/task/get", foov1.LoginHandler)
	adminv1Group.Handle(http.MethodGet, "/ping", func(c *gin.Context) ginx.Render {
		return ginx.Success("pong")
	})
}
