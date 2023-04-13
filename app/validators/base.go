package validators

import "github.com/gin-gonic/gin"

type V struct{}

type Handler func(ctx *gin.Context, v interface{})

func bindData(ctx *gin.Context, v interface{}, h Handler) {
	if err := ctx.ShouldBind(v); err != nil {
		return
	}

	h(ctx, v)
}
