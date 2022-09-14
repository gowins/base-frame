package ginhelper

import (
	"github.com/zeromicro/go-zero/rest"

	"github.com/gin-gonic/gin"
)

type ZeroGinRouter interface {
	GinRouters
	Group(path string, handler ...GinHandler) ZeroGinRouter
}

type ZeroGinServer interface {
	Run()
	Shutdown()
}

type GinRouter struct {
	ginGroup

	engine *gin.Engine
	server *rest.Server

	root bool
}

func NewZeroGinRouter(c rest.RestConf) *GinRouter {
	g := gin.New()
	g.Use(gin.Recovery()) // 默认注册recovery
	zg := &zeroGin{g}
	r := &GinRouter{
		ginGroup: ginGroup{g: &g.RouterGroup},
		engine:   g,
		root:     true,
	}
	r.server = rest.MustNewServer(c, rest.WithRouter(zg))
	return r
}

func (r *GinRouter) Group(path string, handler ...GinHandler) ZeroGinRouter {
	g := r.engine.Group(path, buildGinHandler(handler...)...)
	return &GinRouter{
		ginGroup: ginGroup{
			g: g,
		},
		engine: r.engine,
		server: r.server,
	}
}

// Run 启动
func (r *GinRouter) Run() {
	if r.root {
		r.server.Start()
	} else {

	}
}

// Shutdown 停止
func (r *GinRouter) Shutdown() {
	if r.root {
		r.server.Stop()
	}
}
