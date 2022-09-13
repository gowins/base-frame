package ginhelper

import (
	"net/http"
	"strings"

	"github.com/zeromicro/go-zero/rest"

	"github.com/gin-gonic/gin"
)

type ZeroGinRouter interface {
	gin.IRouter
	Run()
	Shutdown()
}

type ginRouter struct {
	*gin.Engine
	server *rest.Server
}

func NewZeroGinRouter(c rest.RestConf) ZeroGinRouter {
	g := gin.New()
	g.Use(gin.Recovery()) // 默认注册recovery
	zg := &zeroGin{g}
	r := &ginRouter{Engine: g}
	r.server = rest.MustNewServer(c, rest.WithRouter(zg))
	return r
}

// Run 启动
func (r *ginRouter) Run() {
	r.server.Start()
}

// Shutdown 停止
func (r *ginRouter) Shutdown() {
	r.server.Stop()
}

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
