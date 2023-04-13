package validators

import (
	"gigame.xyz/app/handlers/box1"
	"gigame.xyz/app/validators/v_data_01"
	"github.com/gin-gonic/gin"
)

func (v V) Topic(ctx *gin.Context) {
	var params v_data_01.VDataTopic
	bindData(ctx, &params, (box1.Api{}).Topic)
}
