package v1

import (
	v1 "base-frame/services/user/api/internal/logic/foo/v1"

	"github.com/gowins/dionysus/ginx"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"

	"book/service/user/api/internal/types"
)

func LoginHandler(c *gin.Context) render.Render {
	var req types.LoginReq
	if err := c.ShouldBind(&req); err != nil {
		return ginx.Error(ginx.GinErrorParams)
	}

	l := v1.NewLoginLogic(req)
	resp, err := l.Login(&req)
	if err != nil {
		return ginx.Error(err)
	} else {
		return ginx.Success(resp)
	}
}
