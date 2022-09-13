package zero_gin

import (
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	. "github.com/smartystreets/goconvey/convey"
)

func TestTimeout(t *testing.T) {
	Convey("超时", t, func() {
		r := NewZeroGinRouter(c)
		r.Use(TimeoutMiddleware(1000))
		r.GET("timeout", func(c *gin.Context) {
			time.Sleep(2 * time.Second)
			c.JSON(200, map[string]interface{}{})
		})

		res := testHttpRequest("GET", "/timeout", nil, r)
		So(res.Code, ShouldEqual, 504)
	})
}
