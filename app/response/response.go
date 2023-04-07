package response

import (
	"fmt"
	"gigame.xyz/app/vars"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func HTML(ctx *gin.Context, page string, data interface{}) {
	tpl := page
	if !strings.HasSuffix(page, vars.TemplateExt) {
		tpl = fmt.Sprintf("%s.%s", page, vars.TemplateExt)
	}
	ctx.HTML(http.StatusOK, tpl, data)
}

func Success(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
		"data":    data,
	})
	ctx.Abort()
}

func AuthFail(ctx *gin.Context) {
	ctx.JSON(http.StatusBadGateway, struct{}{})
	ctx.Abort()
}
