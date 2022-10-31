package v1

import (
	v1 "base-frame/service/foo-multi/internal/logic/foo/v1"
	"base-frame/service/foo-multi/internal/types"

	"github.com/gowins/dionysus/ginx"

	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) ginx.Render {
	var req types.LoginReq
	if err := c.ShouldBind(&req); err != nil {
		return ginx.Error(ginx.GinErrorParams)
	}

	resp, err := v1.Login(&req)
	if err != nil {
		return ginx.Error(err)
	} else {
		return ginx.Success(resp)
	}
}
