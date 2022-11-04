package example

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gowins/dionysus/ginx"
	"github.com/spf13/viper"
)

func TestError(c *gin.Context) ginx.Render {
	return ginx.Error(ginx.NewGinError(500100, "内部错误"))
}

func DemoRoute(c *gin.Context) ginx.Render {
	//会拿到demoMiddle中间件塞入gin.Contex的值
	va, ok := c.Get("demoMiddle")
	if !ok {
		return ginx.Error(ginx.NewGinError(500100, "内部错误"))
	}
	return ginx.Success(va)
}

func HttpCode(c *gin.Context) ginx.Render {
	// 设置http返回状态码
	c.Status(http.StatusGatewayTimeout)
	//c.JSON(http.StatusInternalServerError, "InternalServerError")
	return ginx.Error(ginx.NewGinError(500100, "内部错误"))
}

func GetConfig(c *gin.Context) ginx.Render {
	d := viper.GetDuration("mysql.max-connection-life-time")
	return ginx.Success(d.Seconds())
}
