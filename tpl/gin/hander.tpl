package {{.PkgName}}

import (
        "net/http"

        "github.com/gowins/dionysus/ginx"

        "github.com/gin-gonic/gin"
        "github.com/gin-gonic/gin/render"

  {{.importPackages}}
)

func {{.HandlerName}}(c *gin.Context) ginx.Render {
    {{if .HasRequest}}var req types.{{.RequestType}}
    if err := c.ShouldBind(&req); err != nil {
        return ginx.Error(ginx.GinErrorParams)
    }

    {{end}}l := {{.LogicName}}.New{{.LogicType}}({{if .HasRequest}}req{{end}})
    {{if .HasResp}}resp, {{end}}err := l.{{.Call}}({{if .HasRequest}}&req{{end}})
    if err != nil {
        return ginx.Error(err)
    } else {
        {{if .HasResp}}return ginx.Success(resp){{else}}return ginx.Success("success"){{end}}
    }
}
