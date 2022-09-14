package ginhelper

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// 实现go-zero httpx.Router
type zeroGin struct {
	g *gin.Engine
}

func (zg *zeroGin) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	zg.g.ServeHTTP(writer, request)
}

func (zg *zeroGin) Handle(method, path string, handler http.Handler) error {
	zg.g.Handle(strings.ToUpper(method), path, func(ctx *gin.Context) {
		handler.ServeHTTP(ctx.Writer, ctx.Request)
	})
	return nil
}

func (zg *zeroGin) SetNotFoundHandler(handler http.Handler) {
	zg.g.NoRoute(gin.WrapH(handler))
}

func (zg *zeroGin) SetNotAllowedHandler(handler http.Handler) {
	zg.g.NoMethod(gin.WrapH(handler))
}
