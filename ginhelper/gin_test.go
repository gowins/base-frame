package ginhelper

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/zeromicro/go-zero/rest"
)

var (
	c = rest.RestConf{
		Host:    "0.0.0.0",
		Port:    9999,
		Timeout: 1000,
	}
)

func TestNewGinRouter(t *testing.T) {
	Convey("new router", t, func() {
		r := NewZeroGinRouter(c)
		r.GET("/test-router", func(ctx *gin.Context) {
			ctx.JSON(200, map[string]interface{}{
				"code": 200,
				"msg":  "success",
				"data": []string{},
			})
		})

		res := testHttpRequest("GET", "/", nil, r)
		So(res.Code, ShouldEqual, 404)

		res = testHttpRequest("GET", "/test-router", nil, r)
		So(res.Code, ShouldEqual, 200)

		res = testHttpRequest("POST", "/test-router", nil, r)
		So(res.Code, ShouldEqual, 404)
	})
}

func TestNewGinRouter2(t *testing.T) {
	Convey("集成go-zero", t, func() {
		r := NewZeroGinRouter(c)

		r.GET("/test-router", func(ctx *gin.Context) {
			ctx.JSON(200, map[string]interface{}{
				"code": 200,
				"msg":  "success",
				"data": []string{},
			})
		})
		go func() {
			r.Run()
		}()

		res := testHttpRequest("GET", "/test-router", nil, r)
		So(res.Code, ShouldEqual, 200)

		r.Shutdown()

	})
}

func TestAddRouter(t *testing.T) {
	Convey("增加路由", t, func() {
		r := NewZeroGinRouter(c)
		ag := r.Group("/admin/v1")
		ug := ag.Group("user")
		ug.Any("add", func(ctx *gin.Context) {
			ctx.JSON(200, map[string]interface{}{})
		})
		ug.Any("delete", func(ctx *gin.Context) {
			ctx.JSON(200, map[string]interface{}{})
		})
		ug.Any("update", func(ctx *gin.Context) {
			ctx.JSON(200, map[string]interface{}{})
		})
		ug.Any("list", func(ctx *gin.Context) {
			ctx.JSON(200, map[string]interface{}{})
		})
		ug.Any("detail", func(ctx *gin.Context) {
			ctx.JSON(200, map[string]interface{}{})
		})

		r.Use(gin.Recovery())
		r.Use(timeout.New())
	})
}

func Test_zeroGin(t *testing.T) {
	Convey("zeroGin", t, func() {
		notFound := false
		notAllow := false
		zg := zeroGin{g: gin.New()}
		zg.g.HandleMethodNotAllowed = true
		zg.g.GET("/test", func(c *gin.Context) {
			c.JSON(http.StatusOK, nil)
		})
		zg.SetNotFoundHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			notFound = true
		}))
		zg.SetNotAllowedHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			notAllow = true
		}))
		req, _ := http.NewRequest("GET", "/notFound", nil)
		writer := httptest.NewRecorder()
		zg.ServeHTTP(writer, req)
		So(notFound, ShouldBeTrue)

		req, _ = http.NewRequest("POST", "/test", nil)
		writer = httptest.NewRecorder()
		zg.ServeHTTP(writer, req)
		So(notAllow, ShouldBeTrue)

		err := zg.Handle(http.MethodGet, "/a/b",
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				httpx.Ok(w)
			}),
		)
		So(err, ShouldBeNil)
		req, _ = http.NewRequest("GET", "/a/b", nil)
		writer = httptest.NewRecorder()
		zg.ServeHTTP(writer, req)
		So(writer.Code, ShouldEqual, http.StatusOK)
	})
}

func testHttpRequest(method, path string, body interface{}, r ZeroGinRouter) *httptest.ResponseRecorder {
	bs, _ := json.Marshal(body)
	req, _ := http.NewRequest(method, path, bytes.NewReader(bs))
	writer := httptest.NewRecorder()
	r.(*ginRouter).Engine.ServeHTTP(writer, req)
	return writer
}
