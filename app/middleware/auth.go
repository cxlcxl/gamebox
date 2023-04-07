package middleware

import (
	"encoding/base64"
	"gigame.xyz/app/response"
	"gigame.xyz/app/vars"
	"gigame.xyz/utils"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"time"
)

type Headers struct {
	Authorization string `json:"Authorization"`
}

func CheckAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		headers := Headers{}
		if err := ctx.ShouldBindHeader(&headers); err != nil {
			response.AuthFail(ctx)
			return
		}
		if headers.Authorization == "" {
			response.AuthFail(ctx)
			return
		}
		token := strings.Split(headers.Authorization, " ")
		if len(token) != 2 || len(token[1]) < 10 {
			response.AuthFail(ctx)
			return
		}
		key := vars.YmlConfig.GetString("BoxKeys." + token[0])
		if key == "" {
			response.AuthFail(ctx)
			return
		}
		if !checkSecret(key, token[1]) {
			response.AuthFail(ctx)
			return
		}
		ctx.Next()
	}
}

func checkSecret(k, s string) bool {
	decodeString, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return false
	}
	ds, err := utils.AesDecrypt(decodeString, []byte(k))
	if err != nil {
		return false
	}
	i, err := strconv.ParseInt(string(ds), 0, 64)
	if err != nil {
		return false
	}
	n := time.Now().Unix()
	if i-n <= 0 || i-n > 30 { // 有效期保证在 30s 内
		return false
	}
	return true
}
