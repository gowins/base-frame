package v1

import (
	"github.com/gowins/dionysus/ginhelper"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"

	"book/service/user/api/internal/logic/foo/v1"
	"book/service/user/api/internal/svc"
	"book/service/user/api/internal/types"
)

func LoginHandler(c *gin.Context) render.Render {
	var req types.LoginReq
	if err := c.ShouldBind(&req); err != nil {
		return ginHelper.Error(ginHelper.GinErrorParams)
	}

	l := v1.NewLoginLogic(req)
	resp, err := l.Login(&req)
	if err != nil {
		return ginHelper.Error(err)
	} else {
		return ginHelper.Success(resp)
	}
}
