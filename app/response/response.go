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
