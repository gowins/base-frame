package zero_gin

import (
	"net/http"
	"strings"

	"github.com/zeromicro/go-zero/rest"

	"github.com/gin-gonic/gin"
)

type ZeroGinInterface interface {
	gin.IRouter
	Run()
	Shutdown()
}

type ZeroGinRouter struct {
	*gin.Engine
	server *rest.Server
}

func NewZeroGinRouter(c rest.RestConf) ZeroGinInterface {
	g := gin.New()
	g.Use(gin.Recovery()) // 默认注册recovery
	zg := &zeroGin{g}
	r := &ZeroGinRouter{Engine: g}
	r.server = rest.MustNewServer(c, rest.WithRouter(zg))
	return r
}

// Run 启动
func (r *ZeroGinRouter) Run() {
	r.server.Start()
}

// Shutdown 停止
func (r *ZeroGinRouter) Shutdown() {
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
